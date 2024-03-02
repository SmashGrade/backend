package api

import (
	"github.com/SmashGrade/backend/app/dao"
	"github.com/SmashGrade/backend/app/db"
	e "github.com/SmashGrade/backend/app/error"
	"github.com/SmashGrade/backend/app/repository"
	"github.com/SmashGrade/backend/app/requestmodels"
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
		Dao: dao.NewCourseDao(
			repository.NewCourseRepository(provider),
			repository.NewModuleRepository(provider),
			repository.NewUserRepository(provider),
			repository.NewSelectedCourseRepository(provider),
			repository.NewExamRepository(provider),
		),
	}
}

//	@Summary		Get all courses
//	@Description	Get all courses
//	@Tags			courses
//	@Produce		json
//	@Success		200	{array}		models.Course
//	@Failure		401	{object}	error.ApiError
//	@Failure		403	{object}	error.ApiError
//	@Failure		500	{object}	error.ApiError
//	@Router			/courses [get]
//	@Security		Bearer
func (c *CourseController) Courses(ctx echo.Context) error {
	res, err := c.Dao.GetAll()
	if err != nil {
		return err
	}
	return c.Yeet(ctx, res)
}

//	@Summary		Get a specific course
//	@Description	Get a specific course
//	@Tags			courses
//	@Param			id		path	uint	true	"Course ID"
//	@Param			version	path	uint	true	"Course Version"
//	@Produce		json
//	@Success		200	{object}	models.Course
//	@Failure		401	{object}	error.ApiError
//	@Failure		403	{object}	error.ApiError
//	@Failure		500	{object}	error.ApiError
//	@Router			/courses/{id}/{version} [get]
//	@Security		Bearer
func (c *CourseController) Course(ctx echo.Context) error {
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

	// Ask the DAO for the course
	res, err := c.Dao.Get(id, version)
	if err != nil {
		return err
	}
	// Return the result to the client
	return c.Yeet(ctx, res)
}

//	@Summary		Create a course
//	@Description	Create a course
//	@Tags			courses
//	@Produce		json
//	@Accept			json
//
//	@Param			request	body		requestmodels.RefCourse	true	"request body"
//
//	@Success		200		{object}	models.Course
//	@Failure		401		{object}	error.ApiError
//	@Failure		403		{object}	error.ApiError
//	@Failure		500		{object}	error.ApiError
//	@Router			/courses [post]
//	@Security		Bearer
func (c *CourseController) Create(ctx echo.Context) error {
	course := new(requestmodels.RefCourse)
	// Read the request into Course
	if err := ctx.Bind(course); err != nil {
		return e.ErrorInvalidRequest("course")
	}
	// Let dao create the Course
	returnCourse, err := c.Dao.Create(*course)
	if err != nil {
		return err
	}
	// return the result from the Post
	return c.Yeet(ctx, returnCourse)
}

//	@Summary		Create a new version of a course
//	@Description	Create a new version of a course
//	@Tags			courses
//	@Produce		json
//	@Accept			json
//
//	@Param			request	body		requestmodels.RefCourse	true	"request body"
//	@Param			id		path		uint					true	"Course ID"
//
//	@Success		200		{object}	models.Course
//	@Failure		401		{object}	error.ApiError
//	@Failure		403		{object}	error.ApiError
//	@Failure		500		{object}	error.ApiError
//	@Router			/courses/{id} [post]
//	@Security		Bearer
func (c *CourseController) CreateVersion(ctx echo.Context) error {
	// Read id parameter from request

	id, err := c.GetPathParamUint(ctx, "id")
	if err != nil {
		return e.NewDaoValidationError("id", "uint", c.GetPathParam(ctx, "id"))
	}

	course := new(requestmodels.RefCourse)
	// Read the request into Course
	if err := ctx.Bind(course); err != nil {
		return e.ErrorInvalidRequest("course")
	}

	course.ID = id
	course.Version = 0

	// Let dao create the Course
	returnCourse, err := c.Dao.Create(*course)
	if err != nil {
		return err
	}
	// return the result from the Post
	return c.Yeet(ctx, returnCourse)
}

//	@Summary		Update a course
//	@Description	Update a course
//	@Tags			courses
//	@Produce		json
//	@Accept			json
//
//	@Param			request	body		requestmodels.RefCourse	true	"request body"
//	@Param			id		path		uint					true	"Course ID"
//	@Param			version	path		uint					true	"Course Version"
//
//	@Success		200		{object}	models.Course
//	@Failure		401		{object}	error.ApiError
//	@Failure		403		{object}	error.ApiError
//	@Failure		500		{object}	error.ApiError
//	@Router			/courses/{id}/{version} [put]
//	@Security		Bearer
func (c *CourseController) Update(ctx echo.Context) error {
	// Read id parameter from request
	id, err := c.GetPathParamUint(ctx, "id")
	if err != nil {
		return e.NewDaoValidationError("id", "uint", c.GetPathParam(ctx, "id"))
	}

	version, err := c.GetPathParamUint(ctx, "version")
	if err != nil {
		return e.NewDaoValidationError("version", "uint", c.GetPathParam(ctx, "version"))
	}

	course := new(requestmodels.RefCourse)
	// Read the request into Course
	if err := ctx.Bind(course); err != nil {
		return e.ErrorInvalidRequest("course")
	}

	course.ID = id
	course.Version = version

	// Let dao create the Course
	err = c.Dao.Update(*course)
	if err != nil {
		return err
	}
	// return the result from the Post
	return c.Yeet(ctx, course)
}

//	@Summary		Delete a course
//	@Description	Delete a course
//	@Tags			courses
//	@Produce		json
//	@Accept			json
//
//	@Param			id		path	uint	true	"Course ID"
//	@Param			version	path	uint	true	"Course Version"
//
//	@Success		200
//	@Failure		401	{object}	error.ApiError
//	@Failure		403	{object}	error.ApiError
//	@Failure		500	{object}	error.ApiError
//	@Router			/courses/{id}/{version} [delete]
//	@Security		Bearer
func (c *CourseController) Delete(ctx echo.Context) error {
	// Read id parameter from request

	id, err := c.GetPathParamUint(ctx, "id")
	if err != nil {
		return e.NewDaoValidationError("id", "uint", c.GetPathParam(ctx, "id"))
	}

	version, err := c.GetPathParamUint(ctx, "version")
	if err != nil {
		return e.NewDaoValidationError("version", "uint", c.GetPathParam(ctx, "version"))
	}

	// Let dao create the Course
	err = c.Dao.Delete(id, version)
	if err != nil {
		return err
	}
	// return the result from the Post
	return c.Yeet(ctx, nil)
}

// register all output endpoints to router
func RegisterV1Courses(g *echo.Group, c *CourseController) {
	g.GET("/courses", c.Courses)
	g.GET("/courses/:id/:version", c.Course)
	g.POST("/courses", c.Create)
	g.POST("/courses/:id", c.CreateVersion)
	g.PUT("/courses/:id/:version", c.Update)
	g.DELETE("/courses/:id/:version", c.Delete)
}
