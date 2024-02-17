package repository

import (
	"testing"

	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
)

func Test_Evaluationtype_Create(t *testing.T) {
	repository := NewEvaluationtypeRepository(db.NewMockProvider())

	evaluationtype_1 := db.Evaluationtype_1()
	evaluationtype_1.ID = 0

	_, err := repository.Create(&evaluationtype_1)

	require.NoError(t, err)
}

func Test_Evaluationtype_Update(t *testing.T) {
	repository := NewEvaluationtypeRepository(db.NewMockProvider())

	// Update Description of Evaluationtype
	evaluationtype := db.Evaluationtype_1()
	evaluationtype.Description = "edited Evaluationtype Description 1"
	err := repository.Update(&evaluationtype)

	// Return all Evaluationtype for comparing
	result2, _ := repository.GetAll()
	evaluationtypes := result2.([]models.Evaluationtype)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(evaluationtype.Description, evaluationtypes[0].Description))
}

func Test_Evaluationtype_Find(t *testing.T) {
	repository := NewEvaluationtypeRepository(db.NewMockProvider())

	// Find Evaluationtype
	result2, err := repository.Find(db.Evaluationtype_1())
	evaluationtypes := result2.([]models.Evaluationtype)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.Field_1().ID, evaluationtypes[0].ID))
}

func Test_Evaluationtype_GetAll(t *testing.T) {
	repository := NewEvaluationtypeRepository(db.NewMockProvider())

	// Get all evaluationtyes
	result, err := repository.GetAll()
	evaluationtypes := result.([]models.Evaluationtype)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.Field_1().ID, evaluationtypes[0].ID))
	require.Nil(t, deep.Equal(db.Field_2().ID, evaluationtypes[1].ID))
}

func Test_Evaluationtype_GetID(t *testing.T) {
	repository := NewEvaluationtypeRepository(db.NewMockProvider())

	// Get by ID
	result, err := repository.GetId(db.Evaluationtype_1().ID)
	evaluationtype := result.(*models.Evaluationtype)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(evaluationtype.Description, db.Evaluationtype_1().Description))
}

func Test_Evaluationtype_DeleteId(t *testing.T) {
	repository := NewEvaluationtypeRepository(db.NewMockProvider())

	// Get length of slice of all evaluationtypes
	result, _ := repository.GetAll()
	afterCreateLength := len(result.([]models.Evaluationtype))

	// Delete evaluationtype
	err := repository.DeleteId(db.Evaluationtype_1().ID)

	result2, _ := repository.GetAll()
	afterDeleteLength := len(result2.([]models.Evaluationtype))

	require.NoError(t, err)
	require.Nil(t, deep.Equal(afterCreateLength-1, afterDeleteLength))
}
