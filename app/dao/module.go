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

// returns module by id with maximum available version
func (db *Database) GetLatestModuleById(id uint) (module *entity.Module, err error) {
	module = &entity.Module{}
	err = db.Db.Preload(clause.Associations).Where("id = ? AND max(version)", id).Find(&module).Error
	return
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

// creates new module, returns module if successful
func (db *Database) CreateModule(moduleReq *schemas.ModuleReq) (*entity.Module, error) {
	module := &entity.Module{}

	// manual increment by one
	err := db.Db.Last(&module).Error
	if err != nil {
		module.ID = 1 // this is the first entry set id to 1
	} else {
		module.ID += 1
	}

	err = db.ParseModuleRefToEnt(moduleReq, module)
	if err != nil {
		return nil, err
	}
	module.Version = 1

	err = db.Db.Create(&module).Error
	if err != nil {
		return nil, err
	}

	return module, nil
}

// updates existing module
func (db *Database) UpdateModule(id uint, version uint, moduleReq *schemas.ModuleReq) (*entity.Module, error) {
	module := &entity.Module{}

	err := db.ParseModuleRefToEnt(moduleReq, module)
	if err != nil {
		return nil, err
	}

	module.ID = id
	module.Version = version

	err = db.Db.Save(&module).Error
	if err != nil {
		return nil, err
	}

	return module, nil
}

// deletes an existing module
func (db *Database) DeleteModule(id uint, version uint) error {
	var module entity.Module

	module.ID = id
	module.Version = version

	return db.Db.Delete(&module).Error
}
