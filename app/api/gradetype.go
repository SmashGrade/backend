package api

import (
	"github.com/SmashGrade/backend/app/dao"
	"github.com/SmashGrade/backend/app/db"
	e "github.com/SmashGrade/backend/app/error"
	"github.com/SmashGrade/backend/app/models"
	"github.com/SmashGrade/backend/app/repository"
	"github.com/labstack/echo/v4"
)

// Handles the requests for the gradetype resource
type GradeTypeController struct {
	*BaseController
	Dao *dao.GradeTypeDao
}

// Constructor for GradeTypeController
func NewGradeTypeController(provider db.Provider) *GradeTypeController {
	return &GradeTypeController{
		BaseController: NewBaseController(provider),
		Dao:            dao.NewGradeTypeDao(repository.NewGradetypeRepository(provider)),
	}
}

// @Summary		Get all gradetypes
// @Description	Get all gradetypes
// @Tags			gradetype
// @Produce		json
// @Success		200	{array}		models.Gradetype
// @Failure		401	{object}	error.ApiError
// @Failure		403	{object}	error.ApiError
// @Failure		500	{object}	error.ApiError
// @Router			/gradetypes [get]
// @Security		Bearer
func (c *GradeTypeController) GradeTypes(ctx echo.Context) error {
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

// @Summary		Get a specific gradetype
// @Description	Get a specific gradetype
// @Tags			gradetype
// @Param			id	path	uint	true	"GradeType ID"
// @Produce		json
// @Success		200	{object}	models.Gradetype
// @Failure		401	{object}	error.ApiError
// @Failure		403	{object}	error.ApiError
// @Failure		500	{object}	error.ApiError
// @Router			/gradetypes/{id}/ [get]
// @Security		Bearer
func (c *GradeTypeController) GradeType(ctx echo.Context) error {
	// Check if the user has any role
	if authErr := c.CheckUserAnyRole(ctx); authErr != nil {
		return authErr
	}

	// Read id parameter from request
	id, err := c.GetPathParamUint(ctx, "id")
	if err != nil {
		return e.NewDaoValidationError("id", "uint", c.GetPathParam(ctx, "id"))
	}

	// Ask the DAO for the gradetype
	res, daoErr := c.Dao.Get(id)
	if daoErr != nil {
		return daoErr
	}
	// Return the result to the client
	return c.Yeet(ctx, res)
}

// @Summary		Create a gradetype
// @Description	Create a gradetype
// @Tags			gradetype
// @Produce		json
// @Accept			json
//
// @Param			request	body		models.Gradetype	true	"request body"
//
// @Success		200		{object}	models.Gradetype
// @Failure		401		{object}	error.ApiError
// @Failure		403		{object}	error.ApiError
// @Failure		500		{object}	error.ApiError
// @Router			/gradetypes [post]
// @Security		Bearer
func (c *GradeTypeController) Create(ctx echo.Context) error {
	// Check if the user has required roles
	authErr := c.CheckUserRoles(ROLEGROUP_ADMIN, ctx)
	if authErr != nil {
		return authErr
	}

	gradeType := new(models.Gradetype)
	// Read the request into the gradetype
	if err := ctx.Bind(gradeType); err != nil {
		return e.ErrorInvalidRequest("gradetype")
	}
	// Let dao create the gradeType
	returnGradeType, err := c.Dao.Create(*gradeType)
	if err != nil {
		return err
	}
	// return the result from the Post
	return c.Yeet(ctx, returnGradeType)
}

// @Summary		Update a gradetype
// @Description	Update a gradetype
// @Tags			grapetype
// @Produce		json
// @Accept			json
//
// @Param			request	body		models.Gradetype	true	"request body"
// @Param			id		path		uint				true	"GradeType ID"
//
// @Success		200		{object}	models.Gradetype
// @Failure		401		{object}	error.ApiError
// @Failure		403		{object}	error.ApiError
// @Failure		500		{object}	error.ApiError
// @Router			/gradetypes/{id} [put]
// @Security		Bearer
func (c *GradeTypeController) Update(ctx echo.Context) error {
	// Check if the user has required roles
	authErr := c.CheckUserRoles(ROLEGROUP_ADMIN, ctx)
	if authErr != nil {
		return authErr
	}

	// Read id parameter from request
	id, err := c.GetPathParamUint(ctx, "id")
	if err != nil {
		return e.NewDaoValidationError("id", "uint", c.GetPathParam(ctx, "id"))
	}

	gradeType := new(models.Gradetype)
	// Read the request into gradeType
	if err := ctx.Bind(gradeType); err != nil {
		return e.ErrorInvalidRequest("gradetype")
	}

	gradeType.ID = id

	// Let dao create the gradeType
	daoErr := c.Dao.Update(*gradeType)
	if daoErr != nil {
		return daoErr
	}
	// return the result from the gradeType
	return c.Yeet(ctx, gradeType)
}

// @Summary		Delete a gradetype
// @Description	Delete a gradetype
// @Tags			gradetypes
// @Produce		json
// @Accept			json
//
// @Param			id	path	uint	true	"GradeType ID"
//
// @Success		200
// @Failure		401	{object}	error.ApiError
// @Failure		403	{object}	error.ApiError
// @Failure		500	{object}	error.ApiError
// @Router			/gradetypes/{id} [delete]
// @Security		Bearer
func (c *GradeTypeController) Delete(ctx echo.Context) error {
	// Check if the user has required roles
	authErr := c.CheckUserRoles(ROLEGROUP_ADMIN, ctx)
	if authErr != nil {
		return authErr
	}

	// Read id parameter from request
	id, err := c.GetPathParamUint(ctx, "id")
	if err != nil {
		return e.NewDaoValidationError("id", "uint", c.GetPathParam(ctx, "id"))
	}

	// Let dao delete the gradetype
	daoErr := c.Dao.Delete(id)
	if daoErr != nil {
		return daoErr
	}
	// return the result from the Post
	return c.Yeet(ctx, nil)
}

// register all output endpoints to router
func RegisterV1GradeTypes(g *echo.Group, c *GradeTypeController) {
	g.GET("/gradetypes", c.GradeType)
	g.GET("/gradetypes/:id", c.GradeTypes)
	g.POST("/gradetypes", c.Create)
	g.PUT("/gradetypes/:id", c.Update)
	g.DELETE("/gradetypes/:id", c.Delete)
}
