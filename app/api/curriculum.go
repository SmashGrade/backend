package api

import (
	"github.com/SmashGrade/backend/app/config"
	"github.com/SmashGrade/backend/app/dao"
	"github.com/SmashGrade/backend/app/db"
	e "github.com/SmashGrade/backend/app/error"
	"github.com/SmashGrade/backend/app/repository"
	"github.com/SmashGrade/backend/app/requestmodels"
	"github.com/labstack/echo/v4"
)

// Handles the requests for the module resource
type CurriculumController struct {
	*BaseController
	Dao *dao.CurriculumDao
}

// Constructor for ModuleController
func NewCurriculumController(provider db.Provider) *CurriculumController {
	return &CurriculumController{
		BaseController: NewBaseController(provider),
		Dao: dao.NewCurriculumDao(
			repository.NewCurriculumRepository(provider),
			repository.NewFocusRepository(provider),
			repository.NewCurriculumtypeRepository(provider),
			repository.NewStateRepository(provider),
			repository.NewModuleRepository(provider),
		),
	}
}

// @Summary		Get all curriculums
// @Description	Get all curriculums
// @Tags			curriculums
// @Produce		json
// @Success		200	{array}		models.Curriculum
// @Failure		401	{object}	error.ApiError
// @Failure		403	{object}	error.ApiError
// @Failure		500	{object}	error.ApiError
// @Router			/curriculums [get]
// @Security		Bearer
func (c *CurriculumController) Curriculums(ctx echo.Context) error {
	// Check Role
	var allowedRoles []uint
	allowedRoles = append(allowedRoles, config.ROLE_COURSEADMIN)
	allowedRoles = append(allowedRoles, config.ROLE_FIELDMANAGER)
	allowedRoles = append(allowedRoles, config.ROLE_TEACHER)
	allowedRoles = append(allowedRoles, config.ROLE_STUDENT)
	roleErr := c.CheckUserRoles(allowedRoles, ctx)
	if roleErr != nil {
		return roleErr
	}

	res, err := c.Dao.GetAll()
	if err != nil {
		return err
	}
	return c.Yeet(ctx, res)
}

// @Summary		Get a specific curriculum
// @Description	Get a specific curriculum
// @Tags			curriculums
// @Param			id		path	uint		true	"Curriculum ID"
// @Param			date	path	time.Time	true	"Curriculum StartValidity"
// @Produce		json
// @Success		200	{object}	models.Curriculum
// @Failure		401	{object}	error.ApiError
// @Failure		403	{object}	error.ApiError
// @Failure		500	{object}	error.ApiError
// @Router			/curriculums/{id}/{date} [get]
// @Security		Bearer
func (c *CurriculumController) Curriculum(ctx echo.Context) error {
	// Check Role
	var allowedRoles []uint
	allowedRoles = append(allowedRoles, config.ROLE_COURSEADMIN)
	allowedRoles = append(allowedRoles, config.ROLE_FIELDMANAGER)
	allowedRoles = append(allowedRoles, config.ROLE_TEACHER)
	allowedRoles = append(allowedRoles, config.ROLE_STUDENT)
	roleErr := c.CheckUserRoles(allowedRoles, ctx)
	if roleErr != nil {
		return roleErr
	}

	// Read id parameter from request
	id, err := c.GetPathParamUint(ctx, "id")
	if err != nil {
		return e.NewDaoValidationError("id", "uint", c.GetPathParam(ctx, "id"))
	}

	// Read date paramater from request
	date, err := c.GetPathParamTime(ctx, "date")
	if err != nil {
		return e.NewDaoValidationError("date", "time.Time", c.GetPathParam(ctx, "date"))
	}

	// Get curriculum from dao
	res, daoErr := c.Dao.Get(id, date)
	if daoErr != nil {
		return err
	}
	return c.Yeet(ctx, res)
}

// @Summary		Create a curriculum
// @Description	Create a curriculum
// @Tags			curriculums
// @Produce		json
// @Accept			json
//
// @Param			request	body		requestmodels.RefCurriculum	true	"request body"
//
// @Success		200		{object}	models.Curriculum
// @Failure		401		{object}	error.ApiError
// @Failure		403		{object}	error.ApiError
// @Failure		500		{object}	error.ApiError
// @Router			/curriculums [post]
// @Security		Bearer
func (c *CurriculumController) Create(ctx echo.Context) error {
	// Check Role
	var allowedRoles []uint
	allowedRoles = append(allowedRoles, config.ROLE_COURSEADMIN)
	allowedRoles = append(allowedRoles, config.ROLE_FIELDMANAGER)
	allowedRoles = append(allowedRoles, config.ROLE_TEACHER)
	roleErr := c.CheckUserRoles(allowedRoles, ctx)
	if roleErr != nil {
		return roleErr
	}

	curriculum := new(requestmodels.RefCurriculum)
	// Read the request into curriculum
	if err := ctx.Bind(curriculum); err != nil {
		return e.ErrorInvalidRequest("curriculum")
	}

	// Create curriculum
	returnCurriculum, err := c.Dao.Create(*curriculum)
	if err != nil {
		return err
	}

	return c.Yeet(ctx, returnCurriculum)
}

