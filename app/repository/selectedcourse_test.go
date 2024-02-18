package repository

import (
	"testing"
	"time"

	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
)

func Test_SelectedCourse_Create(t *testing.T) {
	repository := NewSelectedCourseRepository(db.NewMockProvider())

	selectedCourse_1 := db.SelectedCourse_1()
	selectedCourse_1.UserID = 2

	_, err := repository.Create(&selectedCourse_1)

	require.NoError(t, err)
}

func Test_SelectedCourse_Update(t *testing.T) {
	repository := NewSelectedCourseRepository(db.NewMockProvider())

	// Update Description of SelectedCourse
	selectedCourse := db.SelectedCourse_1()
	selectedCourse.Dispensed = true
	err := repository.Update(&selectedCourse)

	// Return all SelectedCourses for comparing
	result2, _ := repository.GetAll()
	selectedCourses := result2.([]models.SelectedCourse)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(selectedCourse.Dispensed, selectedCourses[0].Dispensed))
}

func Test_SelectedCourse_Find(t *testing.T) {
	repository := NewSelectedCourseRepository(db.NewMockProvider())

	// Find SelectedCourse
	result2, err := repository.Find(db.SelectedCourse_1())
	selectedCourses := result2.([]models.SelectedCourse)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.SelectedCourse_1().UserID, selectedCourses[0].UserID))
}

func Test_SelectedCourse_GetAll(t *testing.T) {
	repository := NewSelectedCourseRepository(db.NewMockProvider())

	// Get all selectedCourses
	result, err := repository.GetAll()
	selectedCourses := result.([]models.SelectedCourse)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.SelectedCourse_1().UserID, selectedCourses[0].UserID))
	require.Nil(t, deep.Equal(db.SelectedCourse_2().UserID, selectedCourses[1].UserID))
}

func Test_SelectedCourse_GetSelectedCourse(t *testing.T) {
	repository := NewSelectedCourseRepository(db.NewMockProvider())

	result, err := repository.GetSelectedCourse(1, 1, 1, time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC))

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.SelectedCourse_1().Dispensed, result.Dispensed))
}

func Test_SelectedCourse_DeleteSelectedCourse(t *testing.T) {
	repository := NewSelectedCourseRepository(db.NewMockProvider())

	// Get length of slice of all selectedcourse
	result, _ := repository.GetAll()
	afterCreateLength := len(result.([]models.SelectedCourse))

	// Delete selectedcourse
	err := repository.DeleteSelectedCourse(1, 1, 1, time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC))

	result2, _ := repository.GetAll()
	afterDeleteLength := len(result2.([]models.SelectedCourse))

	require.NoError(t, err)
	require.Nil(t, deep.Equal(afterCreateLength-1, afterDeleteLength))
}
