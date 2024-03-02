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
	// returns: all users, all modules, all examtypes
	// These are Preselected Items
	var metaCourse models.MetaCourse

	// TODO: only allow course admin role on this endpoint

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
	// returns: all evaluation types, all curriculum, all curriculum types, all courses
	// These are Preselected Items
	var metaModules models.MetaModules

	// TODO: only allow course admin role on this endpoint

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

	// Get all courses
	courses, err := r.course.Dao.GetAll()
	if err != nil {
		return err
	}

	metaModules.Evaluationtypes = evaluationtypes
	metaModules.Curriculums = curriculums
	metaModules.Curriculumstype = curriculumstype
	metaModules.Courses = courses

	return r.module.Yeet(ctx, metaModules)
}

func (r *Router) MetaCurriculums(ctx echo.Context) error {
	// MetaCurriculum contains all form choice data to create or modify a curriculum (Studiengang)
	// returns: all focus (Fachrichtung), all fields (Schwerpunkt), all curriculumtypes, all users

	// TODO: only allow course admin role on this endpoint

	return nil
}

func (r *Router) MyCoursesAsTeacher(ctx echo.Context) error {
	// View of the course teacher
	// returns: list of courses teached by current user with modules and study stage, list of all users
	// TODO: check if user is teacher

	return nil
}

func (r *Router) MyCurriculumsAsStudent(ctx echo.Context) error {
	// View of the student, general info
	// returns: chosen curriculum with start year and curriculum type
	// TODO: check if user is student

	return nil
}

// register all output endpoints to router
func RegisterV1MetaCourse(g *echo.Group, r *Router) {
	g.GET("/courses/meta", r.MetaCourses)
	g.GET("/modules/meta", r.MetaModules)
	g.GET("/curriculums/meta", r.MetaCurriculums)
	g.GET("/courses/me", r.MyCoursesAsTeacher)
	g.GET("/curriculums/me", r.MetaCurriculums)
}
