package dao

import (
	"time"

	e "github.com/SmashGrade/backend/app/error"
	"github.com/SmashGrade/backend/app/models"
)

// curriculum type / Studiengang art
// has description like Vollzeit or Berufsbegleitend
type CurriculumTypeDao struct {
}

// Creates new dao with required repositories
func NewCurriculumTypeDao(curriculumTypeRepository *interface{}) (dao *CurriculumTypeDao, err e.DaoError) {
	return nil, e.DAOUnimplemented
}

func (c *CurriculumTypeDao) GetAll() (entities []models.Curriculumtype, err e.DaoError) {
	return nil, e.DAOUnimplemented
}

// Curriculum / Studiengang
// Highest level of categorization
type CurriculumDao struct{}

// Returns existing curriculum
func (c *CurriculumDao) Get(id uint, startValidity time.Time) (entity *models.Curriculum, err e.DaoError) {
	return nil, e.DAOUnimplemented
}

// Creates new curriculum
func (c *CurriculumDao) Create(entity *models.Curriculum) e.DaoError {
	return e.DAOUnimplemented
}

// Module / Modul
// A collection of multiple courses
type ModuleDao struct{}

// Creates a new dao with required repositories
func NewModuleDao(moduleRepository *interface{}) (dao *ModuleDao, err e.DaoError) {
	return nil, e.DAOUnimplemented
}

// Returns module identified by id and version
func (m *ModuleDao) Get(id, version uint) (entity *models.Module, err e.DaoError) {
	return nil, e.DAOUnimplemented
}

// Returns module by id with highest version
func (m *ModuleDao) GetLatest(id uint) (entity *models.Module, err e.DaoError) {
	return nil, e.DAOUnimplemented
}

// Will create a new module if neither id nor version are set
// Will create a new module version if only id is set
func (m *ModuleDao) Create(entity models.Module) (returnEntity *models.Module, err e.DaoError) {
	return nil, e.DAOUnimplemented
}

// Will update an existing module specified by id and version
func (m *ModuleDao) Update(entity models.Module) e.DaoError {
	return e.DAOUnimplemented
}

// Deletes a single module with id and version
func (m *ModuleDao) Delete(id, version uint) e.DaoError {
	return e.DAOUnimplemented
}

type CourseDao struct{}

// Create new dao with required repositories
func NewCourseDao(courseRepository *interface{}) (dao *CourseDao, err e.DaoError) {
	return nil, e.DAOUnimplemented
}

// Returns course by id and version
func (c *CourseDao) Get(id, version uint) (entity *models.Course, err e.DaoError) {
	return nil, e.DAOUnimplemented
}

// Returns course by id with highest version
func (c *CourseDao) GetLatest(id uint) (entity *models.Course, err e.DaoError) {
	return nil, e.DAOUnimplemented
}

// Will create a new course if neither id nor version are set
// Will create a new course version if only id is set
func (c *CourseDao) Create(entity models.Course) (returnEntity *models.Course, err e.DaoError) {
	return nil, e.DAOUnimplemented
}

// updates an existing course via id and version
func (c *CourseDao) Update(entity models.Course) e.DaoError {
	return e.DAOUnimplemented
}

// Deletes a course by id and version
func (c *CourseDao) Delete(id, version uint) e.DaoError {
	return e.DAOUnimplemented
}

type ExamDao struct{}

// Create new exam dao with all used providers
func NewDoaExam(examProvider *interface{}, courseProvider *interface{}) (daoExam *ExamDao, err e.DaoError) {
	return nil, e.DAOUnimplemented
}

// Returns a list of exams for a specific course
func (ex *ExamDao) GetForCourse(courseId, courseVersion uint) (entities []models.Exam, err e.DaoError) {
	return nil, e.DAOUnimplemented
}

// Returns a single exam for a course selected by id
func (ex *ExamDao) Get(courseId, courseVersion, examId uint) (entity *models.Exam, err e.DaoError) {
	return nil, e.DAOUnimplemented
}

// Creates a new exam for a course
func (ex *ExamDao) Create(entity models.Exam) (returnEntity *models.Exam, err e.DaoError) {
	return nil, e.DAOUnimplemented
}

// Updates an existing exam for a course
func (ex *ExamDao) Update(entity models.Exam) e.DaoError {
	return e.DAOUnimplemented
}

// Deletes exam for a course
func (ex *ExamDao) Delete(courseId, courseVersion, examId uint) e.DaoError {
	return e.DAOUnimplemented
}

type UserDao struct{}

// Creates new dao from required repositories
func NewUserDao(userRepository *interface{}) (dao *UserDao, err e.DaoError) {
	return nil, e.DAOUnimplemented
}

// Returns a list of courses a user has assigned
func (u *UserDao) GetCourses(uid uint) (courses []models.SelectedCourse, err e.DaoError) {
	return nil, e.DAOUnimplemented
}

// Returns a list of courses a user has assigned in a specific start year
func (u *UserDao) GetCoursesForYear(uid uint, startYear time.Time) (courses []models.SelectedCourse, err e.DaoError) {
	return nil, e.DAOUnimplemented
}

func (u *UserDao) GetExamEvaluations()
