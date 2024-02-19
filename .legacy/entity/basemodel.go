package entity

import (
	"time"

	"gorm.io/gorm"
)

// base model for gorm without deletedAt field
type Basemodel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
