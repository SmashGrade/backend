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
// Returns id of newly created curriculum
func (db *Database) CreateCurriculum(curriculum *schemas.CurriculumReq) (uint, error) {

	// get our in between ids of field and focus from the strings
	focus, err := db.GetFocusByDescription(curriculum.Focus)
	if err != nil {
		return 0, err // probably no focus with that desc
	}

	field, err := db.GetFieldByDescription(curriculum.Field)
	if err != nil {
		return 0, err // probably no field with that desc
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
			return 0, fmt.Errorf("no user with id '%v' found", uid)
		}
	}

	// TODO: get modules

	curriculumEntity := &entity.Curriculum{}

	err = ParseSchemaToEntity(&curriculum, &curriculumEntity)
	if err != nil {
		return 0, err
	}

	curriculumEntity.Focus = *focus // add found focus to the autocasted curriculum

	err = db.Db.Create(&curriculumEntity).Error
	if err != nil {
		return 0, err
	}

	return curriculumEntity.ID, nil
}
