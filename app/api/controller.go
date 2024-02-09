package api

import (
	"strconv"

	"github.com/SmashGrade/backend/app/db"
	"github.com/labstack/echo/v4"
)

// Controller is the base controller for all API controllers
type BaseController struct {
	Provider db.Provider
}

// Constructor for BaseController
func NewBaseController(provider db.Provider) *BaseController {
	return &BaseController{
		Provider: provider,
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
