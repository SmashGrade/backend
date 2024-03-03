package api

import (
	"github.com/SmashGrade/backend/app/dao"
	"github.com/SmashGrade/backend/app/db"
	e "github.com/SmashGrade/backend/app/error"
	"github.com/SmashGrade/backend/app/repository"
	"github.com/labstack/echo/v4"
)

// Handles the requests for the module resource
type ModuleController struct {
	*BaseController
	Dao *dao.ModuleDao
}

// Constructor for ModuleController
func NewModuleController(provider db.Provider) *ModuleController {
	return &ModuleController{
		BaseController: NewBaseController(provider),
		Dao:            dao.NewModuleDao(repository.NewModuleRepository(provider)),
	}
}

// @Summary		Get all modules
// @Description	Get all modules
// @Tags			modules
// @Produce		json
// @Success		200	{array}		models.Module
// @Failure		401	{object}	error.ApiError
// @Failure		403	{object}	error.ApiError
// @Failure		500	{object}	error.ApiError
// @Router			/modules [get]
// @Security		Bearer
func (c *ModuleController) Modules(ctx echo.Context) error {
	res, err := c.Dao.GetAll()
	if err != nil {
		return err
	}
	return c.Yeet(ctx, res)
}

// @Summary		Get a specific module
// @Description	Get a specific module
// @Tags			modules
// @Param			id		path	uint	true	"Module ID"
// @Param			version	path	uint	true	"Module Version"
// @Produce		json
// @Success		200	{object}	models.Module
// @Failure		401	{object}	error.ApiError
// @Failure		403	{object}	error.ApiError
// @Failure		500	{object}	error.ApiError
// @Router			/modules/{id}/{version} [get]
// @Security		Bearer
func (c *ModuleController) Module(ctx echo.Context) error {
	// Read id parameter from request
	id, err := c.GetPathParamUint(ctx, "id")
	if err != nil {
		return e.NewDaoValidationError("id", "uint", c.GetPathParam(ctx, "id"))
	}
	// Read version parameter from request
	version, err := c.GetPathParamUint(ctx, "version")
	if err != nil {
		return e.NewDaoValidationError("version", "uint", c.GetPathParam(ctx, "version"))
	}
	// Ask the DAO for the module
	res, daoErr := c.Dao.Get(id, version)
	if daoErr != nil {
		return daoErr
	}
	// Return the result to the client
	return c.Yeet(ctx, res)
}

// register all output endpoints to router
func RegisterV1Modules(g *echo.Group, c *ModuleController) {
	g.GET("/modules", c.Modules)
	g.GET("/modules/:id/:version", c.Module)
}
