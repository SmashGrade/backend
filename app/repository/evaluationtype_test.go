package repository

import (
	"testing"

	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
)

func Test_Evaluationtype_Create(t *testing.T) {
	repository := NewEvaluationtypeRepository(db.NewPrefilledMockProvider())

	evaluationtype_1 := db.EvaluationType1
	evaluationtype_1.ID = 0

	_, err := repository.Create(&evaluationtype_1)

	require.NoError(t, err)
}

func Test_Evaluationtype_Update(t *testing.T) {
	repository := NewEvaluationtypeRepository(db.NewPrefilledMockProvider())

	// Update Description of Evaluationtype
	evaluationtype := db.EvaluationType1
	evaluationtype.Description = "edited Evaluationtype Description 1"
	err := repository.Update(&evaluationtype)

	// Return all Evaluationtype for comparing
	result2, _ := repository.GetAll()
	evaluationtypes := result2.([]models.Evaluationtype)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(evaluationtype.Description, evaluationtypes[0].Description))
}

func Test_Evaluationtype_Find(t *testing.T) {
	repository := NewEvaluationtypeRepository(db.NewPrefilledMockProvider())

	// Find Evaluationtype
	result2, err := repository.Find(db.EvaluationType1)
	evaluationtypes := result2.([]models.Evaluationtype)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.Field1.ID, evaluationtypes[0].ID))
}

func Test_Evaluationtype_GetAll(t *testing.T) {
	repository := NewEvaluationtypeRepository(db.NewPrefilledMockProvider())

	// Get all evaluationtyes
	result, err := repository.GetAll()
	evaluationtypes := result.([]models.Evaluationtype)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.Field1.ID, evaluationtypes[0].ID))
	require.Nil(t, deep.Equal(db.Field2.ID, evaluationtypes[1].ID))
}

func Test_Evaluationtype_GetID(t *testing.T) {
	repository := NewEvaluationtypeRepository(db.NewPrefilledMockProvider())

	// Get by ID
	result, err := repository.GetId(db.EvaluationType1.ID)
	evaluationtype := result.(*models.Evaluationtype)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(evaluationtype.Description, db.EvaluationType1.Description))
}

func Test_Evaluationtype_DeleteId(t *testing.T) {
	repository := NewEvaluationtypeRepository(db.NewPrefilledMockProvider())

	// Get length of slice of all evaluationtypes
	result, _ := repository.GetAll()
	afterCreateLength := len(result.([]models.Evaluationtype))

	// Delete evaluationtype
	err := repository.DeleteId(db.EvaluationType1.ID)

	result2, _ := repository.GetAll()
	afterDeleteLength := len(result2.([]models.Evaluationtype))

	require.NoError(t, err)
	require.Nil(t, deep.Equal(afterCreateLength-1, afterDeleteLength))
}
