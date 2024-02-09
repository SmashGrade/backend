package dao

import (
	"time"

	e "github.com/SmashGrade/backend/app/error"
	"github.com/SmashGrade/backend/app/models"
)

// curriculum type / Studiengang art
// has description like Vollzeit or Berufsbegleitend
type DaoCurriculumType struct {
}

func (c *DaoCurriculumType) GetAll() (entities []models.Curriculumtype, err e.DaoError) {
	return nil, e.DAOUnimplemented
}

// Curriculum / Studiengang
// Highest level of categorization
type DaoCurriculum struct{}

// Returns existing curriculum
func (c *DaoCurriculum) Get(id uint, startValidity time.Time) (entity *models.Curriculum, err e.DaoError) {
	return nil, e.DAOUnimplemented
}

// Creates new curriculum
func (c *DaoCurriculum) Create(entity *models.Curriculum) e.DaoError {
	return e.DAOUnimplemented
}

// Module / Modul
// A collection of multiple courses
type DaoModule struct{}

// Returns module identified by id and version
func (m *DaoModule) Get(id, version uint) (entity *models.Module, err e.DaoError) {
	return nil, e.DAOUnimplemented
}

// Returns module by id with highest version
func (m *DaoModule) GetLatest(id uint) (entity *models.Module, err e.DaoError) {
	return nil, e.DAOUnimplemented
}

// Will create a new module if neither id nor version are set
// Will create a new module version if only id is set
func (m *DaoModule) Create(entity models.Module) (returnEntity *models.Module, err e.DaoError) {
	return nil, e.DAOUnimplemented
}

// Will update an existing module specified by id and version
func (m *DaoModule) Update(entity models.Module) e.DaoError {
	return e.DAOUnimplemented
}

// Deletes a single module with id and version
func (m *DaoModule) Delete(id, version uint) e.DaoError {
	return e.DAOUnimplemented
}

type DaoCourse struct{}

// Returns course by id and version
func (c *DaoCourse) Get(id, version uint) (entity *models.Course, err e.DaoError) {
	return nil, e.DAOUnimplemented
}

// Returns course by id with highest version
func (c *DaoCourse) GetLatest(id uint) (entity *models.Course, err e.DaoError) {
	return nil, e.DAOUnimplemented
}

// Will create a new course if neither id nor version are set
// Will create a new course version if only id is set
func (c *DaoCourse) Create(entity models.Course) (returnEntity *models.Course, err e.DaoError) {
	return nil, e.DAOUnimplemented
}

// updates an existing course via id and version
func (c *DaoCourse) Update(entity models.Course) e.DaoError {
	return e.DAOUnimplemented
}

// Deletes a course by id and version
func (c *DaoCourse) Delete(id, version uint) e.DaoError {
	return e.DAOUnimplemented
}

type DaoExam struct{}

// Returns a list of exams for a specific course
func (ex *DaoExam) GetForCourse(courseId, courseVersion uint) (entities []models.Exam, err e.DaoError) {
	return nil, e.DAOUnimplemented
}

// Returns a single exam for a course selected by id
func (ex *DaoExam) Get(courseId, courseVersion, examId uint) (entity *models.Exam, err e.DaoError) {
	return nil, e.DAOUnimplemented
}

// Creates a new exam for a course
func (ex *DaoExam) Create(entity models.Exam) (returnEntity *models.Exam, err e.DaoError) {
	return nil, e.DAOUnimplemented
}

// Updates an existing exam for a course
func (ex *DaoExam) Update(entity models.Exam) e.DaoError {
	return e.DAOUnimplemented
}

// Deletes exam for a course
func (ex *DaoExam) Delete(courseId, courseVersion, examId uint) e.DaoError {
	return e.DAOUnimplemented
}

type DaoUser struct{}

// Returns a list of courses a
func (u *DaoUser) GetCourses(id uint, startYear time.Time) (courses []models.SelectedCourse, err e.DaoError) {
	return nil, e.DAOUnimplemented
}
