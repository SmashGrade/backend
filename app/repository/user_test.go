package repository

import (
	"testing"

	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/models"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
)

func TestUserCreate(t *testing.T) {
	repository := NewUserRepository(db.NewMockProvider())

	_, err := repository.Create(&models.User{
		Name:  "Max Müller",
		Email: "max.mueller@hftm.ch",
	})

	require.NoError(t, err)
}

func TestUserUpdate(t *testing.T) {
	repository := NewUserRepository(db.NewPrefilledMockProvider())

	checkEmail := "rafael.stauffer@gmx.ch"

	user, err := repository.GetId(1)
	require.NoError(t, err)

	castedUser := user.(*models.User)
	castedUser.Email = checkEmail

	err = repository.Update(castedUser)
	require.NoError(t, err)

	// Return all Users for comparing
	secondUser, err := repository.GetId(1)
	require.NoError(t, err)

	castedUser = secondUser.(*models.User)

	require.NoError(t, err)
	require.Equal(t, checkEmail, castedUser.Email)
}

func TestUserFindByName(t *testing.T) {
	repository := NewUserRepository(db.NewPrefilledMockProvider())

	checkName := "Rafael Stauffer"

	// Find User
	result2, err := repository.Find(&models.User{
		Name: checkName,
	})
	users := result2.([]models.User)

	require.NoError(t, err)

	require.Equal(t, checkName, users[0].Name)
}

func TestUserFindByEmail(t *testing.T) {
	repository := NewUserRepository(db.NewPrefilledMockProvider())

	checkEmail := "rafael.stauffer@hftm.ch"

	// Find User
	result2, err := repository.Find(&models.User{
		Email: checkEmail,
	})
	users := result2.([]models.User)

	require.NoError(t, err)

	require.Equal(t, checkEmail, users[0].Email)
}

func TestUserGetAll(t *testing.T) {
	repository := NewUserRepository(db.NewPrefilledMockProvider())

	// Get all users
	result, err := repository.GetAll()
	users := result.([]models.User)

	require.NoError(t, err)
	if len(users) < 1 {
		t.Fatal("No users returned")
	}
}

func TestUserGetById(t *testing.T) {
	repository := NewUserRepository(db.NewPrefilledMockProvider())

	user, err := repository.GetId(1)

	require.NoError(t, err)
	require.NotNil(t, user)
}

func TestUserDeleteById(t *testing.T) {
	repository := NewUserRepository(db.NewPrefilledMockProvider())

	// Get length of slice of all users
	result, _ := repository.GetAll()
	afterCreateLength := len(result.([]models.User))

	// Delete user
	err := repository.DeleteId(1)

	result2, _ := repository.GetAll()
	afterDeleteLength := len(result2.([]models.User))

	require.NoError(t, err)
	require.Nil(t, deep.Equal(afterCreateLength-1, afterDeleteLength))
}
