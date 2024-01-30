package dao

import (
	"github.com/SmashGrade/backend/app/api/v1/schemas"
	"github.com/SmashGrade/backend/app/entity"
)

func (db *Database) GetCourseEntity(course *entity.Course, id uint, version uint) error {
	db.Db.Where("id = ?", id).Where("version = ?", version).Find(&course)
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
func (db *Database) CreateCourse(courseReq *schemas.CourseReqPost) (*entity.Course, error) {
	course := &entity.Course{}

	course.Version = 1

	err := db.Db.Create(course).Error
	if err != nil {
		return nil, err
	}

	return course, nil
}

func (db *Database) PutCourse(courseReq *schemas.CourseReqPut, id uint) error {
	//var course entity.Course

	return nil
}

// selects all matching courses entities from courseRef, saves in outCourseEnt, throws if one is not found
func (db *Database) GetCourseListFromCourseRefList(inCourseRef []schemas.CourseRef, outCourseEnt []*entity.Course) error {
	for _, courseRef := range inCourseRef {
		var course entity.Course
		err := db.GetCourseEntity(&course, courseRef.Id, courseRef.Version)
		if err != nil {
			return err
		}
		outCourseEnt = append(outCourseEnt, &course)
	}
	return nil
}
