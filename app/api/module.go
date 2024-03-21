package api

import (
	"github.com/SmashGrade/backend/app/dao"
	"github.com/SmashGrade/backend/app/db"
	e "github.com/SmashGrade/backend/app/error"
	"github.com/SmashGrade/backend/app/repository"
	"github.com/SmashGrade/backend/app/requestmodels"
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

// @Summary		Create a module
// @Description	Create a module
// @Tags			modules
// @Produce		json
// @Accept			json
//
// @Param			request	body		requestmodels.RefModule	true	"request body"
//
// @Success		200		{object}	models.Module
// @Failure		401		{object}	error.ApiError
// @Failure		403		{object}	error.ApiError
// @Failure		500		{object}	error.ApiError
// @Router			/modules [post]
// @Security		Bearer
func (c *ModuleController) Create(ctx echo.Context) error {
	// Check if the user has required roles
	authErr := c.CheckUserRoles(ROLEGROUP_TEACHER, ctx)
	if authErr != nil {
		return authErr
	}

	module := new(requestmodels.RefModule)
	// Read the request into Reference
	if err := ctx.Bind(module); err != nil {
		return e.ErrorInvalidRequest("module")
	}
	// Let dao create the module
	returnModule, err := c.Dao.Create(*module)
	if err != nil {
		return err
	}
	// return the result from the Post
	return c.Yeet(ctx, returnModule)
}

// @Summary		Create a new version of a module
// @Description	Create a new version of a module
// @Tags			modules
// @Produce		json
// @Accept			json
//
// @Param			request	body		requestmodels.RefModule	true	"request body"
// @Param			id		path		uint					true	"Module ID"
//
// @Success		200		{object}	models.Module
// @Failure		401		{object}	error.ApiError
// @Failure		403		{object}	error.ApiError
// @Failure		500		{object}	error.ApiError
// @Router			/modules/{id} [post]
// @Security		Bearer
func (c *ModuleController) CreateVersion(ctx echo.Context) error {
	// Check if the user has required roles
	authErr := c.CheckUserRoles(ROLEGROUP_TEACHER, ctx)
	if authErr != nil {
		return authErr
	}

	id, err := c.GetPathParamUint(ctx, "id")
	if err != nil {
		return e.NewDaoValidationError("id", "uint", c.GetPathParam(ctx, "id"))
	}

	module := new(requestmodels.RefModule)
	// Read the request into module
	if err := ctx.Bind(module); err != nil {
		return e.ErrorInvalidRequest("module")
	}

	module.ID = id
	module.Version = 0

	// Let dao create the module
	returnModule, daoErr := c.Dao.Create(*module)
	if daoErr != nil {
		return daoErr
	}
	// return the result from the Post
	return c.Yeet(ctx, returnModule)
}

// @Summary		Update a module
// @Description	Update a module
// @Tags			modules
// @Produce		json
// @Accept			json
//
// @Param			request	body		requestmodels.RefModule	true	"request body"
// @Param			id		path		uint					true	"Module ID"
// @Param			version	path		uint					true	"Module Version"
//
// @Success		200		{object}	models.Module
// @Failure		401		{object}	error.ApiError
// @Failure		403		{object}	error.ApiError
// @Failure		500		{object}	error.ApiError
// @Router			/modules/{id}/{version} [put]
// @Security		Bearer
func (c *ModuleController) Update(ctx echo.Context) error {
	// Check if the user has required roles
	authErr := c.CheckUserRoles(ROLEGROUP_TEACHER, ctx)
	if authErr != nil {
		return authErr
	}

	// Read id parameter from request
	id, err := c.GetPathParamUint(ctx, "id")
	if err != nil {
		return e.NewDaoValidationError("id", "uint", c.GetPathParam(ctx, "id"))
	}

	version, err := c.GetPathParamUint(ctx, "version")
	if err != nil {
		return e.NewDaoValidationError("version", "uint", c.GetPathParam(ctx, "version"))
	}

	module := new(requestmodels.RefModule)
	// Read the request into Module
	if err := ctx.Bind(module); err != nil {
		return e.ErrorInvalidRequest("module")
	}

	module.ID = id
	module.Version = version

	// Let dao create the module
	daoErr := c.Dao.Update(*module)
	if daoErr != nil {
		return daoErr
	}

	return c.Yeet(ctx, module)
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
	// Check if the user has any role
	if authErr := c.CheckUserAnyRole(ctx); authErr != nil {
		return authErr
	}

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
	// Check if the user has any role
	if authErr := c.CheckUserAnyRole(ctx); authErr != nil {
		return authErr
	}

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

func (c *ModuleController) Delete(ctx echo.Context) error {
	// Check if the user has required roles
	authErr := c.CheckUserRoles(ROLEGROUP_ADMIN, ctx)
	if authErr != nil {
		return authErr
	}

	id, err := c.GetPathParamUint(ctx, "id")
	if err != nil {
		return e.NewDaoValidationError("id", "uint", c.GetPathParam(ctx, "version"))
	}

	version, err := c.GetPathParamUint(ctx, "version")
	if err != nil {
		return e.NewDaoValidationError("version", "uint", c.GetPathParam(ctx, "version"))
	}

	// Let dao create the module
	daoErr := c.Dao.Delete(id, version)
	if daoErr != nil {
		return daoErr
	}
	// return the result from the Post
	return c.Yeet(ctx, nil)
}

// register all output endpoints to router
func RegisterV1Modules(g *echo.Group, c *ModuleController) {
	g.GET("/modules", c.Modules)
	g.GET("/modules/:id/:version", c.Module)
	g.POST("/modules", c.Create)
	g.POST("/modules/:id", c.CreateVersion)
	g.PUT("/modules/:id/:version", c.Update)
	g.DELETE("/modules/:id/:version", c.Delete)
}
