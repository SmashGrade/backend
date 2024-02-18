package repository

import (
	"testing"

	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
)

func Test_Exam_Create(t *testing.T) {
	repository := NewExamRepository(db.NewPrefilledMockProvider())

	exam_1 := db.Exam_1()
	exam_1.ID = 0

	_, err := repository.Create(&exam_1)

	require.NoError(t, err)
}

func Test_Exam_Update(t *testing.T) {
	repository := NewExamRepository(db.NewPrefilledMockProvider())

	// Update Description of Exam
	exam := db.Exam_1()
	exam.Description = "edited description Exam 1"
	err := repository.Update(&exam)

	// Return all Exams for comparing
	result2, _ := repository.GetAll()
	exams := result2.([]models.Exam)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(exam.Description, exams[0].Description))
}

func Test_Exam_Find(t *testing.T) {
	repository := NewExamRepository(db.NewPrefilledMockProvider())

	// Find Exam
	result2, err := repository.Find(db.Exam_1())
	exams := result2.([]models.Exam)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.Exam_1().ID, exams[0].ID))
}

func Test_Exam_GetAll(t *testing.T) {
	repository := NewExamRepository(db.NewPrefilledMockProvider())

	// Get all exams
	result, err := repository.GetAll()
	exams := result.([]models.Exam)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.Exam_1().ID, exams[0].ID))
	require.Nil(t, deep.Equal(db.Exam_2().ID, exams[1].ID))
}

func Test_Exam_GetID(t *testing.T) {
	repository := NewExamRepository(db.NewPrefilledMockProvider())

	// Get by ID
	result, err := repository.GetId(db.Exam_1().ID)
	exam := result.(*models.Exam)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(exam.Description, db.Exam_1().Description))
}

func Test_Exam_DeleteId(t *testing.T) {
	repository := NewExamRepository(db.NewPrefilledMockProvider())

	// Get length of slice of all exams
	result, _ := repository.GetAll()
	afterCreateLength := len(result.([]models.Exam))

	// Delete exams
	err := repository.DeleteId(db.Exam_1().ID)

	result2, _ := repository.GetAll()
	afterDeleteLength := len(result2.([]models.Exam))

	require.NoError(t, err)
	require.Nil(t, deep.Equal(afterCreateLength-1, afterDeleteLength))
}
