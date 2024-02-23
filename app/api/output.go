package api

import (
	"log"

	"github.com/SmashGrade/backend/app/dao"
	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/repository"
	"github.com/labstack/echo/v4"
)

// contains all output list only endpoints to fill frontend drop downs and co

// All list outputs under one big object
type OutputController struct {
	*BaseController
	stateDao           *dao.StateDao
	curriculumytypeDao *dao.CurriculumTypeDao
	gradetypeDao       *dao.GradeTypeDao
	evaluationtypeDao  *dao.EvaluationTypeDao
	fieldDao           *dao.FieldDao
	focusDao           *dao.FocusDao
	examtypeDao        *dao.ExamtypeDao
}

// Constructor for output controller
func NewOutputController(provider db.Provider) *OutputController {
	ctrl := &OutputController{
		BaseController:     NewBaseController(provider),
		stateDao:           dao.NewStateDao(repository.NewStateRepository(provider)),
		curriculumytypeDao: dao.NewCurriculumTypeDao(repository.NewCurriculumtypeRepository(provider)),
		gradetypeDao:       dao.NewGradeTypeDao(repository.NewGradetypeRepository(provider)),
		evaluationtypeDao:  dao.NewEvaluationTypeDao(repository.NewEvaluationtypeRepository(provider)),
		fieldDao:           dao.NewFieldDao(repository.NewFieldRepository(provider)),
		focusDao:           dao.NewFocusDao(repository.NewFocusRepository(provider)),
		examtypeDao:        dao.NewExamtypeDao(repository.NewExamtypeRepository(provider)),
	}

	if provider.Config().AutoMigrate {
		log.Println("Creating defaults from config...")
		ctrl.stateDao.CreateDefaults()
		ctrl.curriculumytypeDao.CreateDefaults()
		ctrl.gradetypeDao.CreateDefaults()
		ctrl.evaluationtypeDao.CreateDefaults()
		ctrl.examtypeDao.CreateDefaults()
		log.Println("Creating defaults from config completed successfully")
	}

	return ctrl
}

//	@Summary		Get all states
//	@Description	Get all states
//	@Tags			states
//	@Produce		json
//	@Success		200	{array}		models.State
//	@Failure		401	{object}	error.ApiError
//	@Failure		403	{object}	error.ApiError
//	@Failure		500	{object}	error.ApiError
//	@Router			/states [get]
//	@Security		Bearer
func (c *OutputController) States(ctx echo.Context) error {
	res, err := c.stateDao.GetAll()
	if err != nil {
		return err
	}
	return c.Yeet(ctx, res)
}

//	@Summary		Get all curriculum types
//	@Description	Get all curriculum types
//	@Tags			curriculumtypes
//	@Produce		json
//	@Success		200	{array}		models.Curriculumtype
//	@Failure		401	{object}	error.ApiError
//	@Failure		403	{object}	error.ApiError
//	@Failure		500	{object}	error.ApiError
//	@Router			/curriculumtypes [get]
//	@Security		Bearer
func (c *OutputController) Curriculumtypes(ctx echo.Context) error {
	res, err := c.curriculumytypeDao.GetAll()
	if err != nil {
		return err
	}
	return c.Yeet(ctx, res)
}

//	@Summary		Get all grade types
//	@Description	Get all grade types
//	@Tags			gradetypes
//	@Produce		json
//	@Success		200	{array}		models.Gradetype
//	@Failure		401	{object}	error.ApiError
//	@Failure		403	{object}	error.ApiError
//	@Failure		500	{object}	error.ApiError
//	@Router			/gradetypes [get]
//	@Security		Bearer
func (c *OutputController) Gradetypes(ctx echo.Context) error {
	res, err := c.gradetypeDao.GetAll()
	if err != nil {
		return err
	}
	return c.Yeet(ctx, res)
}

//	@Summary		Get all evaluation types
//	@Description	Get all evaluation types
//	@Tags			evaluationtypes
//	@Produce		json
//	@Success		200	{array}		models.Evaluationtype
//	@Failure		401	{object}	error.ApiError
//	@Failure		403	{object}	error.ApiError
//	@Failure		500	{object}	error.ApiError
//	@Router			/evaluationtypes [get]
//	@Security		Bearer
func (c *OutputController) Evaluationtypes(ctx echo.Context) error {
	res, err := c.evaluationtypeDao.GetAll()
	if err != nil {
		return err
	}
	return c.Yeet(ctx, res)
}

//	@Summary		Get all fields
//	@Description	Get all fields
//	@Tags			fields
//	@Produce		json
//	@Success		200	{array}		models.Field
//	@Failure		401	{object}	error.ApiError
//	@Failure		403	{object}	error.ApiError
//	@Failure		500	{object}	error.ApiError
//	@Router			/fields [get]
//	@Security		Bearer
func (c *OutputController) Fields(ctx echo.Context) error {
	res, err := c.fieldDao.GetAll()
	if err != nil {
		return err
	}
	return c.Yeet(ctx, res)
}

//	@Summary		Get all focuses
//	@Description	Get all focuses
//	@Tags			focuses
//	@Produce		json
//	@Success		200	{array}		models.Focus
//	@Failure		401	{object}	error.ApiError
//	@Failure		403	{object}	error.ApiError
//	@Failure		500	{object}	error.ApiError
//	@Router			/focuses [get]
//	@Security		Bearer
func (c *OutputController) Focuses(ctx echo.Context) error {
	res, err := c.focusDao.GetAll()
	if err != nil {
		return err
	}
	return c.Yeet(ctx, res)
}

//	@Summary		Get all exam types
//	@Description	Get all exam types
//	@Tags			examtypes
//	@Produce		json
//	@Success		200	{array}		models.Examtype
//	@Failure		401	{object}	error.ApiError
//	@Failure		403	{object}	error.ApiError
//	@Failure		500	{object}	error.ApiError
//	@Router			/examtypes [get]
//	@Security		Bearer
func (c *OutputController) Examtypes(ctx echo.Context) error {
	res, err := c.examtypeDao.GetAll()
	if err != nil {
		return err
	}
	return c.Yeet(ctx, res)
}

// register all output endpoints to router
func RegisterV1Output(g *echo.Group, o *OutputController) {
	g.GET("/states", o.States)
	g.GET("/curriculumtypes", o.Curriculumtypes)
	g.GET("/gradetypes", o.Gradetypes)
	g.GET("/evaluationtypes", o.Evaluationtypes)
	g.GET("/fields", o.Fields)
	g.GET("/focuses", o.Focuses)
	g.GET("/examtypes", o.Examtypes)
}
