package api

import (
	"github.com/SmashGrade/backend/app/dao"
	"github.com/SmashGrade/backend/app/db"
	e "github.com/SmashGrade/backend/app/error"
	"github.com/SmashGrade/backend/app/repository"
	"github.com/labstack/echo/v4"
)

// Handles the requests for the exam evaluation resource
type ExamevaluationController struct {
	*BaseController
	Dao *dao.ExamDao
}

// Constructor for ExamevaluationController
func NewExamevaluationController(provider db.Provider) *ExamevaluationController {
	return &ExamevaluationController{
		BaseController: NewBaseController(provider),
		Dao:            dao.NewExamDao(repository.NewExamRepository(provider), repository.NewCourseRepository(provider)),
	}
}

// @Summary		Get all exam evaluations
// @Description	Get all exam evaluations
// @Tags			evaluations
// @Produce		json
// @Success		200	{array}		models.ExamEvaluation
// @Failure		401	{object}	error.ApiError
// @Failure		403	{object}	error.ApiError
// @Failure		500	{object}	error.ApiError
// @Router			/evaluations [get]
// @Security		Bearer
func (c *ExamevaluationController) GetAllExamEvaluations(ctx echo.Context) error {
	return e.NewApiUnimplementedError()
}

// @Summary		Get all exam evaluations as student
// @Description	Get all exam evaluations as student uncategorized
// @Tags			evaluations
// @Produce		json
// @Success		200	{array}		models.ExamEvaluation
// @Failure		401	{object}	error.ApiError
// @Failure		403	{object}	error.ApiError
// @Failure		500	{object}	error.ApiError
// @Router			/evaluations/me [get]
// @Security		Bearer
func (c *ExamevaluationController) GetAllMyExamEvaluationsAsStudent(ctx echo.Context) error {
	return e.NewApiUnimplementedError()
}

// @Summary		Get all exam evaluations for a specific class as teacher
// @Description	Get all exam evaluations for a specific class as teacher
// @Tags			evaluations
// @Produce		json
//
// @Param			courseid		path	uint		true	"Course ID"
// @Param			courseversion		path	uint		true	"Course Version"
// @Param			year	path	time.Time	true	"SelectedCourse StartYear"
//
// @Success		200	{array}		models.ExamEvaluation
// @Failure		401	{object}	error.ApiError
// @Failure		403	{object}	error.ApiError
// @Failure		500	{object}	error.ApiError
// @Router			/evaluations/{courseid}/{courseversion}/{year}/teacher [get]
// @Security		Bearer
func (c *ExamevaluationController) GetAllExamEvaluationsForAClassAsTeacher(ctx echo.Context) error {
	return e.NewApiUnimplementedError()
}

// @Summary		Get all my exam evaluations for a specific class as student
// @Description	Get all my exam evaluations for a specific class as student
// @Tags			evaluations
// @Produce		json
//
// @Param			courseid		path	uint		true	"Course ID"
// @Param			courseversion		path	uint		true	"Course Version"
// @Param			year	path	time.Time	true	"SelectedCourse StartYear"
//
// @Success		200	{array}		models.ExamEvaluation
// @Failure		401	{object}	error.ApiError
// @Failure		403	{object}	error.ApiError
// @Failure		500	{object}	error.ApiError
// @Router			/evaluations/{courseid}/{courseversion}/{year}/me [get]
// @Security		Bearer
func (c *ExamevaluationController) GetAllExamEvaluationsForAClassAsStudent(ctx echo.Context) error {
	return e.NewApiUnimplementedError()
}

// @Summary		Get my exam evaluations for a specific exam as student
// @Description	Get my exam evaluations for a specific exam as student
// @Tags			evaluations
// @Produce		json
//
// @Param			courseid		path	uint		true	"Course ID"
// @Param			courseversion		path	uint		true	"Course Version"
// @Param			year	path	time.Time	true	"SelectedCourse StartYear"
// @Param			examid		path	uint		true	"Exam ID"
//
// @Success		200	{array}		models.ExamEvaluation
// @Failure		401	{object}	error.ApiError
// @Failure		403	{object}	error.ApiError
// @Failure		500	{object}	error.ApiError
// @Router			/evaluations/{courseid}/{courseversion}/{year}/{examid}/me [get]
// @Security		Bearer
func (c *ExamevaluationController) GetExamEvaluationsForATestAsStudent(ctx echo.Context) error {
	return e.NewApiUnimplementedError()
}

// @Summary		Get exam evaluations for a specific exam
// @Description	Get exam evaluations for a specific exam
// @Tags			evaluations
// @Produce		json
//
// @Param			courseid		path	uint		true	"Course ID"
// @Param			courseversion		path	uint		true	"Course Version"
// @Param			year	path	time.Time	true	"SelectedCourse StartYear"
// @Param			examid		path	uint		true	"Exam ID"
// @Param			userid		path	uint		true	"User ID"
//
// @Success		200	{array}		models.ExamEvaluation
// @Failure		401	{object}	error.ApiError
// @Failure		403	{object}	error.ApiError
// @Failure		500	{object}	error.ApiError
// @Router			/evaluations/{courseid}/{courseversion}/{year}/{examid}/{userid} [get]
// @Security		Bearer
func (c *ExamevaluationController) GetExamEvaluationsForATest(ctx echo.Context) error {
	return e.NewApiUnimplementedError()
}

