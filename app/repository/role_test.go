package repository

import (
	"testing"

	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
)

func Test_Role_Create(t *testing.T) {
	repository := NewRoleRepository(db.NewPrefilledMockProvider())

	role_1 := db.RoleStudent
	role_1.ID = 0

	_, err := repository.Create(&role_1)

	require.NoError(t, err)
}

func Test_Role_Update(t *testing.T) {
	repository := NewRoleRepository(db.NewPrefilledMockProvider())

	// Update Description of Role
	role := db.RoleStudent
	role.Description = "edited description Role 1"
	err := repository.Update(&role)

	require.NoError(t, err)

	// Get the updated role back from db
	result2, err := repository.GetId(role.ID)
	updatedRole := result2.(*models.Role)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(role.Description, updatedRole.Description))
}

func Test_Role_Find(t *testing.T) {
	repository := NewRoleRepository(db.NewPrefilledMockProvider())

	// Find Role
	result2, err := repository.Find(db.RoleStudent)
	roles := result2.([]models.Role)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.RoleStudent.ID, roles[0].ID))
}

func Test_Role_GetAll(t *testing.T) {
	repository := NewRoleRepository(db.NewPrefilledMockProvider())

	// Get all roles
	result, err := repository.GetAll()
	roles := result.([]models.Role)

	require.NoError(t, err)
	require.NotEmpty(t, roles)
}

func Test_Role_GetID(t *testing.T) {
	repository := NewRoleRepository(db.NewPrefilledMockProvider())

	// Get by ID
	result, err := repository.GetId(db.RoleStudent.ID)
	role := result.(*models.Role)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(role.Description, db.RoleStudent.Description))
}

func Test_Role_DeleteId(t *testing.T) {
	repository := NewRoleRepository(db.NewPrefilledMockProvider())

	// Get length of slice of all roles
	result, _ := repository.GetAll()
	afterCreateLength := len(result.([]models.Role))

	// Delete role
	err := repository.DeleteId(db.RoleStudent.ID)

	result2, _ := repository.GetAll()
	afterDeleteLength := len(result2.([]models.Role))

	require.NoError(t, err)
	require.Nil(t, deep.Equal(afterCreateLength-1, afterDeleteLength))
}
