package repository

import (
	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
)

type StudyStageRepository struct {
	*BaseRepository
}

func NewStudyStageRepository(provider db.Provider) *StudyStageRepository {
	return &StudyStageRepository{
		BaseRepository: NewBaseRepository(provider),
	}
}

func (r *StudyStageRepository) DeleteId(id uint) error {
	return r.Provider.DB().Delete(&models.StudyStage{}, id).Error
}
