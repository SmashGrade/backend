package api

import (
	"github.com/SmashGrade/backend/app/dao"
	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/repository"
)

// Handles the requests for the module resource
type CurriculumController struct {
	*BaseController
	Dao *dao.CurriculumDao
}

// Constructor for ModuleController
func NewCurriculumController(provider db.Provider) *CurriculumController {
	return &CurriculumController{
		BaseController: NewBaseController(provider),
		Dao:            dao.NewCurriculumDao(repository.NewCurriculumRepository(provider)),
	}
}
