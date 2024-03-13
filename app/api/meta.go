package api

import (
	"time"

	"github.com/SmashGrade/backend/app/config"
	"github.com/SmashGrade/backend/app/dao"
	"github.com/SmashGrade/backend/app/db"
	e "github.com/SmashGrade/backend/app/error"
	"github.com/SmashGrade/backend/app/models"
	"github.com/SmashGrade/backend/app/repository"
	"github.com/labstack/echo/v4"
)

// Handles the requests for the meta resources
type MetaController struct {
	*BaseController
	userDao           *dao.UserDao
	moduleDao         *dao.ModuleDao
	courseDao         *dao.CourseDao
	curriculumDao     *dao.CurriculumDao
	examtypeDao       *dao.ExamtypeDao
	evaluationtypeDao *dao.EvaluationTypeDao
	curriculumtypeDao *dao.CurriculumTypeDao
	focusDao          *dao.FocusDao
	fieldDao          *dao.FieldDao
	classDao          *dao.ClassDao
}

// Constructor for MetaController
func NewMetaController(provider db.Provider) *MetaController {
	return &MetaController{
		BaseController:    NewBaseController(provider),
		userDao:           dao.NewUserDao(repository.NewUserRepository(provider), repository.NewRoleRepository(provider)),
		moduleDao:         dao.NewModuleDao(repository.NewModuleRepository(provider)),
		courseDao:         dao.NewCourseDao(repository.NewCourseRepository(provider), repository.NewModuleRepository(provider), repository.NewUserRepository(provider), repository.NewSelectedCourseRepository(provider), repository.NewExamRepository(provider), repository.NewRoleRepository(provider)),
		curriculumDao:     dao.NewCurriculumDao(repository.NewCurriculumRepository(provider), repository.NewFocusRepository(provider), repository.NewCurriculumtypeRepository(provider), repository.NewStateRepository(provider), repository.NewModuleRepository(provider)),
		examtypeDao:       dao.NewExamtypeDao(repository.NewExamtypeRepository(provider)),
		evaluationtypeDao: dao.NewEvaluationTypeDao(repository.NewEvaluationtypeRepository(provider)),
		curriculumtypeDao: dao.NewCurriculumTypeDao(repository.NewCurriculumtypeRepository(provider)),
		focusDao:          dao.NewFocusDao(repository.NewFocusRepository(provider)),
		fieldDao:          dao.NewFieldDao(repository.NewFieldRepository(provider)),
		classDao:          dao.NewClassDao(repository.NewCourseRepository(provider), repository.NewUserRepository(provider), repository.NewSelectedCourseRepository(provider), repository.NewExamRepository(provider), repository.NewRoleRepository(provider), repository.NewExamEvaluationRepository(provider)),
	}
}

// @Summary		Get Metadata for Course
// @Description	Get Metadata for Course
// @Tags			meta, courses
// @Produce		json
// @Success		200	{array}		models.MetaCourse
// @Failure		401	{object}	error.ApiError
// @Failure		403	{object}	error.ApiError
// @Failure		500	{object}	error.ApiError
// @Router			/courses/meta [get]
// @Security		Bearer
func (m *MetaController) MetaCourses(ctx echo.Context) error {
	// MetaCourse has all the data the Frontend needs to Create or Modify a Course.
	// returns: all users, all modules, all examtypes
	// These are Preselected Items
	var metaCourse models.MetaCourse

	err := m.CheckUserRole(config.ROLE_COURSEADMIN, ctx)
	if err != nil {
		return err
	}

	// Get all Teachers
	teachers, err := m.userDao.GetTeachers()
	if err != nil {
		return err
	}
	// Get all Modules
	modules, err := m.moduleDao.GetAll()
	if err != nil {
		return err
	}
	// Get all ExamTypes
	examtypes, err := m.examtypeDao.GetAll()
	if err != nil {
		return err
	}

	metaCourse.Teachers = teachers
	metaCourse.Modules = modules
	metaCourse.Examtypes = examtypes

	return m.Yeet(ctx, metaCourse)
}

