package v1

import (
	"net/http"

	"github.com/SmashGrade/backend/app/api"
	"github.com/SmashGrade/backend/app/api/v1/controllers"
	"github.com/labstack/echo/v4"
)

func RoutesV1(ctx *api.SetupContext) {
	v1 := ctx.Echo.Group("/api/v1")
	v1.GET("/apples", apples)

	// Curriculum
	v1.GET("/curriculum", controllers.GetCurriculums)
	v1.POST("/curriculum", controllers.PostCurriculum)
	v1.GET("/curriculum/:id", controllers.GetCurriculum)
	v1.PUT("/curriculum/:id", controllers.PutCurriculum)
	v1.DELETE("/curriculum/:id", controllers.DeleteCurriculum)
	v1.GET("/curriculum/filter", controllers.GetCurriculumFilter)
	// Module
	v1.GET("/module", controllers.GetModules)
	v1.POST("/module", controllers.PostModule)
	v1.GET("/module/:id", controllers.GetModule)
	v1.PUT("/module/:id", controllers.PutModule)
	v1.DELETE("/module/:id", controllers.DeleteModule)
	v1.GET("/module/student/:studyStageId", controllers.GetModuleStudent)
	v1.GET("/module/teacher/:studyStageId", controllers.GetModuleTeacher)
	v1.GET("/module/filter", controllers.GetModuleFilter)
	// Course
	v1.GET("/course", controllers.GetCourses)
	v1.POST("/course", controllers.PostCourse)
	v1.GET("/course/:id", controllers.GetCourse)
	v1.PUT("/course/:id", controllers.PutCourse)
	v1.DELETE("/course/:id", controllers.DeleteCourse)
	v1.GET("/course/:id/student", controllers.GetCourseStudent)
	v1.GET("/course/:id/teacher", controllers.GetCourseTeacher)
	v1.GET("/course/filter", controllers.GetCourseFilter)
	// Exam
	v1.GET("/exam", controllers.GetExams)
	v1.POST("/exam", controllers.PostExam)
	v1.GET("/exam/:id", controllers.GetExam)
	v1.PUT("/exam/:id", controllers.PutExam)
	v1.DELETE("/exam/:id", controllers.DeleteExam)
	v1.POST("/exam/:id/student", controllers.PostExamGradeStudent)
	v1.POST("exam/:id/teacher", controllers.PostExamGradeTeacher)
	// Onboarding
	v1.POST("/onboarding", controllers.PostOnboarding)
	v1.PUT("/onboarding/:id", controllers.PutOnboarding)
	v1.GET("/onboarding/filter", controllers.GetOnboardingFilter)
	// User
	v1.GET("/user", controllers.GetUsers)
	v1.POST("/user", controllers.PostUser)
	v1.GET("/user/:id", controllers.GetUser)
	v1.PUT("/user/:id", controllers.PutUser)
	v1.GET("/user/:id/course", controllers.GetUserCourses)
	v1.GET("/user/:id/exam", controllers.GetUserExams)
	v1.GET("/user/:id/exam/:examId", controllers.GetUserExam)

}

func apples(c echo.Context) error {
	err := c.String(http.StatusOK, "Lululu i've got some Apples")
	if err != nil {
		return err
	}
	return nil
}
