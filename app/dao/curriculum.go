package dao

import (
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

	curriculumEntity := &entity.Curriculum{}

	err := ParseSchemaToEntity(&curriculum, &curriculumEntity)
	if err != nil {
		return 0, err
	}

	err = db.Db.Create(&curriculumEntity).Error
	if err != nil {
		return 0, err
	}

	return curriculumEntity.ID, nil
}
