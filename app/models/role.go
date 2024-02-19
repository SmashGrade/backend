package models

type Role struct {
	Basemodel
	Description string  `json:"description"`
	Users       []*User `gorm:"many2many:user_has_role;" json:"users"`
}
