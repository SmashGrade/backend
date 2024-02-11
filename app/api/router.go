package api

import (
	"github.com/SmashGrade/backend/app/db"
	"github.com/labstack/echo/v4"
)

// Router is the main router for the API
type Router struct {
	echo   *echo.Echo
	course *CourseController
}

// NewRouter creates a new router
func NewRouter(e *echo.Echo, p db.Provider) *Router {
	return &Router{
		echo:   e,
		course: NewCourseController(p),
	}
}

// Register all v1 routes
func (r *Router) RegisterV1() {
	// Create a new group for v1
	v1 := r.echo.Group("/v1")
	// Register all v1 routes
	v1.GET("/courses", r.course.Courses)
	v1.GET("/courses/:id/:version", r.course.Course)
}
