package dao

import (
	"github.com/SmashGrade/backend/app/api/v1/schemas"
	"github.com/SmashGrade/backend/app/entity"
)

func (db *Database) GetModuleEntity(module *entity.Module, id uint, version uint) error {
	db.Db.Where("id = ? AND version = ?", id, version).Find(&module)
	return nil
}

func (db *Database) ListModules(modulesRes *[]schemas.ModuleRes) error {
	var modules []entity.Module

	db.Db.Preload("Courses").Find(&modules)

	for _, module := range modules {
		// TODO: Get evaluationType with EvaluationID
		var evaluationType entity.Evaluationtype
		var valuationCategory schemas.ValuationCategory

		db.Db.Where("id = ?", module.EvaluationTypeID).First(&evaluationType)
		ParseEntityToSchema(&evaluationType, &valuationCategory)

		// TODO: Get studyStage ????

		// Get Courses
		var courseRes []schemas.CourseRes
		db.ListCourses(&courseRes)

		var moduleRes schemas.ModuleRes
		ParseEntityToSchema(&module, &moduleRes)
		moduleRes.ValuationCategory = valuationCategory

		*modulesRes = append(*modulesRes, moduleRes)
	}

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
