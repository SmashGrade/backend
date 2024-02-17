package db

import (
	"time"

	"github.com/SmashGrade/backend/app/models"
)

// Fill MockDB with prefill functions
func prefillMockDB(p Provider) {
	prefillFields(p)
	prefillConversions(p)
}

// Field that will be added to the mock DB
func Field_1() models.Field {
	var field models.Field
	field.Description = "description Field 1"
	field.ID = 1
	return field
}

// Field that will be added to the mock DB
func Field_2() models.Field {
	var field models.Field
	field.Description = "description Field 2"
	field.ID = 2
	return field
}

// add all the fields to the fields table of the mockDB
func prefillFields(p Provider) {
	field_1 := Field_1()
	field_2 := Field_2()
	p.DB().Table("fields").Create(&field_1)
	p.DB().Table("fields").Create(&field_2)
}

// Conversion that will be added to the mock DB
func Conversion_1() models.Conversion {
	var conversion models.Conversion
	conversion.EEExamID = 1
	conversion.EESelectedCourseUserID = 1
	conversion.EESelectedCourseCourseVersion = 1
	conversion.EESelectedCourseClassStartyear = time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)
	conversion.EEExamID = 1
	conversion.Value = 3.8
	conversion.ID = 1
	return conversion
}

// Conversion that will be added to the mock DB
func Conversion_2() models.Conversion {
	var conversion models.Conversion
	conversion.EEExamID = 2
	conversion.EESelectedCourseUserID = 2
	conversion.EESelectedCourseCourseVersion = 2
	conversion.EESelectedCourseClassStartyear = time.Date(2024, time.February, 2, 0, 0, 0, 0, time.UTC)
	conversion.EEExamID = 2
	conversion.Value = 5.5
	conversion.ID = 2
	return conversion
}

// add all the fields to the fields table of the mockDB
func prefillConversions(p Provider) {
	conversion_1 := Conversion_1()
	conversion_2 := Conversion_2()
	p.DB().Table("conversions").Create(&conversion_1)
	p.DB().Table("conversions").Create(&conversion_2)
}
