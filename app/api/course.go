package api

import (
	"github.com/SmashGrade/backend/app/dao"
	"github.com/SmashGrade/backend/app/db"
	e "github.com/SmashGrade/backend/app/error"
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

// @Summary Get all courses
// @Description Get all courses
// @Tags courses
// @Produce json
// @Success 200 {array} models.Course
// @Failure 401 {object} error.ApiError
// @Failure 403 {object} error.ApiError
// @Failure 500 {object} error.ApiError
// @Router /courses [get]
func (c *CourseController) Courses(ctx echo.Context) error {
	res, err := c.Dao.GetAll()
	if err != nil {
		return err
	}
	return c.Yeet(ctx, res)
}

// @Summary Get a specific course
// @Description Get a specific course
// @Tags courses
// @Param id path uint true "Course ID"
// @Param version path uint true "Course Version"
// @Produce json
// @Success 200 {object} models.Course
// @Failure 401 {object} error.ApiError
// @Failure 403 {object} error.ApiError
// @Failure 500 {object} error.ApiError
// @Router /courses/{id}/{version} [get]
func (c *CourseController) Course(ctx echo.Context) error {
	// Read id parameter from request
	id := c.GetPathParamInt(ctx, "id")
	if id == -1 {
		return e.ErrorInvalidRequest("course id")
	}
	// Read version parameter from request
	version := c.GetPathParamInt(ctx, "version")
	if version == -1 {
		return e.ErrorInvalidRequest("course version")
	}
	// Ask the DAO for the course
	res, err := c.Dao.Get(uint(id), uint(version))
	if err != nil {
		return err
	}
	// Return the result to the client
	return c.Yeet(ctx, res)
}
