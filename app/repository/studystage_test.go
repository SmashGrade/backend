package repository

import (
	"testing"

	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
)

func Test_StudyStage_Create(t *testing.T) {
	repository := NewStudyStageRepository(db.NewMockProvider())

	studyStage_1 := db.StudyStage_1()
	studyStage_1.ID = 0

	_, err := repository.Create(&studyStage_1)

	require.NoError(t, err)
}

func Test_StudyStage_Update(t *testing.T) {
	repository := NewStudyStageRepository(db.NewMockProvider())

	// Update Description of StudyStage
	studyStage := db.StudyStage_1()
	studyStage.Description = "edited description StudyStage 1"
	err := repository.Update(&studyStage)

	// Return all StudyStages for comparing
	result2, _ := repository.GetAll()
	studyStages := result2.([]models.StudyStage)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(studyStage.Description, studyStages[0].Description))
}

func Test_StudyStage_Find(t *testing.T) {
	repository := NewStudyStageRepository(db.NewMockProvider())

	// Find StudyStage
	result2, err := repository.Find(db.StudyStage_1())
	studyStages := result2.([]models.StudyStage)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.StudyStage_1().ID, studyStages[0].ID))
}

func Test_StudyStage_GetAll(t *testing.T) {
	repository := NewStudyStageRepository(db.NewMockProvider())

	// Get all studyStages
	result, err := repository.GetAll()
	studyStages := result.([]models.StudyStage)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.StudyStage_1().ID, studyStages[0].ID))
	require.Nil(t, deep.Equal(db.StudyStage_2().ID, studyStages[1].ID))
}

func Test_StudyStage_GetID(t *testing.T) {
	repository := NewStudyStageRepository(db.NewMockProvider())

	// Get by ID
	result, err := repository.GetId(db.StudyStage_1().ID)
	studyStage := result.(*models.StudyStage)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(studyStage.Description, db.StudyStage_1().Description))
}

func Test_StudyStage_DeleteId(t *testing.T) {
	repository := NewStudyStageRepository(db.NewMockProvider())

	// Get length of slice of all studyStages
	result, _ := repository.GetAll()
	afterCreateLength := len(result.([]models.StudyStage))

	// Delete studyStage
	err := repository.DeleteId(db.StudyStage_1().ID)

	result2, _ := repository.GetAll()
	afterDeleteLength := len(result2.([]models.StudyStage))

	require.NoError(t, err)
	require.Nil(t, deep.Equal(afterCreateLength-1, afterDeleteLength))
}