// @Summary		Get Metadata for Module
// @Description	Get Metadata for Module
// @Tags			meta, modules
// @Produce		json
// @Success		200	{array}		models.MetaModules
// @Failure		401	{object}	error.ApiError
// @Failure		403	{object}	error.ApiError
// @Failure		500	{object}	error.ApiError
// @Router			/modules/meta [get]
// @Security		Bearer
func (m *MetaController) MetaModules(ctx echo.Context) error {
	// MetaModules has all the data the Frontend needs to Create or Modify a Module.
	// returns: all evaluation types, all curriculum, all curriculum types, all courses
	// These are Preselected Items
	var metaModules models.MetaModules

	err := m.CheckUserRole(config.ROLE_COURSEADMIN, ctx)
	if err != nil {
		return err
	}

	// Get all Evaluationtypes
	evaluationtypes, err := m.evaluationtypeDao.GetAll()
	if err != nil {
		return err
	}

	// Get all Curriculums
	curriculums, err := m.curriculumDao.GetAll()
	if err != nil {
		return err
	}

	// Get all Curriculumtypes
	curriculumtypes, err := m.curriculumtypeDao.GetAll()
	if err != nil {
		return err
	}

	// Get all courses
	courses, err := m.courseDao.GetAll()
	if err != nil {
		return err
	}

	metaModules.Evaluationtypes = evaluationtypes
	metaModules.Curriculums = curriculums
	metaModules.Curriculumtypes = curriculumtypes
	metaModules.Courses = courses

	return m.Yeet(ctx, metaModules)
}

// @Summary		Get Metadata for Curriculums
// @Description	Get Metadata for Curriculums
// @Tags			meta, curriculums
// @Produce		json
// @Success		200	{array}		models.MetaCurriculums
// @Failure		401	{object}	error.ApiError
// @Failure		403	{object}	error.ApiError
// @Failure		500	{object}	error.ApiError
// @Router			/curriculums/meta [get]
// @Security		Bearer
func (m *MetaController) MetaCurriculums(ctx echo.Context) error {
	// MetaCurriculum contains all form choice data to create or modify a curriculum (Studiengang)
	// returns: all focus (Fachrichtung), all fields (Schwerpunkt), all curriculumtypes, all users

	err := m.CheckUserRole(config.ROLE_COURSEADMIN, ctx)
	if err != nil {
		return err
	}

	var metaCurriculums models.MetaCurriculums

	// Get all Teachers
	teachers, err := m.userDao.GetTeachers()
	if err != nil {
		return err
	}

	// Get all Curriculumtypes
	curriculumtypes, err := m.curriculumtypeDao.GetAll()
	if err != nil {
		return err
	}

	// get all focus
	focuses, err := m.focusDao.GetAll()
	if err != nil {
		return err
	}

	// get all fields
	fields, err := m.fieldDao.GetAll()
	if err != nil {
		return err
	}

	metaCurriculums.Teachers = teachers
	metaCurriculums.Curriculumtypes = curriculumtypes
	metaCurriculums.Focuses = focuses
	metaCurriculums.Fields = fields

	return m.Yeet(ctx, metaCurriculums)
}

// @Summary		Get your Courses as a teacher
// @Description	Get Courses as a teacher selected by teached by userinfo from accesstoken
// @Tags			meta, courses
// @Produce		json
// @Success		200	{array}		models.TeacherCourses
// @Failure		401	{object}	error.ApiError
// @Failure		403	{object}	error.ApiError
// @Failure		500	{object}	error.ApiError
// @Router			/courses/teacher [get]
// @Security		Bearer
func (m *MetaController) MyCoursesAsTeacher(ctx echo.Context) error {
	// View of the course teacher
	// returns: list of courses teached by current user with modules and study stage, list of all users

	err := m.CheckUserRole(config.ROLE_TEACHER, ctx)
	if err != nil {
		return err
	}

	user, err := m.GetUser(ctx)
	if err != nil {
		return err
	}

	var teacherCourses models.TeacherCourses

	teacherCourses.Courses = make([]models.Course, len(user.TeachesCourses))
	teacherCourses.Classes = make([]models.Class, 0)
	for i := range user.TeachesCourses {
		teacherCourses.Courses[i] = *user.TeachesCourses[i]

		// here we create a list of all unique class start dates in selected courses
		uniqueStartDates := make([]time.Time, 0)
		for _, newSelected := range user.TeachesCourses[i].SelectedCourses {
			alreadyExists := false
			for _, existingDate := range uniqueStartDates {
				if newSelected.ClassStartyear == existingDate {
					alreadyExists = true
					break
				}
			}
			if !alreadyExists {
				uniqueStartDates = append(uniqueStartDates, newSelected.ClassStartyear)
			}
		}

		for _, uniqueDate := range uniqueStartDates {
			newClass, err := m.classDao.Get(teacherCourses.Courses[i].ID, teacherCourses.Courses[i].Version, uniqueDate)
			if err != nil {
				return err
			}
			teacherCourses.Classes = append(teacherCourses.Classes, *newClass)
		}
	}

	return m.Yeet(ctx, teacherCourses)
}

