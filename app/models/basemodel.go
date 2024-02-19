package models

import (
	"time"

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
// Bug: Autoincrement does not work with composite primary keys
type VersionedBasemodel struct {
	ID        uint           `gorm:"primarykey;autoincrement:false" json:"id"`
	Version   uint           `gorm:"primarykey" json:"version"`
	CreatedAt time.Time      `json:"created"`
	UpdatedAt time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted,omitempty"`
}

// TerminatedBasemodel is the base model for all models with a start time
// Bug: Autoincrement does not work with composite primary keys
type TerminatedBasemodel struct {
	ID            uint           `gorm:"primaryKey;autoincrement:false" json:"id"`
	StartValidity time.Time      `gorm:"primarykey;autoincrement:false" json:"startvalidity"`
	CreatedAt     time.Time      `json:"created"`
	UpdatedAt     time.Time      `json:"updated"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted,omitempty"`
}
