package repository

import (
	"testing"

	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
)

func Test_Course_Create(t *testing.T) {
	repository := NewCourseRepository(db.NewPrefilledMockProvider())
	//repository := NewCourseRepository(db.NewProvider(config.NewAPIConfig()))

	var err error

	course_1 := db.Course1
	course_1.ID, err = repository.GetNextId()
	require.NoError(t, err)

	_, err = repository.Create(&course_1)

	require.NoError(t, err)
}

func Test_Course_Update(t *testing.T) {
	repository := NewCourseRepository(db.NewPrefilledMockProvider())

	// Update Description of Field
	course := db.Course1
	course.Description = "edited Course 1"
	err := repository.Update(&course)

	// Return all Fields for comparing
	result2, _ := repository.GetAll()
	courses := result2.([]models.Course)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(course.Description, courses[0].Description))
}

func Test_Course_Find(t *testing.T) {
	repository := NewCourseRepository(db.NewPrefilledMockProvider())

	// Find Field
	result2, err := repository.Find(db.Course1)
	courses := result2.([]models.Course)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.Course1.ID, courses[0].ID))
}

func Test_Course_GetAll(t *testing.T) {
	repository := NewCourseRepository(db.NewPrefilledMockProvider())

	// Get all fields
	result, err := repository.GetAll()
	courses := result.([]models.Course)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.Course1.ID, courses[0].ID))
	require.Nil(t, deep.Equal(db.Course2.ID, courses[1].ID))
}

func Test_Course_GetVersioned(t *testing.T) {
	repository := NewCourseRepository(db.NewPrefilledMockProvider())

	// Get by ID
	result, err := repository.GetVersioned(db.Course1.ID, db.Course1.Version)
	course := result.(*models.Course)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(course.Description, db.Course1.Description))
}

func Test_Course_DeleteVersioned(t *testing.T) {
	repository := NewCourseRepository(db.NewPrefilledMockProvider())

	// Get length of slice of all fields
	result, _ := repository.GetAll()
	afterCreateLength := len(result.([]models.Course))

	// Delete field
	err := repository.DeleteVersioned(db.Course1.ID, db.Course1.Version)

	result2, _ := repository.GetAll()
	afterDeleteLength := len(result2.([]models.Course))

	require.NoError(t, err)
	require.Nil(t, deep.Equal(afterCreateLength-1, afterDeleteLength))
}

func Test_Course_GetLatestVersioned(t *testing.T) {
	repository := NewCourseRepository(db.NewPrefilledMockProvider())

	// Get latest Version
	result, err := repository.GetLatestVersioned(db.Course2.ID)
	course := result.(*models.Course)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(course.Version, db.Course3.Version))
}
