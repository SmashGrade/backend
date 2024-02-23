package repository

import (
	"testing"

	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
)

func Test_Gradetype_Create(t *testing.T) {
	repository := NewGradetypeRepository(db.NewPrefilledMockProvider())

	gradetype_1 := db.Gradetype_1()
	gradetype_1.ID = 0

	_, err := repository.Create(&gradetype_1)

	require.NoError(t, err)
}

func Test_Gradetype_Update(t *testing.T) {
	repository := NewGradetypeRepository(db.NewPrefilledMockProvider())

	// Update Description of Gradetype
	gradetype := db.Gradetype_1()
	gradetype.Description = "edited description Gradetype 1"
	err := repository.Update(&gradetype)

	// Return all Gradetypes for comparing
	result2, _ := repository.GetAll()
	gradetypes := result2.([]models.Gradetype)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(gradetype.Description, gradetypes[0].Description))
}

func Test_Gradetype_Find(t *testing.T) {
	repository := NewGradetypeRepository(db.NewPrefilledMockProvider())

	// Find Gradetype
	result2, err := repository.Find(db.Gradetype_1())
	gradetypes := result2.([]models.Gradetype)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.Gradetype_1().ID, gradetypes[0].ID))
}

func Test_Gradetype_GetAll(t *testing.T) {
	repository := NewGradetypeRepository(db.NewPrefilledMockProvider())

	// Get all gradetypes
	result, err := repository.GetAll()
	gradetypes := result.([]models.Gradetype)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.Gradetype_1().ID, gradetypes[0].ID))
	require.Nil(t, deep.Equal(db.Gradetype_2().ID, gradetypes[1].ID))
}

func Test_Gradetype_GetID(t *testing.T) {
	repository := NewGradetypeRepository(db.NewPrefilledMockProvider())

	// Get by ID
	result, err := repository.GetId(db.Gradetype_1().ID)
	gradetype := result.(*models.Gradetype)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(gradetype.Description, db.Gradetype_1().Description))
}

func Test_Gradetype_DeleteId(t *testing.T) {
	repository := NewGradetypeRepository(db.NewPrefilledMockProvider())

	// Get length of slice of all gradetypes
	result, _ := repository.GetAll()
	afterCreateLength := len(result.([]models.Gradetype))

	// Delete gradetype
	err := repository.DeleteId(db.Gradetype_1().ID)

	result2, _ := repository.GetAll()
	afterDeleteLength := len(result2.([]models.Gradetype))

	require.NoError(t, err)
	require.Nil(t, deep.Equal(afterCreateLength-1, afterDeleteLength))
}
