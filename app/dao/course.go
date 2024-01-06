package dao

import (
	"github.com/SmashGrade/backend/app/api/v1/schemas"
	"github.com/SmashGrade/backend/app/entity"
)

func (db *Database) GetCourseEntity(course *entity.Course, id uint) error {
	db.Db.Where("id = ?", id).Find(&course)
	return nil
}

func (db *Database) ListCourses(coursesRes *[]schemas.CourseRes) error {
	var courses []entity.Course
	db.Db.Find(&courses)
	err := ParseEntityToSchema(&courses, &coursesRes)
	if err != nil {
		return err
	}
	return nil
}

// Post Course is always a new Course with Version 1.
func (db *Database) PostCourse(courseReq *schemas.CourseReqPost) error {
	var course entity.Course

	// Get List of Modules
	var modules []*entity.Module
	for _, moduleRef := range courseReq.ModuleRef {
		var module entity.Module
		db.GetModuleEntity(&module, moduleRef.Id, moduleRef.Version)
		modules = append(modules, &module)
	}

	// Get List of Teachers
	var teachers []*entity.User
	for _, teacherRef := range courseReq.TeacherRef {
		var teacher entity.User
		db.GetUser(&teacher, uint(teacherRef))
		teachers = append(teachers, &teacher)
	}

	err := ParseSchemaToEntity(&courseReq, &course)
	if err != nil {
		return err
	}

	// create course
	course.Modules = modules
	course.TeachedBy = teachers
	course.Version = 1
	db.Db.Create(&course)

	// Create Exams
	var exams []*entity.Exam
	ParseSchemaToEntity(&courseReq.Exams, &exams)
	for _, exam := range exams {
		exam.Course = course
		db.PostExam(exam)
	}

	return nil
}

func (db *Database) GetCourse(courseRes *schemas.CourseRes, id uint, version uint) error {
	var course entity.Course

	if version != 0 {
		db.Db.Where("id = ? AND version = ?", id, version).Find(&course)
	} else {
		db.Db.Preload("Modules").Preload("TeachedBy").Where("id = ?", id).Order("version desc").First(&course)
	}

	err := ParseEntityToSchema(&course, &courseRes)
	if err != nil {
		return err
	}

	// Add Teachers
	var teachers []schemas.Teacher
	err = ParseEntityToSchema(&course.TeachedBy, &teachers)
	if err != nil {
		return err
	}
	courseRes.Teachers = teachers

	// Get all Versions
	var versions []uint
	db.Db.Model(&entity.Course{}).Where("id = ?", id).Pluck("version", &versions)
	courseRes.Versions = versions

	// Get all Exams
	var exams []entity.Exam
	var examsRes []schemas.ExamRes
	db.Db.Where("course_id = ?", id).Find(&exams)
	err = ParseEntityToSchema(&exams, &examsRes)
	if err != nil {
		return err
	}
	courseRes.Exams = examsRes

	return nil
}
