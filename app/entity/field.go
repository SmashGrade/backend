package entity

type Field struct {
	Basemodel
	Description string
	Users       []*User `gorm:"many2many:fieldmanager;"`
}
