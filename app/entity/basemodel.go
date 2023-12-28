package entity

import "time"

// base model for gorm without deletedAt field
type Basemodel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
