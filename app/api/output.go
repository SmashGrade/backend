package api

import (
	"github.com/SmashGrade/backend/app/dao"
	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/repository"
	"github.com/labstack/echo/v4"
)

// contains all output list only endpoints to fill frontend drop downs and co

// All list outputs under one big object
type OutputController struct {
	*BaseController
	//fieldDao dao.fieldDao
	stateDao *dao.StateDao
}

// Constructor for output controller
func NewOutputController(provider db.Provider) *OutputController {
	return &OutputController{
		BaseController: NewBaseController(provider),
		stateDao:       dao.NewStateDao(repository.NewStateRepository(provider)),
	}
}

// @Summary Get all states
// @Description Get all states
// @Tags states
// @Produce json
// @Success 200 {array} models.State
// @Failure 401 {object} error.ApiError
// @Failure 403 {object} error.ApiError
// @Failure 500 {object} error.ApiError
// @Router /states [get]
func (c *OutputController) States(ctx echo.Context) error {
	res, err := c.stateDao.GetAll()
	if err != nil {
		return err
	}
	return c.Yeet(ctx, res)
}

// register all output endpoints to router
func RegisterV1Output(g *echo.Group, o *OutputController) {
	g.GET("/states", o.States)
}
