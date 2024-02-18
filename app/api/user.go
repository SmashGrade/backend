package api

import (
	"github.com/SmashGrade/backend/app/dao"
	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/repository"
	"github.com/labstack/echo/v4"
)

// Handles the requests for the user resource
type UserController struct {
	*BaseController
	Dao *dao.UserDao
}

// Constructor for UserController
func NewUserController(provider db.Provider) *UserController {
	return &UserController{
		BaseController: NewBaseController(provider),
		Dao:            dao.NewUserDao(repository.NewUserRepository(provider)),
	}
}

// @Summary		Get the current user
// @Description	Get the current user (from Bearer token)
// @Tags			users
// @Produce		json
// @Success		200	{object}	models.User
// @Failure		401	{object}	error.ApiError
// @Failure		403	{object}	error.ApiError
// @Failure		500	{object}	error.ApiError
// @Router			/user [get]
// @Security		Bearer
func (c *UserController) User(ctx echo.Context) error {
	user, err := c.GetUser(ctx)
	// Return error if error occurred
	if err != nil {
		return err
	}
	// Return the user
	return c.Yeet(ctx, user)
}

// register all output endpoints to router
func RegisterV1User(g *echo.Group, u *UserController) {
	g.GET("/user", u.User)
}
