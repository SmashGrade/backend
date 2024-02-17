package dao

import (
	"time"

	e "github.com/SmashGrade/backend/app/error"
	"github.com/SmashGrade/backend/app/models"
	"github.com/SmashGrade/backend/app/repository"
)

// Asserts that an []any slice is a specific model slice via type assertion
// out := assertSlice[models.Course](courses)
func assertSlice[outputModel any](inputSlice []any) (outputSlice []outputModel) {
	outSlice := make([]outputModel, len(inputSlice))
	for i := range inputSlice {
		outSlice[i] = inputSlice[i].(outputModel)
	}

	return outSlice
}

// Returns specific outputModel slice from a repository getAll call
func getAllOrError[outputModel any](repo repository.Repository) (outputSlice []outputModel, err *e.ApiError) {
	var entity outputModel
	internalSlice, internalErr := repo.GetAll(entity)
	if internalErr != nil {
		err = e.NewDaoDbError()
		return
	}

	outputSlice = internalSlice.([]outputModel)
	return
}

// Returns sepcific outputModel entity reference from repository getId
func getOrError[outputModel any](repo repository.IdRepository, id uint) (outputEntity *outputModel, err *e.ApiError) {
	var entity outputModel
	ent, internalError := repo.GetId(id, entity)
	if internalError != nil {
		return nil, e.NewDaoDbError()
	}
	return ent.(*outputModel), nil
}

// Returns specific outputModel entity reference from repository getVersioned call
func getVersionedOrError[outputModel any](repo repository.VersionedRepository, id, version uint) (outputEntity *outputModel, err *e.ApiError) {
	var entity outputModel
	ent, internalError := repo.GetVersioned(id, version, entity)
	if internalError != nil {
		return nil, e.NewDaoDbError()
	}
	return ent.(*outputModel), nil
}

// Returns specific outputModel entity reference from repository getLatestVersioned call
func getLatestVersionedOrError[outputModel any](repo repository.VersionedRepository, id uint) (outputEntity *outputModel, err *e.ApiError) {
	var entity outputModel
	ent, internalError := repo.GetLatestVersioned(id, entity)
	if internalError != nil {
		return nil, e.NewDaoDbError()
	}
	return ent.(*outputModel), nil
}

// curriculum type / Studiengang art
// has description like Vollzeit or Berufsbegleitend
type CurriculumTypeDao struct {
	repo *repository.CurriculumtypeRepository
}

// Creates new dao with required repositories
func NewCurriculumTypeDao(curriculumTypeRepository *repository.CurriculumtypeRepository) *CurriculumTypeDao {
	return &CurriculumTypeDao{
		repo: curriculumTypeRepository,
	}
}

// Returns all curriculum types as slice
func (c *CurriculumTypeDao) GetAll() (entities []models.Curriculumtype, err *e.ApiError) {
	return getAllOrError[models.Curriculumtype](c.repo)
}

func (c *CurriculumTypeDao) Get(id uint) (entity *models.Curriculumtype, err *e.ApiError) {
	return getOrError[models.Curriculumtype](c.repo, id)
}

// gradetype / benotungssystem
// has a description
type GradeTypeDao struct {
	repo *repository.GradetypeRepository
}

// Creates new dao with required repositories
func NewGradeTypeDao(gradetypeRepository *repository.GradetypeRepository) *GradeTypeDao {
	return &GradeTypeDao{
		repo: gradetypeRepository,
	}
}

// Returns all grade types as slice
func (c *GradeTypeDao) GetAll() (entities []models.Gradetype, err *e.ApiError) {
	return getAllOrError[models.Gradetype](c.repo)
}

func (c *GradeTypeDao) Get(id uint) (entity *models.Gradetype, err *e.ApiError) {
	return getOrError[models.Gradetype](c.repo, id)
}

// evaluation type / Bewertungstyp for Module
type EvaluationTypeDao struct {
	repo *repository.EvaluationtypeRepository
}

// Creates new dao with required repositories
func NewEvaluationTypeDao(evaluationTypeRepository *repository.EvaluationtypeRepository) *EvaluationTypeDao {
	return &EvaluationTypeDao{
		repo: evaluationTypeRepository,
	}
}

// returns slice of all evaluation types
func (et *EvaluationTypeDao) GetAll() (entities []models.Evaluationtype, err *e.ApiError) {
	return getAllOrError[models.Evaluationtype](et.repo)
}

func (et *EvaluationTypeDao) Get(id uint) (entity *models.Evaluationtype, err *e.ApiError) {
	return getOrError[models.Evaluationtype](et.repo, id)
}

// state / Zustand for Module and Curriculum
type StateDao struct {
	repo *repository.StateRepository
}

// Creates new dao with required repositories
func NewStateDao(stateRepository *repository.StateRepository) *StateDao {
	return &StateDao{
		repo: stateRepository,
	}
}

