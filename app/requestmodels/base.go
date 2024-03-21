package requestmodels

import "time"

type RefId struct {
	ID uint `json:"id"`
}

type RefVersioned struct {
	ID      uint `json:"id"`
	Version uint `json:"version"`
}

type RefTimed struct {
	ID            uint   `json:"id" example:"1"`
	StartValidity string `json:"startvalidity" example:"02.01.2006"`
}

type RefSelectedCourse struct {
	UserID         uint      `gorm:"primarykey" json:"userId"`
	CourseID       uint      `gorm:"primarykey" json:"courseId"`
	CourseVersion  uint      `gorm:"primarykey" json:"courseVersion"`
	ClassStartyear time.Time `gorm:"primarykey" json:"classStartYear"`
	Dispensed      bool
}
