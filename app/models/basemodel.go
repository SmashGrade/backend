package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Basemodel is the base model for all models
type Basemodel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created"`
	UpdatedAt time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted,omitempty"`
}

// VersionedBasemodel is the base model for all models with versioning
type VersionedBasemodel struct {
	ID        uuid.UUID      `gorm:"type:uuid;primarykey;" json:"id"`
	CreatedAt time.Time      `json:"created"`
	UpdatedAt time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted,omitempty"`
	Version   uint           `gorm:"primarykey" json:"version"`
}
