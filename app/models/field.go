package models

type Field struct {
	Basemodel
	Description string  `json:"description"`
	Users       []*User `gorm:"many2many:fieldmanager;" json:"users"`
}
