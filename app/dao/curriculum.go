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

// creates new curriculum and returns id if successful
func (db *Database) CreateCurriculum(curriculum *schemas.Curriculum) (uint, error) {
	var newCurriculum entity.Curriculum

	err := ParseSchemaToEntity(curriculum, newCurriculum)
	newCurriculum.ID = 0

	return 0, err // TODO
}

// updates all values of an existing curriculum
func (db *Database) UpdateCurriculum(curriculum *schemas.Curriculum) error {
	return nil // TODO
}

// deletes an existing curriculum via id
func (db *Database) DeleteCurriculum(id uint) error {
	return nil // TODO
}
