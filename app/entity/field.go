package entity

type Field struct {
	basemodel
	Description string
	Users       []*User `gorm:"many2many:fieldmanager;"`
}
