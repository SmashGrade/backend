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

}

// Curriculum / Studiengang
// Highest level of categorization
type DaoCurriculum struct{}

// Returns existing curriculum
func (c *DaoCurriculum) Get(id uint, startValidity time.Time) (entity models.Curriculum, err e.DaoError) {

}

// Creates new curriculum
func (c *DaoCurriculum) Create(entity *models.Curriculum) e.DaoError {

}

// Module / Modul
// A collection of multiple courses
type DaoModule struct{}

// Returns module identified by id and version
func (m *DaoModule) Get(id, version uint) (entity models.Module, err e.DaoError) {

}

// Returns module by id with highest version
func (m *DaoModule) GetLatest(id uint) (entity models.Module, err e.DaoError) {

}

// Will create a new module if neither id nor version are set
// Will create a new module version if only id is set
func (m *DaoModule) Create(entity models.Module) (returnEntity models.Module, err e.DaoError) {

}

// Will update an existing module specified by id and version
func (m *DaoModule) Update(entity models.Module) e.DaoError {

}

// Deletes a single module with id and version
func (m *DaoModule) Delete(id, version uint) {

}
