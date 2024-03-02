package models

type Role struct {
	Basemodel
	Description string  `json:"description"`
	Claim       string  `json:"claim"`
	Users       []*User `gorm:"many2many:user_has_role;" json:"users"`
}
