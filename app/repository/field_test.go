package repository

import (
	"fmt"
	"log"
	"testing"

	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMockDB() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening gorm database", err)
	}

	return gormDB, mock
}

// TODO: Quick test of unit testing. Delete at the end
func Test_FieldRepository_TestExamples(t *testing.T) {

	provider := db.NewMockProvider()
	repository := NewFieldRepository(provider)

	field := models.Field{
		Description: "test-Description",
	}
	field2 := models.Field{
		Description: "test-lululul",
	}
	repository.Create(&field)
	repository.Create(&field2)

	var myField models.Field
	res, err := repository.GetAll(myField)
	if err != nil {
		fmt.Println(err)
	}

	res2, err := repository.GetId(1, models.Field{})
	newField := res2.(*models.Field)
	fmt.Println(newField.Description)

	fields, _ := res.([]models.Field)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(field.Description, fields[0].Description))

}
