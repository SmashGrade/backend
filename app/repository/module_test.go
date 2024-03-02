package repository

import (
	"testing"

	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
)

func Test_Module_Create(t *testing.T) {
	repository := NewModuleRepository(db.NewPrefilledMockProvider())

	var err error
	module_1 := db.Module1
	module_1.ID, err = repository.GetNextId()
	require.NoError(t, err)

	_, err = repository.Create(&module_1)

	require.NoError(t, err)
}

func Test_Module_Update(t *testing.T) {
	repository := NewModuleRepository(db.NewPrefilledMockProvider())

	// Update Description of Field
	module := db.Module1
	module.Description = "edited Module 1"
	err := repository.Update(&module)

	// Return all Fields for comparing
	result2, _ := repository.GetAll()
	modules := result2.([]models.Module)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(module.Description, modules[0].Description))
}

func Test_Module_Find(t *testing.T) {
	repository := NewModuleRepository(db.NewPrefilledMockProvider())

	// Find Field
	result2, err := repository.Find(db.Module1)
	modules := result2.([]models.Module)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.Module1.ID, modules[0].ID))
}

func Test_Module_GetAll(t *testing.T) {
	repository := NewModuleRepository(db.NewPrefilledMockProvider())

	// Get all fields
	result, err := repository.GetAll()
	modules := result.([]models.Module)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.Module1.ID, modules[0].ID))
	require.Nil(t, deep.Equal(db.Module2.ID, modules[1].ID))
}

func Test_Module_GetVersioned(t *testing.T) {
	repository := NewModuleRepository(db.NewPrefilledMockProvider())

	// Get by ID
	result, err := repository.GetVersioned(db.Module1.ID, db.Module1.Version)
	module := result.(*models.Module)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(module.Description, db.Module1.Description))
}

func Test_Module_DeleteVersioned(t *testing.T) {
	repository := NewModuleRepository(db.NewPrefilledMockProvider())

	// Get length of slice of all fields
	result, _ := repository.GetAll()
	afterCreateLength := len(result.([]models.Module))

	// Delete field
	err := repository.DeleteVersioned(db.Module1.ID, db.Module1.Version)

	result2, _ := repository.GetAll()
	afterDeleteLength := len(result2.([]models.Module))

	require.NoError(t, err)
	require.Nil(t, deep.Equal(afterCreateLength-1, afterDeleteLength))
}

func Test_Module_GetLatestVersioned(t *testing.T) {
	repository := NewModuleRepository(db.NewPrefilledMockProvider())

	// Get latest Version
	result, err := repository.GetLatestVersioned(db.Module2.ID)
	module := result.(*models.Module)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(module.Version, db.Module2.Version))
}

// check if the version is correctly incremented and tracked
func TestModuleVersionIncrement(t *testing.T) {
	repository := NewModuleRepository(db.NewPrefilledMockProvider())

	all, err := repository.GetAll()
	require.NoError(t, err)

	modules, ok := all.([]models.Module)
	require.True(t, ok, "slice type assertion failed")

	latestVersion, err := repository.GetLatestVersion(modules[0].ID) // TODO: this fails
	require.NoError(t, err)

	nextVersion, err := repository.GetNextVersion(modules[0].ID)
	require.NoError(t, err)

	if nextVersion <= latestVersion {
		t.Fatalf("Next version %v should be greater than current version %v", nextVersion, latestVersion)
	}

	currentEntity, err := repository.GetLatestVersioned(modules[0].ID)
	require.NoError(t, err)

	currentModule, ok := currentEntity.(*models.Module)
	require.True(t, ok, "type assertion failed")

	currentModule.Version = nextVersion

	createdEntity, err := repository.Create(currentModule)
	require.NoError(t, err)

	createdModule, ok := createdEntity.(*models.Module)
	require.True(t, ok, "type assertion failed")

	latestVersion, err = repository.GetLatestVersion(modules[0].ID)
	require.NoError(t, err)

	nextVersion, err = repository.GetNextVersion(modules[0].ID)
	require.NoError(t, err)

	require.Equal(t, currentModule.ID, createdModule.ID)
	require.Equal(t, latestVersion, createdModule.Version)

	if nextVersion <= latestVersion {
		t.Fatalf("Next version %v should be greater than current version %v", nextVersion, latestVersion)
	}
}