// @Summary		Get all exam evaluations for a specific exam as teacher
// @Description	Get all exam evaluations for a specific exam as teacher
// @Tags			evaluations
// @Produce		json
//
// @Param			courseid		path	uint		true	"Course ID"
// @Param			courseversion		path	uint		true	"Course Version"
// @Param			year	path	time.Time	true	"SelectedCourse StartYear"
// @Param			examid		path	uint		true	"Exam ID"
//
// @Success		200	{array}		models.ExamEvaluation
// @Failure		401	{object}	error.ApiError
// @Failure		403	{object}	error.ApiError
// @Failure		500	{object}	error.ApiError
// @Router			/evaluations/{courseid}/{courseversion}/{year}/{examid}/teacher [get]
// @Security		Bearer
func (c *ExamevaluationController) GetAllExamEvaluationsForATestAsTeacher(ctx echo.Context) error {
	return e.NewApiUnimplementedError()
}

// @Summary		Create new exam evaluation
// @Description	Create new exam evaluation
// @Tags			evaluations
// @Produce		json
// @Accept			json
//
// @Param			request	body		requestmodels.RefExamEvaluation	true	"request body"
// @Success		200	{object}		models.ExamEvaluation
// @Failure		401	{object}	error.ApiError
// @Failure		403	{object}	error.ApiError
// @Failure		500	{object}	error.ApiError
// @Router			/evaluations [post]
// @Security		Bearer
func (c *ExamevaluationController) CreateNewExamEvaluation(ctx echo.Context) error {
	return e.NewApiUnimplementedError()
}

// @Summary		Update a exam evaluation
// @Description	Update a exam evaluation
// @Tags			evaluations
// @Produce		json
// @Accept			json
//
// @Param			request	body		requestmodels.RefExamEvaluation	true	"request body"
// @Param			courseid		path	uint		true	"Course ID"
// @Param			courseversion		path	uint		true	"Course Version"
// @Param			year	path	time.Time	true	"SelectedCourse StartYear"
// @Param			examid		path	uint		true	"Exam ID"
// @Param			userid		path	uint		true	"User ID"
// @Param			evaluationid		path	uint		true	"Evaluation ID"
//
// @Success		200		{object}	models.Curriculum
// @Failure		401		{object}	error.ApiError
// @Failure		403		{object}	error.ApiError
// @Failure		500		{object}	error.ApiError
// @Router			/evaluations/{courseid}/{courseversion}/{year}/{examid}/{userid}/{evaluationid} [put]
// @Security		Bearer
func (c *ExamevaluationController) UpdateExamEvaluation(ctx echo.Context) error {
	return e.NewApiUnimplementedError()
}

// @Summary		Delete a curriculum
// @Description	Delete a curriculum
// @Tags			curriculums
// @Produce		json
// @Accept			json
//
// @Param			courseid		path	uint		true	"Course ID"
// @Param			courseversion		path	uint		true	"Course Version"
// @Param			year	path	time.Time	true	"SelectedCourse StartYear"
// @Param			examid		path	uint		true	"Exam ID"
// @Param			userid		path	uint		true	"User ID"
// @Param			evaluationid		path	uint		true	"Evaluation ID"
//
// @Success		200
// @Failure		401	{object}	error.ApiError
// @Failure		403	{object}	error.ApiError
// @Failure		500	{object}	error.ApiError
// @Router			/evaluations/{courseid}/{courseversion}/{year}/{examid}/{userid}/{evaluationid} [delete]
// @Security		Bearer
func (c *ExamevaluationController) DeleteExamEvaluation(ctx echo.Context) error {
	return e.NewApiUnimplementedError()
}

// register all output endpoints to router
func RegisterV1Examevaluations(g *echo.Group, c *ExamevaluationController) {
	g.GET("/evaluations", c.GetAllExamEvaluations)
	g.GET("/evaluations/me", c.GetAllMyExamEvaluationsAsStudent)
	g.GET("/evaluations/:courseid/:courseversion/:year/teacher", c.GetAllExamEvaluationsForAClassAsTeacher)
	g.GET("/evaluations/:courseid/:courseversion/:year/me", c.GetAllExamEvaluationsForAClassAsStudent)
	g.GET("/evaluations/:courseid/:courseversion/:year/:examid/me", c.GetExamEvaluationsForATestAsStudent)
	g.GET("/evaluations/:courseid/:courseversion/:year/:examid/:userid", c.GetExamEvaluationsForATest)
	g.GET("/evaluations/:courseid/:courseversion/:year/:examid/teacher", c.GetAllExamEvaluationsForATestAsTeacher)
	g.POST("/evaluations", c.CreateNewExamEvaluation)
	g.PUT("/evaluations/:courseid/:courseversion/:year/:examid/:userid/:evaluationid", c.UpdateExamEvaluation)
	g.DELETE("/evaluations/:courseid/:courseversion/:year/:examid/:userid/:evaluationid", c.DeleteExamEvaluation)
}
