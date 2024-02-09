package dao

import (
	"time"

	e "github.com/SmashGrade/backend/app/error"
	"github.com/SmashGrade/backend/app/models"
	"github.com/SmashGrade/backend/app/repository"
)

// curriculum type / Studiengang art
// has description like Vollzeit or Berufsbegleitend
type CurriculumTypeDao struct {
}

// Creates new dao with required repositories
func NewCurriculumTypeDao(curriculumTypeRepository *interface{}) (dao *CurriculumTypeDao, err *e.ApiError) {
	return nil, e.NewDaoUnimplementedError()
}

func (c *CurriculumTypeDao) GetAll() (entities []models.Curriculumtype, err *e.ApiError) {
	return nil, e.NewDaoUnimplementedError()
}

// Curriculum / Studiengang
// Highest level of categorization
type CurriculumDao struct{}

// Returns existing curriculum
func (c *CurriculumDao) Get(id uint, startValidity time.Time) (entity *models.Curriculum, err *e.ApiError) {
	return nil, e.NewDaoUnimplementedError()
}

// Creates new curriculum
func (c *CurriculumDao) Create(entity *models.Curriculum) e.DaoError {
	return e.NewDaoUnimplementedError()
}

// Module / Modul
// A collection of multiple courses
type ModuleDao struct{}

// Creates a new dao with required repositories
func NewModuleDao(moduleRepository *interface{}) (dao *ModuleDao, err *e.ApiError) {
	return nil, e.NewDaoUnimplementedError()
}

// Returns module identified by id and version
func (m *ModuleDao) Get(id, version uint) (entity *models.Module, err *e.ApiError) {
	return nil, e.NewDaoUnimplementedError()
}

// Returns module by id with highest version
func (m *ModuleDao) GetLatest(id uint) (entity *models.Module, err *e.ApiError) {
	return nil, e.NewDaoUnimplementedError()
}

// Will create a new module if neither id nor version are set
// Will create a new module version if only id is set
func (m *ModuleDao) Create(entity models.Module) (returnEntity *models.Module, err *e.ApiError) {
	return nil, e.NewDaoUnimplementedError()
}

// Will update an existing module specified by id and version
func (m *ModuleDao) Update(entity models.Module) e.DaoError {
	return e.NewDaoUnimplementedError()
}

// Deletes a single module with id and version
func (m *ModuleDao) Delete(id, version uint) e.DaoError {
	return e.NewDaoUnimplementedError()
}

type CourseDao struct {
	repo *repository.CourseRepository
}

// Create new dao with required repositories
func NewCourseDao(courseRepository *repository.CourseRepository) *CourseDao {
	return &CourseDao{
		repo: courseRepository,
	}
}

// Returns course by id and version
func (c *CourseDao) GetAll(id, version uint) (entities []models.Course, err *e.ApiError) {
	courses, internalErr := c.repo.GetAll()
	if internalErr != nil {
		return nil, e.NewDaoDbError()
	}
	return courses, nil
}

// Returns course by id and version
func (c *CourseDao) Get(id, version uint) (entity *models.Course, err *e.ApiError) {
	return nil, e.NewDaoUnimplementedError()
}

// Returns course by id with highest version
func (c *CourseDao) GetLatest(id uint) (entity *models.Course, err *e.ApiError) {
	return nil, e.NewDaoUnimplementedError()
}

// Will create a new course if neither id nor version are set
// Will create a new course version if only id is set
func (c *CourseDao) Create(entity models.Course) (returnEntity *models.Course, err *e.ApiError) {
	return nil, e.NewDaoUnimplementedError()
}

// updates an existing course via id and version
func (c *CourseDao) Update(entity models.Course) *e.ApiError {
	return e.NewDaoUnimplementedError()
}

// Deletes a course by id and version
func (c *CourseDao) Delete(id, version uint) *e.ApiError {
	return e.NewDaoUnimplementedError()
}

type ExamDao struct{}

// Create new exam dao with all used providers
func NewDoaExam(examProvider *interface{}, courseProvider *interface{}) (daoExam *ExamDao, err *e.ApiError) {
	return nil, e.NewDaoUnimplementedError()
}

// Returns a list of exams for a specific course
func (ex *ExamDao) GetForCourse(courseId, courseVersion uint) (entities []models.Exam, err *e.ApiError) {
	return nil, e.NewDaoUnimplementedError()
}

// Returns a single exam for a course selected by id
func (ex *ExamDao) Get(courseId, courseVersion, examId uint) (entity *models.Exam, err *e.ApiError) {
	return nil, e.NewDaoUnimplementedError()
}

// Creates a new exam for a course
func (ex *ExamDao) Create(entity models.Exam) (returnEntity *models.Exam, err *e.ApiError) {
	return nil, e.NewDaoUnimplementedError()
}

// Updates an existing exam for a course
func (ex *ExamDao) Update(entity models.Exam) *e.ApiError {
	return e.NewDaoUnimplementedError()
}

// Deletes exam for a course
func (ex *ExamDao) Delete(courseId, courseVersion, examId uint) *e.ApiError {
	return e.NewDaoUnimplementedError()
}

type UserDao struct{}

// Creates new dao from required repositories
func NewUserDao(userRepository *interface{}) (dao *UserDao, err *e.ApiError) {
	return nil, e.NewDaoUnimplementedError()
}

// Returns a list of courses a user has assigned
func (u *UserDao) GetCourses(uid uint) (courses []models.SelectedCourse, err *e.ApiError) {
	return nil, e.NewDaoUnimplementedError()
}

// Returns a list of courses a user has assigned in a specific start year
func (u *UserDao) GetCoursesForYear(uid uint, startYear time.Time) (courses []models.SelectedCourse, err *e.ApiError) {
	return nil, e.NewDaoUnimplementedError()
}

func (u *UserDao) GetExamEvaluations()
