package repository

import (
	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
)

type GradetypeRepository struct {
	*BaseRepository
}

func NewGradetypeRepository(provider db.Provider) *GradetypeRepository {
	return &GradetypeRepository{
		BaseRepository: NewBaseRepository(provider, models.Gradetype{}),
	}
}
