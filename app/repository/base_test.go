package repository

import (
	"fmt"
	"testing"
	"time"

	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {

	tests := []struct {
		name       string
		entity     any
		repository Repository
		want       string
	}{
		{
			name: "ExamEvaluation",
			entity: models.ExamEvaluation{
				RegisteredBy:   db.User5,
				SelectedCourse: db.SelectedCourse2,
				Exam:           db.Exam2,
				OriginalValue:  "6.0",
				EntryDate:      time.Date(2023, time.September, 1, 12, 0, 0, 0, time.UTC),
			},
			repository: NewExamEvaluationRepository(db.NewPrefilledMockProvider()),
			want:       "200",
		},
		{
			name: "Conversion",
			entity: models.Conversion{
				ExamEvaluation: db.ExamEvaluation3,
				Gradetype:      db.GradeType1,
				Value:          4.8,
			},
			repository: NewConversionRepository(db.NewPrefilledMockProvider()),
			want:       "lulu",
		},
	}
	for _, testData := range tests {
		var err error

		switch testData.name {
		case "ExamEvaluation":
			entity := testData.entity.(models.ExamEvaluation)
			_, err = testData.repository.Create(&entity)
		case "Conversion":
			entity := testData.entity.(models.Conversion)
			_, err = testData.repository.Create(&entity)
		}

		require.NoError(t, err)
	}

}

func TestGetAll(t *testing.T) {
	tests := []struct {
		name       string
		repository Repository
		types      any
		want       string
	}{
		{
			name:       "ExamEvaluation",
			repository: NewExamEvaluationRepository(db.NewPrefilledMockProvider()),
			types:      models.ExamEvaluation{},
			want:       "200",
		},
	}

	for _, testData := range tests {

		res, err := testData.repository.GetAll()
		entities := res.([]models.ExamEvaluation)
		fmt.Print(entities)

		require.NoError(t, err)
	}
}

func getAllOrError[outputModel any](repo Repository) (outputSlice []outputModel) {
	internalSlice, internalErr := repo.GetAll()
	if internalErr != nil {
		//err = e.NewDaoDbError()
		return
	}

	outputSlice = internalSlice.([]outputModel)
	return
}
