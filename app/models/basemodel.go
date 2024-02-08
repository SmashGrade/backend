package models

import (
	"time"

	"gorm.io/gorm"
)

// Basemodel is the base model for all models
type Basemodel struct {
	Id        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created"`
	UpdatedAt time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted,omitempty"`
}

// VersionedBasemodel is the base model for all models with versioning
type VersionedBasemodel struct {
	Basemodel
	Version uint `gorm:"primarykey" json:"version"`
}
