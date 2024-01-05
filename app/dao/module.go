package dao

import (
	"github.com/SmashGrade/backend/app/api/v1/schemas"
	"github.com/SmashGrade/backend/app/entity"
)

func (db *Database) GetModuleEntity(module *entity.Module, id uint, version uint) error {
	db.Db.Where("id = ? AND version = ?", id, version).Find(&module)
	return nil
}

func (db *Database) PostModule(moduleReq *schemas.ModuleReq) error {
	var module entity.Module

	// Get List of Courses
	var courses []*entity.Course
	for _, courseRef := range moduleReq.CoursesRef {
		var course entity.Course
		db.GetCourseEntity(&course, uint(courseRef))
		courses = append(courses, &course)
	}

	err := ParseSchemaToEntity(&moduleReq, &module)
	if err != nil {
		return err
	}

	module.Courses = courses
	module.Version = 1

	db.Db.Create(&module)

	return nil
}
