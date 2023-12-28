package entity

type Focus struct {
	basemodel
	FieldID     uint
	Field       Field `gorm:"foreignKey:FieldID;references:ID"`
	Description string
}
