package repository

import (
	"testing"

	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
)

func Test_Module_Create(t *testing.T) {
	repository := NewModuleRepository(db.NewMockProvider())

	module_1 := db.Module_1()
	module_1.ID = 0

	_, err := repository.Create(&module_1)

	require.NoError(t, err)
}

func Test_Module_Update(t *testing.T) {
	repository := NewModuleRepository(db.NewMockProvider())

	// Update Description of Field
	module := db.Module_1()
	module.Description = "edited Module 1"
	err := repository.Update(&module)

	// Return all Fields for comparing
	result2, _ := repository.GetAll()
	modules := result2.([]models.Module)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(module.Description, modules[0].Description))
}

func Test_Module_Find(t *testing.T) {
	repository := NewModuleRepository(db.NewMockProvider())

	// Find Field
	result2, err := repository.Find(db.Module_1())
	modules := result2.([]models.Module)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.Module_1().ID, modules[0].ID))
}

func Test_Module_GetAll(t *testing.T) {
	repository := NewModuleRepository(db.NewMockProvider())

	// Get all fields
	result, err := repository.GetAll()
	modules := result.([]models.Module)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.Module_1().ID, modules[0].ID))
	require.Nil(t, deep.Equal(db.Module_2_1().ID, modules[1].ID))
}

func Test_Module_GetVersioned(t *testing.T) {
	repository := NewModuleRepository(db.NewMockProvider())

	// Get by ID
	result, err := repository.GetVersioned(db.Module_1().ID, db.Module_1().Version)
	module := result.(*models.Module)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(module.Description, db.Module_1().Description))
}

func Test_Module_DeleteVersioned(t *testing.T) {
	repository := NewModuleRepository(db.NewMockProvider())

	// Get length of slice of all fields
	result, _ := repository.GetAll()
	afterCreateLength := len(result.([]models.Module))

	// Delete field
	err := repository.DeleteVersioned(db.Module_1().ID, db.Module_1().Version)

	result2, _ := repository.GetAll()
	afterDeleteLength := len(result2.([]models.Module))

	require.NoError(t, err)
	require.Nil(t, deep.Equal(afterCreateLength-1, afterDeleteLength))
}

func Test_Module_GetLatestVersioned(t *testing.T) {
	repository := NewModuleRepository(db.NewMockProvider())

	// Get latest Version
	result, err := repository.GetLatestVersioned(db.Module_2_1().ID)
	module := result.(*models.Module)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(module.Version, db.Module_2_2().Version))
}
