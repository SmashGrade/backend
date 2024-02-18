package repository

import (
	"testing"

	"github.com/SmashGrade/backend/app/config"
	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
)

func Test_Curriculum_Create(t *testing.T) {
	repository := NewCurriculumRepository(db.NewPrefilledMockProvider())

	curriculum_1 := db.Curriculum_1()
	curriculum_1.ID = 0

	_, err := repository.Create(&curriculum_1)

	require.NoError(t, err)
}

func Test_Curriculum_Update(t *testing.T) {
	repository := NewCurriculumRepository(db.NewPrefilledMockProvider())

	// Update Description of Curriculum
	curriculum := db.Curriculum_1()
	curriculum.Description = "edited description Curriculum 1"
	err := repository.Update(&curriculum)

	// Return all Curriculums for comparing
	result2, _ := repository.GetAll()
	curriculums := result2.([]models.Curriculum)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(curriculum.Description, curriculums[0].Description))
}

func Test_Curriculum_Find(t *testing.T) {
	repository := NewCurriculumRepository(db.NewPrefilledMockProvider())

	// Find Curriculum
	result2, err := repository.Find(db.Curriculum_1())
	curriculums := result2.([]models.Curriculum)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.Curriculum_1().ID, curriculums[0].ID))
}

func Test_Curriculum_GetAll(t *testing.T) {
	repository := NewCurriculumRepository(db.NewPrefilledMockProvider())

	// Get all fields
	result, err := repository.GetAll()
	curriculums := result.([]models.Curriculum)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.Curriculum_1().ID, curriculums[0].ID))
	require.Nil(t, deep.Equal(db.Curriculum_2().ID, curriculums[1].ID))
}

func Test_Curriculum_GetID(t *testing.T) {
	repository := NewCurriculumRepository(db.NewPrefilledMockProvider())

	// Get by ID
	result, err := repository.GetId(db.Curriculum_1().ID)
	curriculum := result.(*models.Curriculum)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(curriculum.Description, db.Curriculum_1().Description))
}

func Test_Curriculum_DeleteId(t *testing.T) {
	repository := NewCurriculumRepository(db.NewPrefilledMockProvider())

	// Get length of slice of all curriculum
	result, _ := repository.GetAll()
	afterCreateLength := len(result.([]models.Curriculum))

	// Delete curriculum
	err := repository.DeleteId(db.Curriculum_1().ID)

	result2, _ := repository.GetAll()
	afterDeleteLength := len(result2.([]models.Curriculum))

	require.NoError(t, err)
	require.Nil(t, deep.Equal(afterCreateLength-1, afterDeleteLength))
}

// create two same id curriculum and different times
// create two different id curriculum and same time
func TestCurriculumCreateWithStartyear(t *testing.T) {
	repository := NewCurriculumRepository(db.NewProvider(config.NewAPIConfig()))

	/*timeCheck := time.Now()

	curriculum1 := &models.Curriculum{}
	curriculum1.ID = 15
	curriculum1.StartValidity = timeCheck

	_, err := repository.Create(curriculum1)
	require.NoError(t, err)

	curriculum1.StartValidity = timeCheck.Add(time.Hour)

	_, err = repository.Create(curriculum1)
	require.NoError(t, err)*/

	var err error
	curriculum1 := &models.Curriculum{}
	curriculum1.ID, err = repository.GetLatestId()
	require.NoError(t, err)
	curriculum1.ID += 1

	_, err = repository.Create(curriculum1)
	require.NoError(t, err)

	curriculum1.ID, err = repository.GetLatestId()
	require.NoError(t, err)
	curriculum1.ID += 1

	_, err = repository.Create(curriculum1)
	require.NoError(t, err)
}
