package repository

import (
	"testing"

	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
)

func Test_ExamEvaluation_Create(t *testing.T) {
	repository := NewExamEvaluationRepository(db.NewPrefilledMockProvider())

	examEvaluation_1 := db.ExamEvaluation_1()
	examEvaluation_1.ID = 0

	_, err := repository.Create(&examEvaluation_1)

	require.NoError(t, err)
}

func Test_ExamEvaluation_Update(t *testing.T) {
	repository := NewExamEvaluationRepository(db.NewPrefilledMockProvider())

	// Update OriginalValue of ExamEvaluation
	examEvaluation := db.ExamEvaluation_1()
	examEvaluation.OriginalValue = "edited description ExamEvaluation 1"
	err := repository.Update(&examEvaluation)

	// Return all ExamEvaluations for comparing
	result2, _ := repository.GetAll()
	examEvaluations := result2.([]models.ExamEvaluation)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(examEvaluation.OriginalValue, examEvaluations[0].OriginalValue))
}

func Test_ExamEvaluation_Find(t *testing.T) {
	repository := NewExamEvaluationRepository(db.NewPrefilledMockProvider())

	// Find ExamEvaluation
	result2, err := repository.Find(db.ExamEvaluation_1())
	examEvaluations := result2.([]models.ExamEvaluation)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.ExamEvaluation_1().ID, examEvaluations[0].ID))
}

func Test_ExamEvaluation_GetAll(t *testing.T) {
	repository := NewExamEvaluationRepository(db.NewPrefilledMockProvider())

	// Get all examEvaluations
	result, err := repository.GetAll()
	examEvaluations := result.([]models.ExamEvaluation)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.ExamEvaluation_1().ID, examEvaluations[0].ID))
	require.Nil(t, deep.Equal(db.ExamEvaluation_2().ID, examEvaluations[1].ID))
}

func Test_ExamEvaluation_GetID(t *testing.T) {
	repository := NewExamEvaluationRepository(db.NewPrefilledMockProvider())

	// Get by ID
	result, err := repository.GetId(db.ExamEvaluation_1().ID)
	examEvaluation := result.(*models.ExamEvaluation)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(examEvaluation.OriginalValue, db.ExamEvaluation_1().OriginalValue))
}

func Test_ExamEvaluation_DeleteId(t *testing.T) {
	repository := NewExamEvaluationRepository(db.NewPrefilledMockProvider())

	// Get length of slice of all examEvaluations
	result, _ := repository.GetAll()
	afterCreateLength := len(result.([]models.ExamEvaluation))

	// Delete examEvaluation
	err := repository.DeleteId(db.ExamEvaluation_1().ID)

	result2, _ := repository.GetAll()
	afterDeleteLength := len(result2.([]models.ExamEvaluation))

	require.NoError(t, err)
	require.Nil(t, deep.Equal(afterCreateLength-1, afterDeleteLength))
}
