package dao

import (
	"github.com/SmashGrade/backend/app/api/v1/schemas"
	"github.com/SmashGrade/backend/app/entity"
)

func (db *Database) GetCourseEntity(course *entity.Course, id uint) error {
	db.Db.Where("id = ?", id).Find(&course)
	return nil
}

func (db *Database) ListCourses(coursesRes *[]schemas.CoursesRes) error {
	err := db.listCoursesRes(coursesRes)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) GetCourse(courseRes *schemas.CourseRes, courseId uint, version uint) error {
	err := db.getCourseRes(courseRes, courseId, version)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) GetCourseResStudent(courseResStudent *schemas.CourseResStudent, courseId uint, version uint, userId uint) error {
	err := db.getCourseResStudent(courseResStudent, courseId, version, userId)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) GetCourseResTeacher(courseResTeacher *schemas.CourseResTeacher, courseId uint, version uint, userId uint) error {
	err := db.getCourseResTeacher(courseResTeacher, courseId, version)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) GetCourseFilter(courseFilter *schemas.CourseFilter) error {
	err := db.getCourseFilter(courseFilter)
	if err != nil {
		return nil
	}

	return nil
}

// Post Course is always a new Course with Version 1.
func (db *Database) PostCourse(courseReq *schemas.CourseReqPost, version uint, id uint) error {
	var course entity.Course

	// If Id is not set, increment by 1
	if id == 0 {
		db.Db.Model(&entity.Course{}).Select("COALESCE(MAX(id), 0)").Scan(&id)
		id += 1
	}

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
		db.GetUser(&teacher, teacherRef)
		teachers = append(teachers, &teacher)
	}

	err := ParseSchemaToEntity(&courseReq, &course)
	if err != nil {
		return err
	}

	// create course
	course.Modules = modules
	course.TeachedBy = teachers
	course.Version = version
	course.ID = id
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

func (db *Database) PutCourse(courseReq *schemas.CourseReqPut, id uint) error {
	var course entity.Course

	// check if entry exist
	result := db.Db.Where("id = ? AND version = ?", id, courseReq.Version).Find(&entity.Course{})
	if result.RowsAffected == 0 {
		// Create New (All Exams are created New)
		var courseReqPost schemas.CourseReqPost
		ParseSchemaToEntity(&courseReq, &courseReqPost)
		db.PostCourse(&courseReqPost, courseReq.Version, id)

		// Create Exams
		var exams []*entity.Exam
		ParseSchemaToEntity(&courseReq.Exams, &exams)
		for _, exam := range exams {
			exam.Course = course
			db.PostExam(exam)
		}

		return nil
	}

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
		db.GetUser(&teacher, teacherRef)
		teachers = append(teachers, &teacher)
	}

	err := ParseSchemaToEntity(&courseReq, &course)
	if err != nil {
		return err
	}

	course.ID = id
	course.Modules = modules
	course.TeachedBy = teachers

	// Update (Only Exams with ID=0 created New)
	db.Db.Model(&course).Updates(&course)

	for _, examReq := range courseReq.Exams {
		var exam *entity.Exam
		ParseSchemaToEntity(&examReq, &exam)
		exam.Course = course

		if examReq.Id == 0 {
			db.PostExam(exam)
		} else {
			db.Db.Model(&exam).Updates(&exam)
		}
	}

	return nil
}
