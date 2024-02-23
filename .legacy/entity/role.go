package entity

type Role struct {
	Basemodel
	Description string
	Users       []*User `gorm:"many2many:user_has_role;"`
}
