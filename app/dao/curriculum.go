package dao

import (
	"fmt"

	"github.com/SmashGrade/backend/app/api/v1/schemas"
	"github.com/SmashGrade/backend/app/entity"
)

func (db *Database) ListCurriculum(curriculumRes *[]schemas.CurriculumRes) error {
	err := db.listCurriculumRes(curriculumRes)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) GetCurriculum(curriculumRes *schemas.CurriculumRes, id uint) error {
	err := db.getCurriculumRes(curriculumRes, id)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) GetCurriculumFilter(curriculumFilter *schemas.CurriculumFilter) error {
	err := db.getCurriculumFilter(curriculumFilter)
	if err != nil {
		return err
	}

	return nil
}

// creates new curriculum out of req struct, all linked ressources must be existing
// Returns entity of newly created curriculum
func (db *Database) CreateCurriculum(curriculum *schemas.CurriculumReq) (*entity.Curriculum, error) {

	// init entity struct
	curriculumEntity := &entity.Curriculum{}

	// do an autoparse
	err := ParseSchemaToEntity(&curriculum, &curriculumEntity)
	if err != nil {
		return nil, err
	}

	// get our in between ids of field and focus from the strings
	focus, err := db.GetFocusByDescription(curriculum.Focus)
	if err != nil {
		return nil, err // probably no focus with that desc
	}

	field, err := db.GetFieldByDescription(curriculum.Field)
	if err != nil {
		return nil, err // probably no field with that desc
	}

	// get curriculum type from description
	curriculumType, err := db.GetCurriculumTypeByDescription(curriculum.CurriculumType)
	if err != nil {
		return nil, err // probably no curriculum type with that desc
	}

	// check if those assigned field managers exist
	for _, uid := range curriculum.FieldmanagerRef {
		user_found := false
		for _, user := range field.Users {
			if uid == user.ID {
				user_found = true
				break
			}
		}
		if !user_found {
			return nil, fmt.Errorf("no user with id '%v' found", uid)
		}
	}

	// get modules
	for _, modId := range curriculum.ModulesRef {
		modEnt, err := db.GetLatestModuleById(modId)
		if err != nil {
			return nil, fmt.Errorf("no module with id '%v' found", modId)
		}
		curriculumEntity.Modules = append(curriculumEntity.Modules, modEnt)
	}

	curriculumEntity.Focus = *focus // add found focus to the autocasted curriculum
	curriculumEntity.Curriculumtype = *curriculumType

	err = db.Db.Create(&curriculumEntity).Error
	if err != nil {
		return nil, err
	}

	return curriculumEntity, nil
}

func (db *Database) GetCurriculumTypeByDescription(description string) (curriculumtype *entity.Curriculumtype, err error) {
	curriculumtype = &entity.Curriculumtype{}
	err = db.Db.Model(&entity.Curriculumtype{}).Where("description = ?", description).First(&curriculumtype).Error
	return
}

func (db *Database) CreateCurriculumType(description string, durationYears uint) (curriculumtype *entity.Curriculumtype, err error) {
	curriculumtype = &entity.Curriculumtype{}
	curriculumtype.Description = description
	curriculumtype.DurationYears = durationYears
	err = db.Db.Create(&curriculumtype).Error
	return
}
