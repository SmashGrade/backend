package dao

import (
	"github.com/SmashGrade/backend/app/api/v1/schemas"
	"github.com/SmashGrade/backend/app/entity"
	"gorm.io/gorm/clause"
)

func (db *Database) GetModuleEntity(module *entity.Module, id uint, version uint) error {
	db.Db.Preload(clause.Associations).Where("id = ? AND version = ?", id, version).Find(&module)
	return nil
}

func (db *Database) ListModule(modulesRes *[]schemas.ModuleRes) error {
	err := db.listModuleRes(modulesRes)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) GetModule(moduleRes *schemas.ModuleRes, moduleId uint, version uint) error {
	err := db.getModuleRes(moduleRes, moduleId, version)
	if err != nil {
		return err
	}

	return nil
}

// TODO:
func (db *Database) GetModuleStudent(modulesRes *[]schemas.ModuleRes, studyStage uint, userId uint) error {
	var allModules []schemas.ModuleRes
	err := db.listModuleRes(&allModules)
	if err != nil {
		return err
	}

	for _, module := range allModules {
		if module.StudyStage.Id == studyStage {
			*modulesRes = append(*modulesRes, module)
		}
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
