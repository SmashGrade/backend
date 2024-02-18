package repository

import (
	"testing"

	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
)

func Test_User_Create(t *testing.T) {
	repository := NewUserRepository(db.NewMockProvider())

	user_1 := db.User_1()
	user_1.ID = 0

	_, err := repository.Create(&user_1)

	require.NoError(t, err)
}

func Test_User_Update(t *testing.T) {
	repository := NewUserRepository(db.NewMockProvider())

	// Update Name of User
	user := db.User_1()
	user.Name = "edited description User 1"
	err := repository.Update(&user)

	// Return all Users for comparing
	result2, _ := repository.GetAll()
	users := result2.([]models.User)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(user.Name, users[0].Name))
}

func Test_User_Find(t *testing.T) {
	repository := NewUserRepository(db.NewMockProvider())

	// Find User
	result2, err := repository.Find(db.User_1())
	users := result2.([]models.User)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.User_1().ID, users[0].ID))
}

func Test_User_GetAll(t *testing.T) {
	repository := NewUserRepository(db.NewMockProvider())

	// Get all users
	result, err := repository.GetAll()
	users := result.([]models.User)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(db.User_1().ID, users[0].ID))
	require.Nil(t, deep.Equal(db.User_2().ID, users[1].ID))
}

func Test_User_GetID(t *testing.T) {
	repository := NewUserRepository(db.NewMockProvider())

	// Get by ID
	result, err := repository.GetId(db.User_1().ID)
	user := result.(*models.User)

	require.NoError(t, err)
	require.Nil(t, deep.Equal(user.Name, db.User_1().Name))
}

func Test_User_DeleteId(t *testing.T) {
	repository := NewUserRepository(db.NewMockProvider())

	// Get length of slice of all users
	result, _ := repository.GetAll()
	afterCreateLength := len(result.([]models.User))

	// Delete user
	err := repository.DeleteId(db.User_1().ID)

	result2, _ := repository.GetAll()
	afterDeleteLength := len(result2.([]models.User))

	require.NoError(t, err)
	require.Nil(t, deep.Equal(afterCreateLength-1, afterDeleteLength))
}
