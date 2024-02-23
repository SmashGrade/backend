package repository

import (
	"testing"

	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
)

func Test_Conversion_Create(t *testing.T) {
	repository := NewConversionRepository(db.NewPrefilledMockProvider())

	conversion_1 := db.Conversion_1()
	conversion_1.ID = 0

	_, err := repository.Create(&conversion_1)

	require.NoError(t, err)
}

func Test_Conversion_Update(t *testing.T) {
	repository := NewConversionRepository(db.NewPrefilledMockProvider())

	// Update Description of Field
	conversion := db.Conversion_1()
	conversion.Value = 4.7
	err := repository.Update(&conversion)

	// Return all Fields for comparing
	result2, _ := repository.GetAll()
	conversions := result2.([]models.Conversion)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(conversion.Value, conversions[0].Value))
}

func Test_Conversion_Find(t *testing.T) {
	repository := NewConversionRepository(db.NewPrefilledMockProvider())

	// Find Field
	result2, err := repository.Find(db.Conversion_1())
	fields := result2.([]models.Conversion)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.Field_1().ID, fields[0].ID))
}

func Test_Conversion_GetAll(t *testing.T) {
	repository := NewConversionRepository(db.NewPrefilledMockProvider())

	// Get all fields
	result, err := repository.GetAll()
	conversions := result.([]models.Conversion)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.Conversion_1().ID, conversions[0].ID))
	require.Nil(t, deep.Equal(db.Conversion_2().ID, conversions[1].ID))
}

func Test_Field_GetTimed(t *testing.T) {
	repository := NewConversionRepository(db.NewPrefilledMockProvider())

	// Get by ID and Start year
	result, err := repository.GetTimed(db.Conversion_1().ID, db.Conversion_1().EESelectedCourseClassStartyear)
	conversion := result.(*models.Conversion)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(conversion.Value, db.Conversion_1().Value))
}

func Test_Field_DeleteTimed(t *testing.T) {
	repository := NewConversionRepository(db.NewPrefilledMockProvider())

	// Get length of slice of all conversions
	result, _ := repository.GetAll()
	afterCreateLength := len(result.([]models.Conversion))

	// Delete conversion
	err := repository.DeleteTimed(db.Conversion_1().ID, db.Conversion_1().EESelectedCourseClassStartyear)

	result2, _ := repository.GetAll()
	afterDeleteLength := len(result2.([]models.Conversion))

	require.NoError(t, err)
	require.Nil(t, deep.Equal(afterCreateLength-1, afterDeleteLength))
}
