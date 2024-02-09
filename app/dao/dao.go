package dao

import (
	e "github.com/SmashGrade/backend/app/error"
	"github.com/SmashGrade/backend/app/models"
)

type CurriculumType struct {
}

func (c *CurriculumType) Get(id uint) (entity models.Curriculumtype, err e.DaoError)
