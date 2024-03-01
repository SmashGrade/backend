package repository

import (
	"testing"

	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
)

func Test_Curriculumtype_Create(t *testing.T) {
	repository := NewCurriculumtypeRepository(db.NewPrefilledMockProvider())

	field_1 := db.Field1
	field_1.ID = 0

	_, err := repository.Create(&field_1)

	require.NoError(t, err)
}

func Test_Curriculumtype_Update(t *testing.T) {
	repository := NewCurriculumtypeRepository(db.NewPrefilledMockProvider())

	// Update Description of Curriculumtype
	curriculumtype := db.CurriculumTyp1
	curriculumtype.ID = 1
	curriculumtype.Description = "Teilzeit"
	err := repository.Update(&curriculumtype)

	// Return all Curriculumtypes for comparing
	result2, _ := repository.GetAll()
	curriculumtypes := result2.([]models.Curriculumtype)

	require.NoError(t, err)
	require.Nil(t, deep.Equal("Teilzeit", curriculumtypes[0].Description))
}

func Test_Curriculumtype_Find(t *testing.T) {
	repository := NewCurriculumtypeRepository(db.NewPrefilledMockProvider())

	// Find Curriculumtype
	result2, err := repository.Find(db.CurriculumTyp1)
	curriculumtypes := result2.([]models.Curriculumtype)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.CurriculumTyp1.ID, curriculumtypes[0].ID))
}

func Test_Curriculumtype_GetAll(t *testing.T) {
	repository := NewCurriculumtypeRepository(db.NewPrefilledMockProvider())

	// Get all curriculumtypes
	result, err := repository.GetAll()
	curriculums := result.([]models.Curriculumtype)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.Curriculum1.ID, curriculums[0].ID))
	require.Nil(t, deep.Equal(db.Curriculum2.ID, curriculums[1].ID))
}

func Test_Curriclumtype_GetID(t *testing.T) {
	repository := NewCurriculumtypeRepository(db.NewPrefilledMockProvider())

	// Get by ID
	result, err := repository.GetId(db.CurriculumTyp1.ID)
	curriculumtype := result.(*models.Curriculumtype)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(curriculumtype.Description, db.CurriculumTyp1.Description))
}

func Test_Curriculumtype_DeleteId(t *testing.T) {
	repository := NewCurriculumtypeRepository(db.NewPrefilledMockProvider())

	// Get length of slice of all fields
	result, _ := repository.GetAll()
	afterCreateLength := len(result.([]models.Curriculumtype))

	// Delete field
	err := repository.DeleteId(db.Curriculum1.ID)

	result2, _ := repository.GetAll()
	afterDeleteLength := len(result2.([]models.Curriculumtype))

	require.NoError(t, err)
	require.Nil(t, deep.Equal(afterCreateLength-1, afterDeleteLength))
}
