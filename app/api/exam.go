package api

import (
	"github.com/SmashGrade/backend/app/dao"
	"github.com/SmashGrade/backend/app/db"
	e "github.com/SmashGrade/backend/app/error"
	"github.com/SmashGrade/backend/app/repository"
	"github.com/SmashGrade/backend/app/requestmodels"
	"github.com/labstack/echo/v4"
)

// Handles the requests for the exam resource
type ExamController struct {
	*BaseController
	Dao *dao.ExamDao
}

// Constructor for ExamController
func NewExamController(provider db.Provider) *ExamController {
	return &ExamController{
		BaseController: NewBaseController(provider),
		Dao:            dao.NewExamDao(repository.NewExamRepository(provider)),
	}
}

// @Summary		Get all exams
// @Description	Get all exams
// @Tags			exams
// @Produce		json
// @Success		200	{array}		models.Exam
// @Failure		401	{object}	error.ApiError
// @Failure		403	{object}	error.ApiError
// @Failure		500	{object}	error.ApiError
// @Router			/exams [get]
// @Security		Bearer
func (c *ExamController) Exams(ctx echo.Context) error {
	// Check if the user has required roles
	authErr := c.CheckUserRoles(ROLEGROUP_STUDENT, ctx)
	if authErr != nil {
		return authErr
	}
	res, err := c.Dao.GetAll()
	if err != nil {
		return err
	}
	return c.Yeet(ctx, res)
}

// @Summary		Get a specific exam
// @Description	Get a specific exam
// @Tags			exams
// @Param			id	path	uint	true	"Exam ID"
// @Produce		json
// @Success		200	{object}	models.Exam
// @Failure		401	{object}	error.ApiError
// @Failure		403	{object}	error.ApiError
// @Failure		500	{object}	error.ApiError
// @Router			/exams/{id} [get]
// @Security		Bearer
func (c *ExamController) Exam(ctx echo.Context) error {
	// Check if the user has required roles
	authErr := c.CheckUserRoles(ROLEGROUP_STUDENT, ctx)
	if authErr != nil {
		return authErr
	}

	// Read id parameter from request
	id, err := c.GetPathParamUint(ctx, "id")
	if err != nil {
		return e.NewDaoValidationError("id", "uint", c.GetPathParam(ctx, "id"))
	}
	// Ask the DAO for the exam
	res, daoErr := c.Dao.Get(id)
	if daoErr != nil {
		return daoErr
	}
	// Return the result to the client
	return c.Yeet(ctx, res)
}

// @Summary		Create a exam
// @Description	Create a exam
// @Tags			exams
// @Produce		json
// @Accept			json
//
// @Param			request	body		requestmodels.RefExam	true	"request body"
//
// @Success		200		{object}	models.Exam
// @Failure		401		{object}	error.ApiError
// @Failure		403		{object}	error.ApiError
// @Failure		500		{object}	error.ApiError
// @Router			/exams [post]
// @Security		Bearer
func (c *ExamController) Create(ctx echo.Context) error {
	// Check if the user has required roles
	authErr := c.CheckUserRoles(ROLEGROUP_TEACHER, ctx)
	if authErr != nil {
		return authErr
	}

	exam := new(requestmodels.RefExam)
	// Read the request into the exam
	if err := ctx.Bind(exam); err != nil {
		return e.ErrorInvalidRequest("exam")
	}
	// Let dao create the exam
	returnExam, err := c.Dao.Create(*exam)
	if err != nil {
		return err
	}
	// return the result from the post
	return c.Yeet(ctx, returnExam)
}

// @Summary		Update a exam
// @Description	Update a exam
// @Tags			exams
// @Produce		json
// @Accept			json
//
// @Param			request	body		requestmodels.RefExam	true	"request body"
// @Param			id		path		uint					true	"Exam ID"
//
// @Success		200		{object}	models.Exam
// @Failure		401		{object}	error.ApiError
// @Failure		403		{object}	error.ApiError
// @Failure		500		{object}	error.ApiError
// @Router			/exams/{id}/{version} [put]
// @Security		Bearer
func (c *ExamController) Update(ctx echo.Context) error {
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

	exam := new(requestmodels.RefExam)
	// Read the request into the exam
	if err := ctx.Bind(exam); err != nil {
		return e.ErrorInvalidRequest("exam")
	}

	exam.ID = id

	// Let dao update the Exam
	daoErr := c.Dao.Update(*exam)
	if daoErr != nil {
		return daoErr
	}

	return c.Yeet(ctx, exam)
}

// @Summary		Delete a exam
// @Description	Delete a exam
// @Tags			exams
// @Produce		json
// @Accept			json
//
// @Param			id	path	uint	true	"Exam ID"
//
// @Success		200
// @Failure		401	{object}	error.ApiError
// @Failure		403	{object}	error.ApiError
// @Failure		500	{object}	error.ApiError
// @Router			/exams/{id}/{version} [delete]
// @Security		Bearer
func (c *ExamController) Delete(ctx echo.Context) error {
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

	daoErr := c.Dao.Delete(id)
	if daoErr != nil {
		return daoErr
	}

	return c.Yeet(ctx, nil)
}

// @Summary		Get all exams from a course
// @Description	Get all exams from a course
// @Tags			exams
// @Param			id		path	uint	true	"Course ID"
// @Param			version	path	uint	true	"Course Version"
// @Produce		json
// @Success		200	{array}		models.Exam
// @Failure		401	{object}	error.ApiError
// @Failure		403	{object}	error.ApiError
// @Failure		500	{object}	error.ApiError
// @Router			/exams/courses/{id}/{version} [get]
// @Security		Bearer
func (c *ExamController) GetFromCourse(ctx echo.Context) error {
	// Check if the user has required roles
	authErr := c.CheckUserRoles(ROLEGROUP_STUDENT, ctx)
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

	res, daoErr := c.Dao.GetForCourse(id, version)
	if daoErr != nil {
		return daoErr
	}
	return c.Yeet(ctx, res)
}

// register all output endpoints to router
func RegisterV1Exams(g *echo.Group, c *ExamController) {
	g.GET("/exams", c.Exams)
	g.GET("/exams/:id", c.Exam)
	g.POST("/exams", c.Create)
	g.PUT("/exams/:id", c.Update)
	g.DELETE("/exams/:id", c.Delete)
	g.GET("/exams/courses/:id/:version", c.GetFromCourse)
}
