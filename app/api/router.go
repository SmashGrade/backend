package api

import (
	"github.com/SmashGrade/backend/app/auth"
	"github.com/SmashGrade/backend/app/db"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// Router is the main router for the API
type Router struct {
	echo   *echo.Echo
	auth   *auth.AuthProvider
	course *CourseController
	output *OutputController
}

// NewRouter creates a new router
func NewRouter(e *echo.Echo, p db.Provider) *Router {
	return &Router{
		echo:   e,
		auth:   auth.NewAuthProvider(p.Config()),
		course: NewCourseController(p),
		output: NewOutputController(p),
	}
}

// Register all v1 routes

func (r *Router) RegisterV1() {
	// Create a new group for v1
	v1 := r.echo.Group("/v1")
	// Enable authentication for v1 endpoints
	v1.Use(echojwt.WithConfig(r.auth.GetJWTConfig()))
	// Register all v1 routes
	RegisterV1Courses(v1, r.course)
	RegisterV1Output(v1, r.output)
}
