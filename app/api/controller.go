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

type TokenClaim struct {
	Email string   `json:"email"`
	Name  string   `json:"name"`
	Roles []string `json:"roles"`
	*jwt.MapClaims
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
func (c *BaseController) GetUser(ctx echo.Context) (*models.User, *e.ApiError) {
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
	claims, ok := user.Claims.(TokenClaim)
	if !ok {
		ctx.Logger().Info("Authorized endpoint called without valid claims. Request denied.")
		return nil, e.NewUnauthorizedError()
	}

	print(claims.Email, claims.Name, claims.Roles)

	/*
	 *Claims: map[acct:1 aio:AYQAe/8WAAAANuzCcU/Fuw3PdUijspv08VrJY1CKopiZXI8nl35szPB1Nu0paNS4UhTcHGEu3pZ0uCkBewp44kVv3N3HvzYgBBusbj67dOT0FFi8h19OljJOyDppvSVcXY/lCd+Uw4McPGH202G12BkFe9RDO4fUCur91+FOiuG37JpZ8k/5qLM= aud:72acf4df-78f6-4e6f-81c6-f5aa1efa8ebc auth_time:1.706943744e+09 email:joshua.lehmann@hftm.ch exp:1.708774097e+09 iat:1.708770197e+09 idp:https://sts.windows.net/7a09aace-3641-41b0-993d-3729201228b3/ ipaddr:178.197.219.87 iss:https://login.microsoftonline.com/744b66c4-2df7-4756-905a-c1127799c955/v2.0 login_hint:O.CiQ4YTVkZWYzMy0zMjJhLTQ0M2EtOGFjOC0zYWMxNGU3MDA0NDQSJDdhMDlhYWNlLTM2NDEtNDFiMC05OTNkLTM3MjkyMDEyMjhiMxoWam9zaHVhLmxlaG1hbm5AaGZ0bS5jaCCOAQ== name:Joshua Lehmann nbf:1.708770197e+09 nonce:c92d3774-3b7c-4f5d-92fa-fc2f539d682b oid:8eae5a8e-d0ae-4b89-8d4a-ada300d72719 preferred_username:joshua.lehmann@hftm.ch rh:0.ATsAxGZLdPctVkeQWsESd5nJVd_0rHL2eG9Ogcb1qh76jrw7ACA. roles:[Kursadministrator] sid:bcff908e-a50c-4c00-a873-3d6611c8bde8 sub:M8wesEuTiwsYzMlbCeS5rU5JjGvE2MZKCR9KzJ2EcbQ tenant_ctry:CH tenant_region_scope:EU tid:744b66c4-2df7-4756-905a-c1127799c955 uti:NfWu2nTF70CBKMrpu6sWAA ver:2.0 verified_primary_email:[joshua.lehmann@hftm.ch] verified_secondary_email:[joshua.lehmann@hftm.onmicrosoft.com] xms_tpl:de]
	 */

	// TODO: Finish this function
	return &models.User{}, nil
}