// @Summary		Get Curriculums as a student
// @Description	Get Curriculums as a student selected by userinfo from accesstoken
// @Tags			meta, curriculums
// @Produce		json
// @Success		200	{array}		models.StudentCurriculums
// @Failure		401	{object}	error.ApiError
// @Failure		403	{object}	error.ApiError
// @Failure		500	{object}	error.ApiError
// @Router			/curriculums/student [get]
// @Security		Bearer
func (m *MetaController) MyCurriculumsAsStudent(ctx echo.Context) error {
	// View of the student, general info
	// returns: chosen curriculum with start year and curriculum type

	err := m.CheckUserRole(config.ROLE_STUDENT, ctx)
	if err != nil {
		return err
	}

	user, err := m.GetUser(ctx)
	if err != nil {
		return err
	}

	studentCurriculum := models.StudentCurriculums{
		StartYear: user.ClassStartyear,
	}

	curriculum, err := m.curriculumDao.GetValidForTimepoint(user.CurriculumID, studentCurriculum.StartYear)
	if err != nil {
		return err
	}

	studentCurriculum.Curriculum = *curriculum

	return m.Yeet(ctx, studentCurriculum)
}

// @Summary		Set start year and curriculumId as student
// @Description	Set start year and curriculumId as student by userinfo from accesstoken
// @Tags			meta, curriculums, users
// @Param			id		path	uint		true	"Curriculum ID"
// @Param			date	path	time.Time	true	"Class start date"
// @Produce		json
// @Success		200	{array}		models.User
// @Failure		401	{object}	error.ApiError
// @Failure		403	{object}	error.ApiError
// @Failure		500	{object}	error.ApiError
// @Router			/onboarding [put]
// @Security		Bearer
func (m *MetaController) SetStudentCurriculumLink(ctx echo.Context) error {
	err := m.CheckUserRole(config.ROLE_STUDENT, ctx)
	if err != nil {
		return err
	}

	user, err := m.GetUser(ctx)
	if err != nil {
		return err
	}

	// Read id parameter from request
	id, intErr := m.GetPathParamUint(ctx, "id")
	if intErr != nil {
		return e.NewDaoValidationError("id", "uint", m.GetPathParam(ctx, "id"))
	}

	// Read date paramater from request
	date, intErr := m.GetPathParamTime(ctx, "date")
	if intErr != nil {
		return e.NewDaoValidationError("date", "time.Time", m.GetPathParam(ctx, "date"))
	}

	// check if there is such a curriculum
	_, tpErr := m.curriculumDao.GetValidForTimepoint(id, date)
	if tpErr != nil {
		return tpErr
	}

	user.CurriculumID = id
	user.ClassStartyear = date

	updErr := m.userDao.Update(*user)
	if updErr != nil {
		return updErr
	}

	return m.Yeet(ctx, *user)
}

// register all output endpoints to router
func RegisterV1MetaCourse(g *echo.Group, m *MetaController) {
	g.GET("/courses/meta", m.MetaCourses)
	g.GET("/modules/meta", m.MetaModules)
	g.GET("/curriculums/meta", m.MetaCurriculums)
	g.GET("/courses/teacher", m.MyCoursesAsTeacher)
	g.GET("/curriculums/student", m.MyCurriculumsAsStudent)
	g.PUT("/onboarding", m.SetStudentCurriculumLink)
}
