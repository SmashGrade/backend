package db

import (
	"github.com/SmashGrade/backend/app/models"
)

func prefillMockDB(p Provider) {
	prefillFields(p)
}

func Field_1() models.Field {
	var field models.Field
	field.Description = "description Field 1"
	field.ID = 1
	return field
}

func Field_2() models.Field {
	var field models.Field
	field.Description = "description Field 2"
	field.ID = 2
	return field
}

func prefillFields(p Provider) {
	field_1 := Field_1()
	field_2 := Field_2()
	p.DB().Table("fields").Create(&field_1)
	p.DB().Table("fields").Create(&field_2)
}
