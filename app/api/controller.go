package api

import (
	"net/mail"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/SmashGrade/backend/app/config"
	"github.com/SmashGrade/backend/app/dao"
	"github.com/SmashGrade/backend/app/db"
	e "github.com/SmashGrade/backend/app/error"
	"github.com/SmashGrade/backend/app/models"
	"github.com/SmashGrade/backend/app/repository"
	"github.com/goccy/go-json"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// Controller is the base controller for all API controllers
type BaseController struct {
	Provider db.Provider
	UserDao  *dao.UserDao
}

type TokenClaim struct {
	Email string   `json:"preferred_username"`
	Name  string   `json:"name"`
	Roles []string `json:"roles"`
	jwt.RegisteredClaims
}

// Constructor for BaseController
func NewBaseController(provider db.Provider) *BaseController {
	return &BaseController{
		Provider: provider,
		UserDao:  dao.NewUserDao(repository.NewUserRepository(provider), repository.NewRoleRepository(provider)),
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
func (c *BaseController) GetPathParamUint(ctx echo.Context, param string) (uint, error) {
	res, err := strconv.Atoi(c.GetPathParam(ctx, param))
	// Check if conversion failed
	if err != nil {
		return 0, err
	}
	// Check if the result is negative
	// This would not be able to be converted to a uint
	if res < 0 {
		return 0, err
	}
	// Return value
	return uint(res), nil
}

// GetPathParamTime retrieves the value of a path parameter from the given context and parses it as a time.Time value.
// It expects the path parameter to be in the format "30.07.2024".
// If the parsing is successful, it returns the parsed time.Time value.
// If the parsing fails, it returns an error.
func (c *BaseController) GetPathParamTime(ctx echo.Context, param string) (time.Time, error) {
	// Get dateString from Parameter
	dateString := c.GetPathParam(ctx, param)

	// parse string to time.Time
	dateTime, err := dao.ParseTime(dateString, "02.01.2006")
	if err != nil {
		return time.Time{}, err
	}

	return dateTime, nil
}

// Retrieves the user from the requests bearer token
// Ensures that the user is authenticated and exists in the database
// Handles the unauthorized and forbidden errors
func (c *BaseController) GetUser(ctx echo.Context) (*models.User, *e.ApiError) {
	userRaw := ctx.Get("user")
	// Middleware does not have a user key, so we return unauthorized
	if userRaw == nil {
		ctx.Logger().Error("Authorized endpoint called without a bearer token. Request denied.")
		return nil, e.NewUnauthorizedError()
	}
	// Check if the user key is a valid jwt token
	user, ok := userRaw.(*jwt.Token)
	if !ok {
		ctx.Logger().Error("Authorized endpoint called without a valid bearer token. Request denied.")
		return nil, e.NewUnauthorizedError()
	}
	// Check if the user is valid
	// This is a workaround to get the claims from the token as the default map can not be converted to a struct
	marshalledClaims, err := json.Marshal(user.Claims)
	if err != nil {
		ctx.Logger().Error("Authorized endpoint called without valid claims. Request denied.")
		return nil, e.NewUnauthorizedError()
	}
	// Build the clains as struct from the marshalled claims
	claims := TokenClaim{}
	err = json.Unmarshal(marshalledClaims, &claims)
	if err != nil {
		ctx.Logger().Error("Authorized endpoint called without valid claims. Request denied.")
		return nil, e.NewUnauthorizedError()
	}

	// Create a list of roles from the claims
	userRoles := make([]*models.Role, 0)
	// Add roles from the claims to the user
	for _, claimRole := range claims.Roles {
		role, err := c.UserDao.GetRoleByClaim(claimRole)
		if err != nil {
			ctx.Logger().Warnf("Encountered a role that does not exist in the database: %s for user %s", claimRole, claims.Email)
		} else {
			userRoles = append(userRoles, role)
		}
	}

	// Check if the email is valid
	_, emailInvalidErr := mail.ParseAddress(claims.Email)
	if emailInvalidErr != nil {
		ctx.Logger().Errorf("Authorized endpoint called with invalid email address: %s. Request denied.", claims.Email)
		return nil, e.NewUnauthorizedError()
	}

	// Check if the email address of the user is allowed to access the application
	emailDomain := claims.Email[strings.Index(claims.Email, "@")+1:]
	if !slices.Contains(c.Provider.Config().AllowedDomains, emailDomain) {
		ctx.Logger().Errorf("Authorized endpoint called with unauthorized email address: %s. Request denied.", claims.Email)
		return nil, e.NewUnauthorizedError()
	}

	// Create the user object from the claims
	userEntity := models.User{
		Email: claims.Email,
		Name:  claims.Name,
		Roles: userRoles,
	}

	// Ensure that the database contains the user and that the user is updated based on the claims
	crudUser, crudErr := c.UserDao.CreateOrUpdateByEmail(userEntity)
	if crudErr != nil {
		ctx.Logger().Error("Error creating or updating user in database. Request denied.")
		return nil, e.NewUnauthorizedError()
	}

	// Return the user
	return crudUser, nil
}

// check if a user has a role by roleId
// returns nil if the claim is valid
func (c *BaseController) CheckUserRole(roleId uint, ctx echo.Context) *e.ApiError {

	// Check if authentication has been disabled
	// Return nil no matter what claims are present
	if !c.Provider.Config().AuthConfig.Enabled {
		return nil
	}

	var requiredRole *config.RoleConfig = nil
	for i := range c.Provider.Config().Roles {
		if c.Provider.Config().Roles[i].Id == roleId {
			requiredRole = &c.Provider.Config().Roles[i]
		}
	}
	if requiredRole == nil {
		return e.NewDaoReferenceIdError("role", roleId)
	}

	user, err := c.GetUser(ctx)
	if err != nil {
		return err
	}

	if !user.HasRole(roleId) {
		return e.NewClaimMissingError(requiredRole.ClaimName)
	}

	return nil
}

// Allows the user to access the endpoint if the user has any role
// returns nil if the claim is valid
func (c *BaseController) CheckUserAnyRole(ctx echo.Context) *e.ApiError {
	// Check if authentication has been disabled
	// Return nil no matter what claims are present
	if !c.Provider.Config().AuthConfig.Enabled {
		return nil
	}

	user, err := c.GetUser(ctx)
	if err != nil {
		return err
	}

	if !user.HasAnyRole() {
		return e.NewClaimMissingError("any")
	}

	return nil

}

// CheckUserRoles loops through multiple roles to check if any is correct.
// If a correct role is found, it returns nil; otherwise, it returns an error.
func (c *BaseController) CheckUserRoles(roleIDs []uint, ctx echo.Context) *e.ApiError {
	if len(roleIDs) == 0 {
		return e.NewDaoReferenceIdError("role", 0)
	}

	var lastError *e.ApiError
	for _, roleID := range roleIDs {
		if err := c.CheckUserRole(roleID, ctx); err == nil {
			return nil
		} else {
			lastError = err
		}
	}
	return lastError
}
