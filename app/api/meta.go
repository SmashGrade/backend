package api

import (
	"github.com/SmashGrade/backend/app/models"
	"github.com/labstack/echo/v4"
)

// @Summary		Get Metadata for Course
// @Description	Get Metadata for Course
// @Tags			Meta
// @Produce		json
// @Success		200	{array}		models.MetaCourse
// @Failure		401	{object}	error.ApiError
// @Failure		403	{object}	error.ApiError
// @Failure		500	{object}	error.ApiError
// @Router			/meta/courses [get]
// @Security		Bearer
func (r *Router) MetaCourses(ctx echo.Context) error {
	// MetaCourse has all the data the Frontend needs to Create or Modify a Course.
	// These are Preselected Items
	var metaCourse models.MetaCourse

	// Get all Teachers
	teachers, err := r.user.UserDao.GetTeachers()
	if err != nil {
		return err
	}
	// Get all Modules
	modules, err := r.module.Dao.GetAll()
	if err != nil {
		return err
	}
	// Get all ExamTypes
	examtypes, err := r.output.examtypeDao.GetAll()
	if err != nil {
		return err
	}

	metaCourse.Teachers = teachers
	metaCourse.Modules = modules
	metaCourse.Examtypes = examtypes

	return r.course.Yeet(ctx, metaCourse)
}

// @Summary		Get Metadata for Module
// @Description	Get Metadata for Module
// @Tags			meta
// @Produce		json
// @Success		200	{array}		models.MetaModules
// @Failure		401	{object}	error.ApiError
// @Failure		403	{object}	error.ApiError
// @Failure		500	{object}	error.ApiError
// @Router			/meta/modules [get]
// @Security		Bearer
func (r *Router) MetaModules(ctx echo.Context) error {
	// MetaModules has all the data the Frontend needs to Create or Modify a Module.
	// These are Preselected Items
	var metaModules models.MetaModules

	// Get all Evaluationtypes
	evaluationtypes, err := r.output.evaluationtypeDao.GetAll()
	if err != nil {
		return err
	}

	// Get all Curriculums
	curriculums, err := r.curriculum.Dao.GetAll()
	if err != nil {
		return err
	}

	// Get all Curriculumtypes
	curriculumstype, err := r.output.curriculumytypeDao.GetAll()
	if err != nil {
		return err
	}

	metaModules.Evaluationtypes = evaluationtypes
	metaModules.Curriculums = curriculums
	metaModules.Curriculumstype = curriculumstype

	return r.module.Yeet(ctx, metaModules)
}

// register all output endpoints to router
func RegisterV1MetaCourse(g *echo.Group, r *Router) {
	g.GET("/meta/courses", r.MetaCourses)
	g.GET("/meta/modules", r.MetaModules)
}
