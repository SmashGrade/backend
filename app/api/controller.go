package api

import (
	"strconv"

	"github.com/SmashGrade/backend/app/dao"
	"github.com/SmashGrade/backend/app/db"
	e "github.com/SmashGrade/backend/app/error"
	"github.com/SmashGrade/backend/app/models"
	"github.com/SmashGrade/backend/app/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// Controller is the base controller for all API controllers
type BaseController struct {
	Provider db.Provider
	UserDao  *dao.UserDao
}

// Constructor for BaseController
func NewBaseController(provider db.Provider) *BaseController {
	return &BaseController{
		Provider: provider,
		UserDao:  dao.NewUserDao(repository.NewUserRepository(provider)),
	}
}

// Yeets the data out of the controller to the client
func (c *BaseController) Yeet(ctx echo.Context, data any) error {
	return ctx.JSON(200, map[string]any{"data": data})
}

// Gets the parameter from the request
func (c *BaseController) GetPathParam(ctx echo.Context, param string) string {
	return ctx.Param(param)
}

// Gets the parameter from the request and converts it to an integer
// Returns -1 if the conversion fails
func (c *BaseController) GetPathParamInt(ctx echo.Context, param string) int {
	res, err := strconv.Atoi(c.GetPathParam(ctx, param))
	// Check if conversion failed
	if err != nil {
		return -1
	}
	// Check if the result is negative
	// This would not be able to be converted to a uint
	if res < 0 {
		return -1
	}
	// Return value
	return res
}

// Retrieves the user from the requests bearer token
// Ensures that the user is authenticated and exists in the database
// Handles the unauthorized and forbidden errors
func (c *BaseController) GetUser(ctx echo.Context) (*models.User, error) {
	userRaw := ctx.Get("user")
	// Middleware does not have a user key, so we return unauthorized
	if userRaw == nil {
		ctx.Logger().Info("Authorized endpoint called without a bearer token. Request denied.")
		return nil, e.NewUnauthorizedError()
	}
	// Check if the user key is a valid jwt token
	user, ok := userRaw.(*jwt.Token)
	if !ok {
		ctx.Logger().Info("Authorized endpoint called without a valid bearer token. Request denied.")
		return nil, e.NewUnauthorizedError()
	}
	// Check if the user is valid
	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		ctx.Logger().Info("Authorized endpoint called without valid claims. Request denied.")
		return nil, e.NewUnauthorizedError()
	}
	// Print claims for debugging
	ctx.Logger().Info("Claims: ", claims)

	// TODO: Finish this function
	return &models.User{}, nil
}
