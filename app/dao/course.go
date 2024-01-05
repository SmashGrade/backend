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
