package entity

import "time"

// base model for gorm without deletedAt field
type basemodel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
