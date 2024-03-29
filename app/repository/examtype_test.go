package repository

import (
	"testing"

	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
)

func Test_Examtype_Create(t *testing.T) {
	repository := NewExamtypeRepository(db.NewPrefilledMockProvider())

	examtype_1 := db.Examtype1
	examtype_1.ID = 0

	_, err := repository.Create(&examtype_1)

	require.NoError(t, err)
}

func Test_Examtype_Update(t *testing.T) {
	repository := NewExamtypeRepository(db.NewPrefilledMockProvider())

	// Update Description of Examtype
	examtype := db.Examtype1
	examtype.Description = "edited description Examtype 1"
	err := repository.Update(&examtype)

	// Return all Examtypes for comparing
	result2, _ := repository.GetAll()
	examtypes := result2.([]models.Examtype)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(examtype.Description, examtypes[0].Description))
}

func Test_Examtype_Find(t *testing.T) {
	repository := NewExamtypeRepository(db.NewPrefilledMockProvider())

	// Find Examtype
	result2, err := repository.Find(db.Examtype1)
	examtypes := result2.([]models.Examtype)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.Examtype1.ID, examtypes[0].ID))
}

func Test_Examtype_GetAll(t *testing.T) {
	repository := NewExamtypeRepository(db.NewPrefilledMockProvider())

	// Get all examtypes
	result, err := repository.GetAll()
	examtypes := result.([]models.Examtype)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.Examtype1.ID, examtypes[0].ID))
	require.Nil(t, deep.Equal(db.Examtype2.ID, examtypes[1].ID))
}

func Test_Examtype_GetID(t *testing.T) {
	repository := NewExamtypeRepository(db.NewPrefilledMockProvider())

	// Get by ID
	result, err := repository.GetId(db.Examtype1.ID)
	examtype := result.(*models.Examtype)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(examtype.Description, db.Examtype1.Description))
}

func Test_Examtype_DeleteId(t *testing.T) {
	repository := NewExamtypeRepository(db.NewPrefilledMockProvider())

	// Get length of slice of all examtypes
	result, _ := repository.GetAll()
	afterCreateLength := len(result.([]models.Examtype))

	// Delete examtype
	err := repository.DeleteId(db.Examtype1.ID)

	result2, _ := repository.GetAll()
	afterDeleteLength := len(result2.([]models.Examtype))

	require.NoError(t, err)
	require.Nil(t, deep.Equal(afterCreateLength-1, afterDeleteLength))
}