// @Summary		Update a curriculum
// @Description	Update a curriculum
// @Tags			curriculums
// @Produce		json
// @Accept			json
//
// @Param			request	body		requestmodels.RefCurriculum	true	"request body"
// @Param			id		path		uint						true	"Curriculum ID"
// @Param			date	path		time.Time					true	"Curriculum StartValidity"
//
// @Success		200		{object}	models.Curriculum
// @Failure		401		{object}	error.ApiError
// @Failure		403		{object}	error.ApiError
// @Failure		500		{object}	error.ApiError
// @Router			/curriculums/{id}/{date} [put]
// @Security		Bearer
func (c *CurriculumController) Update(ctx echo.Context) error {
	// Check Role
	var allowedRoles []uint
	allowedRoles = append(allowedRoles, config.ROLE_COURSEADMIN)
	allowedRoles = append(allowedRoles, config.ROLE_FIELDMANAGER)
	allowedRoles = append(allowedRoles, config.ROLE_TEACHER)
	roleErr := c.CheckUserRoles(allowedRoles, ctx)
	if roleErr != nil {
		return roleErr
	}

	// Read id parameter from request
	id, err := c.GetPathParamUint(ctx, "id")
	if err != nil {
		return e.NewDaoValidationError("id", "uint", c.GetPathParam(ctx, "id"))
	}

	// Read date paramater from request
	date := c.GetPathParam(ctx, "date")

	curriculum := new(requestmodels.RefCurriculum)
	// Read the request into curriculum
	if err := ctx.Bind(curriculum); err != nil {
		return e.ErrorInvalidRequest("curriculum")
	}

	curriculum.ID = id
	curriculum.StartValidity = date

	// Create Curriculum
	daoErr := c.Dao.Update(*curriculum)
	if daoErr != nil {
		return daoErr
	}

	// return the result from the update
	return c.Yeet(ctx, curriculum)
}

// @Summary		Delete a curriculum
// @Description	Delete a curriculum
// @Tags			curriculums
// @Produce		json
// @Accept			json
//
// @Param			id		path	uint		true	"Curriculum ID"
// @Param			date	path	time.Time	true	"Curriculum StartValidity"
//
// @Success		200
// @Failure		401	{object}	error.ApiError
// @Failure		403	{object}	error.ApiError
// @Failure		500	{object}	error.ApiError
// @Router			/curriculums/{id}/{date} [delete]
// @Security		Bearer
func (c *CurriculumController) Delete(ctx echo.Context) error {
	// Check Role
	var allowedRoles []uint
	allowedRoles = append(allowedRoles, config.ROLE_COURSEADMIN)
	allowedRoles = append(allowedRoles, config.ROLE_FIELDMANAGER)
	roleErr := c.CheckUserRoles(allowedRoles, ctx)
	if roleErr != nil {
		return roleErr
	}

	// Read id parameter from request
	id, err := c.GetPathParamUint(ctx, "id")
	if err != nil {
		return e.NewDaoValidationError("id", "uint", c.GetPathParam(ctx, "id"))
	}

	// Read date paramater from request
	date, err := c.GetPathParamTime(ctx, "date")
	if err != nil {
		return e.NewDaoValidationError("date", "time.Time", c.GetPathParam(ctx, "date"))
	}

	// Delete Curriculum
	daoErr := c.Dao.Delete(id, date)
	if daoErr != nil {
		return daoErr
	}

	return c.Yeet(ctx, nil)
}

// register all output endpoints to router
func RegisterV1Curriculums(g *echo.Group, c *CurriculumController) {
	g.GET("/curriculums", c.Curriculums)
	g.GET("/curriculums/:id/:date", c.Curriculum)
	g.POST("/curriculums", c.Create)
	g.PUT("/curriculums/:id/:date", c.Update)
	g.DELETE("/curriculums/:id/:date", c.Delete)
}
