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

func (db *Database) ListCoursesModuleStudent(modulesRes *[]schemas.ModuleRes, studyStage uint, userId uint) error {
	err := db.listModulesStudent(modulesRes, userId, studyStage)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) ListModuleTeacher(modulesRes *[]schemas.ModuleRes, userId uint, studyStage uint) error {
	err := db.listTeacherModules(modulesRes, userId, studyStage)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) GetModuleFilter(moduleFilter *schemas.ModuleFilter) error {
	err := db.getModuleFilter(moduleFilter)
	if err != nil {
		return err
	}

	return nil
}

// base parsing module schema to entity
func (db *Database) ParseModuleRefToEnt(moduleReq *schemas.ModuleReq, moduleEnt *entity.Module) error {
	// Get List of Courses
	var courses []*entity.Course
	err := db.GetCourseListFromCourseRefList(moduleReq.CoursesRef, courses)
	if err != nil {
		return err
	}

	err = ParseSchemaToEntity(&moduleReq, &moduleEnt)
	if err != nil {
		return err
	}

	moduleEnt.Courses = courses
	return nil
}

// creates new module, returns id if successful
func (db *Database) CreateModule(moduleReq *schemas.ModuleReq) (uint, error) {
	var module entity.Module

	err := db.ParseModuleRefToEnt(moduleReq, &module)
	if err != nil {
		return 0, err
	}
	module.Version = 1

	err = db.Db.Create(&module).Error
	if err != nil {
		return 0, err
	}

	return module.ID, nil
}

// updates existing module
func (db *Database) UpdateModule(id uint, version uint, moduleReq *schemas.ModuleReq) error {
	var module entity.Module

	err := db.ParseModuleRefToEnt(moduleReq, &module)
	if err != nil {
		return err
	}

	module.ID = id
	module.Version = version

	err = db.Db.Save(&module).Error
	if err != nil {
		return err
	}

	return nil
}

// deletes an existing module
func (db *Database) DeleteModule(id uint, version uint) error {
	var module entity.Module

	module.ID = id
	module.Version = version

	return db.Db.Delete(&module).Error
}
