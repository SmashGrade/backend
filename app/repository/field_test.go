package repository

import (
	"testing"

	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
)

func Test_Field_Create(t *testing.T) {
	repository := NewFieldRepository(db.NewMockProvider())

	field_1 := db.Field_1()
	field_1.ID = 0

	_, err := repository.Create(&field_1)

	require.NoError(t, err)
}

func Test_Field_Update(t *testing.T) {
	repository := NewFieldRepository(db.NewMockProvider())

	// Update Description of Field
	field := db.Field_1()
	field.Description = "edited description Field 1"
	err := repository.Update(&field)

	// Return all Fields for comparing
	result2, _ := repository.GetAll(models.Field{})
	fields := result2.([]models.Field)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(field.Description, fields[0].Description))
}

func Test_Field_Find(t *testing.T) {
	repository := NewFieldRepository(db.NewMockProvider())

	// Find Field
	result2, err := repository.Find(db.Field_1())
	fields := result2.([]models.Field)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.Field_1().ID, fields[0].ID))
}

func Test_Field_GetAll(t *testing.T) {
	repository := NewFieldRepository(db.NewMockProvider())

	// Get all fields
	result, err := repository.GetAll(models.Field{})
	fields := result.([]models.Field)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.Field_1().ID, fields[0].ID))
	require.Nil(t, deep.Equal(db.Field_2().ID, fields[1].ID))
}

func Test_Field_GetID(t *testing.T) {
	repository := NewFieldRepository(db.NewMockProvider())

	// Get by ID
	result, err := repository.GetId(db.Field_1().ID, models.Field{})
	field := result.(*models.Field)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(field.Description, db.Field_1().Description))
}

func Test_Field_DeleteId(t *testing.T) {
	repository := NewFieldRepository(db.NewMockProvider())

	// Get length of slice of all fields
	result, _ := repository.GetAll(models.Field{})
	afterCreateLength := len(result.([]models.Field))

	// Delete field
	err := repository.DeleteId(db.Field_1().ID, models.Field{})

	result2, _ := repository.GetAll((models.Field{}))
	afterDeleteLength := len(result2.([]models.Field))

	require.NoError(t, err)
	require.Nil(t, deep.Equal(afterCreateLength-1, afterDeleteLength))
}
