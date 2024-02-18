package repository

import (
	"testing"

	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
)

func Test_Role_Create(t *testing.T) {
	repository := NewRoleRepository(db.NewMockProvider())

	role_1 := db.Role_1()
	role_1.ID = 0

	_, err := repository.Create(&role_1)

	require.NoError(t, err)
}

func Test_Role_Update(t *testing.T) {
	repository := NewRoleRepository(db.NewMockProvider())

	// Update Description of Role
	role := db.Role_1()
	role.Description = "edited description Role 1"
	err := repository.Update(&role)

	// Return all Roles for comparing
	result2, _ := repository.GetAll()
	roles := result2.([]models.Role)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(role.Description, roles[0].Description))
}

func Test_Role_Find(t *testing.T) {
	repository := NewRoleRepository(db.NewMockProvider())

	// Find Role
	result2, err := repository.Find(db.Role_1())
	roles := result2.([]models.Role)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.Role_1().ID, roles[0].ID))
}

func Test_Role_GetAll(t *testing.T) {
	repository := NewRoleRepository(db.NewMockProvider())

	// Get all roles
	result, err := repository.GetAll()
	roles := result.([]models.Role)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.Role_1().ID, roles[0].ID))
	require.Nil(t, deep.Equal(db.Role_2().ID, roles[1].ID))
}

func Test_Role_GetID(t *testing.T) {
	repository := NewRoleRepository(db.NewMockProvider())

	// Get by ID
	result, err := repository.GetId(db.Role_1().ID)
	role := result.(*models.Role)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(role.Description, db.Role_1().Description))
}

func Test_Role_DeleteId(t *testing.T) {
	repository := NewRoleRepository(db.NewMockProvider())

	// Get length of slice of all roles
	result, _ := repository.GetAll()
	afterCreateLength := len(result.([]models.Role))

	// Delete role
	err := repository.DeleteId(db.Role_1().ID)

	result2, _ := repository.GetAll()
	afterDeleteLength := len(result2.([]models.Role))

	require.NoError(t, err)
	require.Nil(t, deep.Equal(afterCreateLength-1, afterDeleteLength))
}
