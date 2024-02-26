package repository

import (
	"testing"

	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
)

func Test_Focus_Create(t *testing.T) {
	repository := NewFocusRepository(db.NewPrefilledMockProvider())

	focus_1 := db.Focus_1()
	focus_1.ID = 0

	_, err := repository.Create(&focus_1)

	require.NoError(t, err)
}

func Test_Focus_Update(t *testing.T) {
	repository := NewFocusRepository(db.NewPrefilledMockProvider())

	// Update Description of Focus
	focus := db.Focus_1()
	focus.Description = "edited description Focus 1"
	err := repository.Update(&focus)

	// Return all Focuss for comparing
	result2, _ := repository.GetAll()
	focuses := result2.([]models.Focus)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(focus.Description, focuses[0].Description))
}

func Test_Focus_Find(t *testing.T) {
	repository := NewFocusRepository(db.NewPrefilledMockProvider())

	// Find Focus
	result2, err := repository.Find(db.Focus_1())
	focuses := result2.([]models.Focus)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.Focus_1().ID, focuses[0].ID))
}

func Test_Focus_GetAll(t *testing.T) {
	repository := NewFocusRepository(db.NewPrefilledMockProvider())

	// Get all focuses
	result, err := repository.GetAll()
	focuses := result.([]models.Focus)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.Focus_1().ID, focuses[0].ID))
	require.Nil(t, deep.Equal(db.Focus_2().ID, focuses[1].ID))
}

func Test_Focus_GetID(t *testing.T) {
	repository := NewFocusRepository(db.NewPrefilledMockProvider())

	// Get by ID
	result, err := repository.GetId(db.Focus_1().ID)
	focus := result.(*models.Focus)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(focus.Description, db.Focus_1().Description))
}

func Test_Focus_DeleteId(t *testing.T) {
	repository := NewFocusRepository(db.NewPrefilledMockProvider())

	// Get length of slice of all focuses
	result, _ := repository.GetAll()
	afterCreateLength := len(result.([]models.Focus))

	// Delete focus
	err := repository.DeleteId(db.Focus_1().ID)

	result2, _ := repository.GetAll()
	afterDeleteLength := len(result2.([]models.Focus))

	require.NoError(t, err)
	require.Nil(t, deep.Equal(afterCreateLength-1, afterDeleteLength))
}
