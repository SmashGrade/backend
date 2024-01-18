package dao

import "github.com/SmashGrade/backend/app/api/v1/schemas"

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
