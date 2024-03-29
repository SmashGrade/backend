package api

import (
	"github.com/SmashGrade/backend/app/dao"
	"github.com/SmashGrade/backend/app/db"
	e "github.com/SmashGrade/backend/app/error"
	"github.com/SmashGrade/backend/app/repository"
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
		Dao:            dao.NewExamDao(repository.NewExamRepository(provider), repository.NewCourseRepository(provider)),
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
	// Check if the user has any role
	if authErr := c.CheckUserAnyRole(ctx); authErr != nil {
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

// register all output endpoints to router
func RegisterV1Exams(g *echo.Group, c *ExamController) {
	g.GET("/exams", c.Exams)
	g.GET("/exams/:id", c.Exam)
}
