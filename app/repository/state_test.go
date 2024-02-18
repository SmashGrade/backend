package repository

import (
	"testing"

	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
)

func Test_State_Create(t *testing.T) {
	repository := NewStateRepository(db.NewMockProvider())

	state_1 := db.State_1()
	state_1.ID = 0

	_, err := repository.Create(&state_1)

	require.NoError(t, err)
}

func Test_State_Update(t *testing.T) {
	repository := NewStateRepository(db.NewMockProvider())

	// Update Description of State
	state := db.State_1()
	state.Description = "edited description State 1"
	err := repository.Update(&state)

	// Return all States for comparing
	result2, _ := repository.GetAll()
	states := result2.([]models.State)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(state.Description, states[0].Description))
}

func Test_State_Find(t *testing.T) {
	repository := NewStateRepository(db.NewMockProvider())

	// Find State
	result2, err := repository.Find(db.State_1())
	states := result2.([]models.State)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.State_1().ID, states[0].ID))
}

func Test_State_GetAll(t *testing.T) {
	repository := NewStateRepository(db.NewMockProvider())

	// Get all states
	result, err := repository.GetAll()
	states := result.([]models.State)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.State_1().ID, states[0].ID))
	require.Nil(t, deep.Equal(db.State_2().ID, states[1].ID))
}

func Test_State_GetID(t *testing.T) {
	repository := NewStateRepository(db.NewMockProvider())

	// Get by ID
	result, err := repository.GetId(db.State_1().ID)
	state := result.(*models.State)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(state.Description, db.State_1().Description))
}

func Test_State_DeleteId(t *testing.T) {
	repository := NewStateRepository(db.NewMockProvider())

	// Get length of slice of all states
	result, _ := repository.GetAll()
	afterCreateLength := len(result.([]models.State))

	// Delete state
	err := repository.DeleteId(db.State_1().ID)

	result2, _ := repository.GetAll()
	afterDeleteLength := len(result2.([]models.State))

	require.NoError(t, err)
	require.Nil(t, deep.Equal(afterCreateLength-1, afterDeleteLength))
}