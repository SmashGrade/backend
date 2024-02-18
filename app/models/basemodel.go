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

// sets id to random uuid
func (v *VersionedBasemodel) GenerateId() {
	v.ID = uuid.New()
}

// sets id to random uuid if it is initial, returns true if id is generated
func (v *VersionedBasemodel) GenerateIdIfEmpty() bool {
	var emptyUUID = uuid.UUID{}
	if v.ID == emptyUUID {
		v.GenerateId()
		return true
	}
	return false
}
