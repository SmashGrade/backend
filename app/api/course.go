package api

import (
	"github.com/SmashGrade/backend/app/dao"
	"github.com/SmashGrade/backend/app/db"
	"github.com/SmashGrade/backend/app/repository"
	"github.com/labstack/echo/v4"
)

// Handles the requests for the course resource
type CourseController struct {
	*BaseController
	Dao *dao.CourseDao
}

// Constructor for CourseController
func NewCourseController(provider db.Provider) *CourseController {
	return &CourseController{
		BaseController: NewBaseController(provider),
		Dao:            dao.NewCourseDao(repository.NewCourseRepository(provider)),
	}
}

func (c *CourseController) Courses(ctx echo.Context) error {
	return nil
}

// Handles the GET /course/(id) request
func (c *CourseController) Course(ctx echo.Context) error {
	return nil
}
