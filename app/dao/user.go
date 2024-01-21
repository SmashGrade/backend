package dao

import (
	"github.com/SmashGrade/backend/app/api/v1/schemas"
	"github.com/SmashGrade/backend/app/entity"
)

func (db *Database) ListUsers(usersRes *[]schemas.User) error {
	err := db.listUsers(usersRes)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) PostUser(userReq *schemas.User) error {
	var user entity.User

	// User Id has to be set to 0 so the User can not set it
	userReq.Id = 0

	// The Role of the User has to be extracted and added after it got parsed
	rolesRef := userReq.Roles
	var roles []*entity.Role
	db.ListRoles(&roles, rolesRef)

	err := ParseSchemaToEntity(&userReq, &user)
	if err != nil {
		return err
	}

	user.Roles = roles

	db.Db.Create(&user)
	return nil
}

func (db *Database) GetUser(user *schemas.User, id uint) error {
	err := db.getUser(user, id)
	if err != nil {
		return err
	}

	return nil
}

// creates a new user and returns id
func (db *Database) CreateUser(user *schemas.User) (uint, error) {
	// TODO: is the user.CurriculumStartYear = ClassStartyear???

	newUser := &entity.User{
		Name:  user.Name,
		Email: user.Email,
	}

	// map roles by searching them by name
	newRoles := make([]*entity.Role, len(user.Roles))
	for _, role := range user.Roles {
		newRole := &entity.Role{}
		err := db.GetRoleByName(newRole, role)
		if err != nil {
			return 0, err
		}
		newRoles = append(newRoles, newRole)
	}

	newUser.Roles = newRoles

	err := db.Db.Create(newUser).Error
	if err != nil {
		return 0, err
	}

	return newUser.ID, nil
}
