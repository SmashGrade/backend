package v1

import (
	"net/http"

	"github.com/SmashGrade/backend/legacy/api"
	c "github.com/SmashGrade/backend/legacy/api/v1/controllers"
	"github.com/SmashGrade/backend/legacy/provider"
	"github.com/labstack/echo/v4"
)

func RoutesV1(ctx *api.SetupContext) {
	prov := &provider.SqliteProvider{}
	prov.Connect()

	v1 := ctx.Echo.Group("/api/v1")
	v1.GET("/apples", apples)

	// Curriculum
	v1.GET("/curriculum", c.GetCurriculums)
	v1.POST("/curriculum", c.PostCurriculum)
	v1.GET("/curriculum/:id", c.GetCurriculum)
	v1.PUT("/curriculum/:id", c.PutCurriculum)
	v1.DELETE("/curriculum/:id", c.DeleteCurriculum)
	v1.GET("/curriculum/filter", c.GetCurriculumFilter)
	// Module
	v1.GET("/module", c.GetModules)
	v1.POST("/module", c.PostModule)
	v1.GET("/module/:id", c.GetModule)
	v1.PUT("/module/:id", c.PutModule)
	/* v1.DELETE("/module/:id", c.DeleteModule) */
	v1.GET("/module/student/:studyStageId", c.GetModuleStudent)
	v1.GET("/module/teacher/:studyStageId", c.GetModuleTeacher)
	v1.GET("/module/filter", c.GetModuleFilter)
	// Course
	v1.GET("/course", c.GetCourses)
	v1.POST("/course", c.PostCourse)
	v1.GET("/course/:id", c.GetCourse)
	v1.PUT("/course/:id", c.PutCourse)
	/* v1.DELETE("/course/:id", c.DeleteCourse) */
	v1.GET("/course/:id/student", c.GetCourseStudent)
	v1.GET("/course/:id/teacher", c.GetCourseTeacher)
	v1.GET("/course/filter", c.GetCourseFilter)
	// Exam
	v1.GET("/exam", c.GetExams)
	v1.POST("/exam", c.PostExam)
	v1.GET("/exam/:id", c.GetExam)
	v1.PUT("/exam/:id", c.PutExam)
	v1.DELETE("/exam/:id", c.DeleteExam)
	v1.POST("/exam/:id/student", c.PostExamGradeStudent)
	v1.POST("exam/:id/teacher", c.PostExamGradeTeacher)
	// Onboarding
	v1.POST("/onboarding", c.PostOnboarding)
	v1.PUT("/onboarding/:id", c.PutOnboarding)
	v1.GET("/onboarding/filter", c.GetOnboardingFilter)
	// User
	v1.GET("/user", c.GetUsers)
	v1.POST("/user", c.PostUser)
	v1.GET("/user/:id", c.GetUser)
	v1.PUT("/user/:id", c.PutUser)
	v1.GET("/user/:id/course", c.GetUserCourses)
	v1.GET("/user/:id/exam", c.GetUserExams)
	v1.GET("/user/:id/exam/:examId", c.GetUserExam)

}

func apples(c echo.Context) error {
	err := c.String(http.StatusOK, "Lululu i've got some Apples")
	if err != nil {
		return err
	}
	return nil
}
