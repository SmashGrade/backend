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
	//var course entity.Course

	return nil
}

func (db *Database) PutCourse(courseReq *schemas.CourseReqPut, id uint) error {
	//var course entity.Course

	return nil
}
