package api

import (
	"log"

	"github.com/SmashGrade/backend/app/auth"
	"github.com/SmashGrade/backend/app/config"
	"github.com/SmashGrade/backend/app/db"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// Router is the main router for the API
type Router struct {
	echo   *echo.Echo
	config *config.APIConfig
	auth   *auth.AuthProvider
	course *CourseController
	output *OutputController
	user   *UserController
	exam   *ExamController
}

// NewRouter creates a new router
func NewRouter(e *echo.Echo, p db.Provider) *Router {
	return &Router{
		echo:   e,
		config: p.Config(),
		auth:   auth.NewAuthProvider(p.Config()),
		course: NewCourseController(p),
		output: NewOutputController(p),
		user:   NewUserController(p),
		exam:   NewExamController(p),
	}
}

// Register all v1 routes

func (r *Router) RegisterV1() {
	// Create a new group for v1
	v1 := r.echo.Group("/v1")
	if r.config.AuthConfig.Enabled {
		log.Println("Enabling authentication for v1 endpoints...")
		// Enable authentication for v1 endpoints
		v1.Use(echojwt.WithConfig(r.auth.GetJWTConfig()))
	} else {
		log.Println("Authentication is disabled for v1 endpoints...")
		log.Println("This is not recommended for production environments!")
	}
	// Register all v1 routes
	RegisterV1Courses(v1, r.course)
	RegisterV1Output(v1, r.output)
	RegisterV1User(v1, r.user)
	RegisterV1Exams(v1, r.exam)
}