// returns all states as slice
func (st *StateDao) GetAll() (entities []models.State, err *e.ApiError) {
	return getAllOrError[models.State](st.repo)
}

func (st *StateDao) Get(id uint) (entity *models.State, err *e.ApiError) {
	return getOrError[models.State](st.repo, id)
}

// Curriculum / Studiengang
// Highest level of categorization
type CurriculumDao struct {
	repo *repository.CurriculumRepository
}

// Create new curriculum with required repository
func NewCurriculumDao(curriculumRepository *repository.CurriculumRepository) *CurriculumDao {
	return &CurriculumDao{
		repo: curriculumRepository,
	}
}

// Returns existing curriculum
func (c *CurriculumDao) Get(id uint, startValidity time.Time) (entity *models.Curriculum, err *e.ApiError) {
	return nil, e.NewDaoUnimplementedError()
}

// Creates new curriculum
func (c *CurriculumDao) Create(entity *models.Curriculum) *e.ApiError {
	return e.NewDaoUnimplementedError()
}

// Module / Modul
// A collection of multiple courses
type ModuleDao struct {
	repo *repository.ModuleRepository
}

// Creates a new dao with required repositories
func NewModuleDao(moduleRepository *repository.ModuleRepository) *ModuleDao {
	return &ModuleDao{
		repo: moduleRepository,
	}
}

// returns all modules as slice
func (m *ModuleDao) GetAll() (entities []models.Module, err *e.ApiError) {
	return getAllOrError[models.Module](m.repo)
}

// Returns module identified by id and version
func (m *ModuleDao) Get(id, version uint) (entity *models.Module, err *e.ApiError) {
	return getVersionedOrError[models.Module](m.repo, id, version)
}

// Returns module by id with highest version
func (m *ModuleDao) GetLatest(id uint) (entity *models.Module, err *e.ApiError) {
	return getLatestVersionedOrError[models.Module](m.repo, id)
}

// Will create a new module if neither id nor version are set
// Will create a new module version if only id is set
func (m *ModuleDao) Create(entity models.Module) (returnEntity *models.Module, err *e.ApiError) {
	return nil, e.NewDaoUnimplementedError()
}

// Will update an existing module specified by id and version
func (m *ModuleDao) Update(entity models.Module) *e.ApiError {
	return e.NewDaoUnimplementedError()
}

// Deletes a single module with id and version
func (m *ModuleDao) Delete(id, version uint) *e.ApiError {
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
func (c *CourseDao) GetAll() (entities []models.Course, err *e.ApiError) {
	return getAllOrError[models.Course](c.repo)
}

// Returns course by id and version
func (c *CourseDao) Get(id, version uint) (entity *models.Course, err *e.ApiError) {
	return getVersionedOrError[models.Course](c.repo, id, version)
}

// Returns course by id with highest version
func (c *CourseDao) GetLatest(id uint) (entity *models.Course, err *e.ApiError) {
	return getLatestVersionedOrError[models.Course](c.repo, id)
}

// Will create a new course if neither id nor version are set
// Will create a new course version if only id is set
func (c *CourseDao) Create(entity *models.Course) (returnEntity *models.Course, err *e.ApiError) {

	internalEntity, internalError := c.repo.Create(entity)

	if internalError != nil {
		return nil, e.NewDaoDbError()
	}

	return internalEntity.(*models.Course), nil
}

// updates an existing course via id and version
func (c *CourseDao) Update(entity models.Course) *e.ApiError {
	internalError := c.repo.Update(entity)
	if internalError != nil {
		return e.NewDaoDbError()
	}

	return nil
}

// Deletes a course by id and version
func (c *CourseDao) Delete(id, version uint) *e.ApiError {
	var entity models.Course
	internalError := c.repo.DeleteVersioned(id, version, entity)
	if internalError != nil {
		return e.NewDaoDbError()
	}

	return nil
}

type ExamDao struct{}

// Create new exam dao with all used providers
func NewDoaExam(examProvider *interface{}, courseProvider *interface{}) *ExamDao {
	return nil
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
func NewUserDao(userRepository *interface{}) *UserDao {
	return nil
}

// Returns a list of courses a user has assigned
func (u *UserDao) GetCourses(uid uint) (courses []models.SelectedCourse, err *e.ApiError) {
	return nil, e.NewDaoUnimplementedError()
}

// Returns a list of courses a user has assigned in a specific start year
func (u *UserDao) GetCoursesForYear(uid uint, startYear time.Time) (courses []models.SelectedCourse, err *e.ApiError) {
	return nil, e.NewDaoUnimplementedError()
}

// Returns a list of evaluations by startYear and exam
func (u *UserDao) GetExamEvaluationsForYear(uid uint, startYear time.Time) {
	// TODO
}
