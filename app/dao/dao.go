package dao

import (
	"fmt"
	"time"

	"github.com/SmashGrade/backend/app/config"
	e "github.com/SmashGrade/backend/app/error"
	"github.com/SmashGrade/backend/app/models"
	"github.com/SmashGrade/backend/app/repository"
	"github.com/SmashGrade/backend/app/requestmodels"
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
	internalSlice, internalErr := repo.GetAll()
	if internalErr != nil {
		err = e.NewDaoDbError()
		return
	}

	outputSlice = internalSlice.([]outputModel)
	return
}

// Returns sepcific outputModel entity reference from repository getId
func getOrError[outputModel any](repo repository.IdRepository, id uint) (outputEntity *outputModel, err *e.ApiError) {
	ent, internalError := repo.GetId(id)
	if internalError != nil {
		return nil, e.NewDaoDbError()
	}
	return ent.(*outputModel), nil
}

// Returns specific outputModel entity reference from repository getVersioned call
func getVersionedOrError[outputModel any](repo repository.VersionedRepository, id, version uint) (outputEntity *outputModel, err *e.ApiError) {
	ent, internalError := repo.GetVersioned(id, version)
	if internalError != nil {
		return nil, e.NewDaoDbError()
	}
	return ent.(*outputModel), nil
}

// Returns specific outputModel entity refernece from repository getTimed call
func getTimedOrError[outputModul any](repo repository.TimedRepository, id uint, date time.Time) (outputEntity *outputModul, err *e.ApiError) {
	ent, internalError := repo.GetTimed(id, date)
	if internalError != nil {
		return nil, e.NewDaoDbError()
	}
	return ent.(*outputModul), nil
}

// Returns specific outputModel entity reference from repository getLatestVersioned call
func getLatestVersionedOrError[outputModel any](repo repository.VersionedRepository, id uint) (outputEntity *outputModel, err *e.ApiError) {
	ent, internalError := repo.GetLatestVersioned(id)
	if internalError != nil {
		return nil, e.NewDaoDbError()
	}
	return ent.(*outputModel), nil
}

// Returns a specific modelType entity reference from generic repository create call
func createOrError[modelType any](repo repository.Repository, entity modelType) (returnEntity *modelType, err *e.ApiError) {
	internalEntity, internalError := repo.Create(&entity)

	if internalError != nil {
		return nil, e.NewDaoDbError()
	}

	return internalEntity.(*modelType), nil
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

func (c *CurriculumTypeDao) Create(entity models.Curriculumtype) (returnEntity *models.Curriculumtype, err *e.ApiError) {
	return createOrError(c.repo, entity)
}

// Create default values for curriculum type
func (c *CurriculumTypeDao) CreateDefaults() *e.ApiError {
	existingEntities, err := c.GetAll()
	if err != nil {
		return err
	}

	for _, v := range c.repo.Provider.Config().CurriculumTypes {

		existingFound := false
		for _, existing := range existingEntities {
			if v.Description == existing.Description {
				existingFound = true
				break
			}
		}
		if existingFound {
			continue
		}

		_, err := c.Create(models.Curriculumtype{
			Description:   v.Description,
			DurationYears: v.DurationYears,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

// Parse string to Time. Give the Layout for example
//
//	"2006-01-02 15:04:05 MST"
//	or "02.01.2006"
func ParseTime(timestring, layout string) (time.Time, error) {
	// Parse string to time.Time
	dateTime, err := time.Parse(layout, timestring)
	if err != nil {
		return time.Time{}, err
	}
	return dateTime, nil
}

// Update Curriculum with id and date
func (c *CurriculumDao) Update(referenceEntity requestmodels.RefCurriculum) *e.ApiError {
	dateTime, err := ParseTime(referenceEntity.StartValidity, "02.01.2006")
	if err != nil {
		return e.NewDaoReferenceError("date", referenceEntity.StartValidity)
	}
	existingEntity, getErr := c.Get(referenceEntity.ID, dateTime)
	if getErr != nil {
		return e.NewDaoNotExistingError("curriculum", fmt.Sprintf("id: %v, date: %v", referenceEntity.ID, referenceEntity.StartValidity))
	}

	entity, convertErr := c.convertRefCurriculumToCurriculum(referenceEntity)
	if convertErr != nil {
		return e.NewApiUnimplementedError()
	}

	// add existing id and date
	entity.ID = existingEntity.ID
	entity.StartValidity = existingEntity.StartValidity

	internalError := c.repo.Update(&entity)
	if internalError != nil {
		return e.NewDaoDbError()
	}

	return nil
}

// Delete Curriculum with id and date
func (c *CurriculumDao) Delete(id uint, date time.Time) *e.ApiError {
	internalError := c.repo.DeleteTimed(id, date)
	if internalError != nil {
		return e.NewDaoDbError()
	}
	return nil
}

// field / Schwerpunkt
type FieldDao struct {
	repo *repository.FieldRepository
}

// Creates new dao with required repositories
func NewFieldDao(fieldRepository *repository.FieldRepository) *FieldDao {
	return &FieldDao{
		repo: fieldRepository,
	}
}

// Returns all fields as slice
func (c *FieldDao) GetAll() (entities []models.Field, err *e.ApiError) {
	return getAllOrError[models.Field](c.repo)
}

func (c *FieldDao) Get(id uint) (entity *models.Field, err *e.ApiError) {
	return getOrError[models.Field](c.repo, id)
}

func (c *FieldDao) Create(entity models.Field) (returnEntity *models.Field, err *e.ApiError) {
	return createOrError(c.repo, entity)
}

// focus / Fachrichtung
type FocusDao struct {
	repo *repository.FocusRepository
}

// Creates new dao with required repositories
func NewFocusDao(focusRepository *repository.FocusRepository) *FocusDao {
	return &FocusDao{
		repo: focusRepository,
	}
}

// Returns all focus as slice
func (c *FocusDao) GetAll() (entities []models.Focus, err *e.ApiError) {
	return getAllOrError[models.Focus](c.repo)
}

func (c *FocusDao) Get(id uint) (entity *models.Focus, err *e.ApiError) {
	return getOrError[models.Focus](c.repo, id)
}

func (c *FocusDao) Create(entity models.Focus) (returnEntity *models.Focus, err *e.ApiError) {
	return createOrError(c.repo, entity)
}

// exam type / Test art
type ExamtypeDao struct {
	repo *repository.ExamtypeRepository
}

// Creates new dao with required repositories
func NewExamtypeDao(examtypeRepository *repository.ExamtypeRepository) *ExamtypeDao {
	return &ExamtypeDao{
		repo: examtypeRepository,
	}
}

// Returns all exam types as slice
func (c *ExamtypeDao) GetAll() (entities []models.Examtype, err *e.ApiError) {
	return getAllOrError[models.Examtype](c.repo)
}

func (c *ExamtypeDao) Get(id uint) (entity *models.Examtype, err *e.ApiError) {
	return getOrError[models.Examtype](c.repo, id)
}

func (c *ExamtypeDao) Create(entity models.Examtype) (returnEntity *models.Examtype, err *e.ApiError) {
	internalEntity, internalError := c.repo.Create(&entity)

	if internalError != nil {
		return nil, e.NewDaoDbError()
	}

	return internalEntity.(*models.Examtype), nil
}

// Create default values for exam type
func (c *ExamtypeDao) CreateDefaults() *e.ApiError {
	existingEntities, err := c.GetAll()
	if err != nil {
		return err
	}

	for _, v := range c.repo.Provider.Config().ExamTypes {

		existingFound := false
		for _, existing := range existingEntities {
			if v == existing.Description {
				existingFound = true
				break
			}
		}
		if existingFound {
			continue
		}

		_, err := c.Create(models.Examtype{
			Description: v,
		})
		if err != nil {
			return err
		}
	}

	return nil
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

func (c *GradeTypeDao) Create(entity models.Gradetype) (returnEntity *models.Gradetype, err *e.ApiError) {
	internalEntity, internalError := c.repo.Create(&entity)

	if internalError != nil {
		return nil, e.NewDaoDbError()
	}

	return internalEntity.(*models.Gradetype), nil
}

// Create default values for grade type
func (c *GradeTypeDao) CreateDefaults() *e.ApiError {
	existingEntities, err := c.GetAll()
	if err != nil {
		return err
	}

	for _, v := range c.repo.Provider.Config().GradeTypes {

		existingFound := false
		for _, existing := range existingEntities {
			if v.Description == existing.Description {
				existingFound = true
				break
			}
		}
		if existingFound {
			continue
		}

		_, err := c.Create(models.Gradetype{
			Description: v.Description,
		})
		if err != nil {
			return err
		}
	}

	return nil
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

func (et *EvaluationTypeDao) Create(entity models.Evaluationtype) (returnEntity *models.Evaluationtype, err *e.ApiError) {
	internalEntity, internalError := et.repo.Create(&entity)

	if internalError != nil {
		return nil, e.NewDaoDbError()
	}

	return internalEntity.(*models.Evaluationtype), nil
}

// Create default values for grade type
func (et *EvaluationTypeDao) CreateDefaults() *e.ApiError {
	existingEntities, err := et.GetAll()
	if err != nil {
		return err
	}

	for _, v := range et.repo.Provider.Config().ExamEvaluationTypes {

		existingFound := false
		for _, existing := range existingEntities {
			if v.Description == existing.Description {
				existingFound = true
				break
			}
		}
		if existingFound {
			continue
		}

		_, err := et.Create(models.Evaluationtype{
			Description: v.Description,
			Code:        v.Code,
		})
		if err != nil {
			return err
		}
	}

	return nil
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

func (st *StateDao) Create(entity models.State) (returnEntity *models.State, err *e.ApiError) {
	internalEntity, internalError := st.repo.Create(&entity)

	if internalError != nil {
		return nil, e.NewDaoDbError()
	}

	return internalEntity.(*models.State), nil
}

// Create default values for state
func (st *StateDao) CreateDefaults() *e.ApiError {
	existingEntities, err := st.GetAll()
	if err != nil {
		return err
	}

	for _, v := range st.repo.Provider.Config().States {

		existingFound := false
		for _, existing := range existingEntities {
			if v == existing.Description {
				existingFound = true
				break
			}
		}
		if existingFound {
			continue
		}

		_, err := st.Create(models.State{
			Description: v,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

// Curriculum / Studiengang
// Highest level of categorization
type CurriculumDao struct {
	repo              *repository.CurriculumRepository
	focusDao          *FocusDao
	curriculumTypeDao *CurriculumTypeDao
	stateDao          *StateDao
	moduleDao         *ModuleDao
}

// Create new curriculum with required repository
func NewCurriculumDao(curriculumRepository *repository.CurriculumRepository,
	focusRepository *repository.FocusRepository,
	curriculumTypeRepository *repository.CurriculumtypeRepository,
	stateRepository *repository.StateRepository,
	moduleRepository *repository.ModuleRepository) *CurriculumDao {
	return &CurriculumDao{
		repo:              curriculumRepository,
		focusDao:          NewFocusDao(focusRepository),
		curriculumTypeDao: NewCurriculumTypeDao(curriculumTypeRepository),
		stateDao:          NewStateDao(stateRepository),
		moduleDao:         NewModuleDao(moduleRepository),
	}
}

// returns all modules as slice
func (c *CurriculumDao) GetAll() (entities []models.Curriculum, err *e.ApiError) {
	return getAllOrError[models.Curriculum](c.repo)
}

// Returns existing curriculum
func (c *CurriculumDao) Get(id uint, startValidity time.Time) (entity *models.Curriculum, err *e.ApiError) {
	return getTimedOrError[models.Curriculum](c.repo, id, startValidity)
}

// returns the active version of a curriculum at the time timePoint
// returns an error if no active curriculum at that time exists
func (c *CurriculumDao) GetValidForTimepoint(id uint, timePoint time.Time) (entity *models.Curriculum, err *e.ApiError) {
	curriculums, err := c.GetAll()
	if err != nil {
		return nil, err
	}

	for i := range curriculums {
		if curriculums[i].StartValidity.After(timePoint) {
			continue
		}

		if !curriculums[i].EndValidity.IsZero() && curriculums[i].EndValidity.Before(timePoint) {
			continue
		}

		return &curriculums[i], nil
	}

	return nil, e.NewDaoNotExistingError("curriculum", timePoint.String())
}

// Creates new curriculum
func (c *CurriculumDao) Create(referenceEntity requestmodels.RefCurriculum) (returnEntity *models.Curriculum, err *e.ApiError) {
	var internalError error

	entity, err := c.convertRefCurriculumToCurriculum(referenceEntity)
	if err != nil {
		return nil, err
	}

	// Check if id == 0,
	if entity.ID == 0 {
		entity.ID, internalError = c.repo.GetNextId()
		if internalError != nil {
			return nil, e.NewDaoDbError()
		}
	}

	internalEntity, internalError := c.repo.Create(&entity)
	if internalError != nil {
		return nil, e.NewDaoDbError()
	}

	return internalEntity.(*models.Curriculum), nil
}

func (c *CurriculumDao) convertRefCurriculumToCurriculum(ent requestmodels.RefCurriculum) (retEnt models.Curriculum, err *e.ApiError) {
	var endValidity time.Time
	var parseErr error
	if ent.EndValidity == "" {
		// the current curriculum could have a end validity not set this means it is valid until the end of time
		endValidity = time.Time{}
	} else {
		endValidity, parseErr = ParseTime(ent.EndValidity, "02.01.2006")
		if parseErr != nil {
			err = e.NewDaoValidationError("EndValidity", "date in format dd.MM.yyyy", ent.EndValidity)
			return
		}
	}
	startValidity, parseErr := ParseTime(ent.StartValidity, "02.01.2006")
	if parseErr != nil {
		err = e.NewDaoValidationError("StartValidity", "date in format dd.MM.yyyy", ent.StartValidity)
		return
	}

	retEnt = models.Curriculum{
		EndValidity: endValidity,
		Description: ent.Description,
	}
	retEnt.StartValidity = startValidity
	retEnt.ID = ent.ID

	// Get linked Focus
	if ent.FocusID.ID != 0 {
		resFocus, internalError := c.focusDao.Get(ent.FocusID.ID)
		if internalError != nil {
			err = e.NewDaoReferenceIdError("focus", ent.ID)
			return
		}
		retEnt.Focus = *resFocus
	}

	// Get linked Curriculumtype
	if ent.CurriculumtypeID.ID != 0 {
		resCurriculumType, internalError := c.curriculumTypeDao.Get(ent.CurriculumtypeID.ID)
		if internalError != nil {
			err = e.NewDaoReferenceIdError("curriculumType", ent.CurriculumtypeID.ID)
			return
		}
		retEnt.Curriculumtype = *resCurriculumType
	}

	// Get linked State
	if ent.StateID.ID != 0 {
		resStateType, internalError := c.stateDao.Get(ent.StateID.ID)
		if internalError != nil {
			err = e.NewDaoReferenceIdError("state", ent.StateID.ID)
			return
		}
		retEnt.State = *resStateType
	}

	// Get linked Module list
	var resolvedModuleList []*models.Module
	for _, module := range ent.Modules {
		if module.ID != 0 {
			resModule, internalError := c.moduleDao.Get(module.ID, module.Version)
			if internalError != nil {
				err = e.NewDaoReferenceVersionedError("module", module.ID, module.Version)
			}
			resolvedModuleList = append(resolvedModuleList, resModule)
		}
	}
	retEnt.Modules = resolvedModuleList

	return
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
	var internalError error

	if entity.ID == 0 {
		entity.ID, internalError = m.repo.GetNextId()
		if internalError != nil {
			return nil, e.NewDaoDbError()
		}
		entity.Version = 1 // set version to one on new id
	} else {
		if entity.Version == 0 { // generate new version if it is initial on existing id
			entity.Version, internalError = m.repo.GetNextVersion(entity.ID)
			if internalError != nil {
				return nil, e.NewDaoDbError()
			}
		}
	}

	internalEntity, internalError := m.repo.Create(&entity)

	if internalError != nil {
		return nil, e.NewDaoDbError()
	}

	return internalEntity.(*models.Module), nil
}

// Will update an existing module specified by id and version
func (m *ModuleDao) Update(entity models.Module) *e.ApiError {
	internalError := m.repo.Update(entity)
	if internalError != nil {
		return e.NewDaoDbError()
	}

	return nil
}

// Deletes a single module with id and version
func (m *ModuleDao) Delete(id, version uint) *e.ApiError {
	internalError := m.repo.DeleteVersioned(id, version)
	if internalError != nil {
		return e.NewDaoDbError()
	}

	return nil
}

type SelectedCourseDao struct {
	repo *repository.SelectedCourseRepository
}

// Create new selected course dao
func NewSelectedCourseDao(selectedCourseRepository *repository.SelectedCourseRepository) *SelectedCourseDao {
	return &SelectedCourseDao{
		repo: selectedCourseRepository,
	}
}

// Returns existing selected course
func (c *SelectedCourseDao) Get(userId, courseId, courseVersion uint, classStartYear time.Time) (entity *models.SelectedCourse, err *e.ApiError) {
	internalEntity, internalError := c.repo.GetSelectedCourse(userId, courseId, courseVersion, classStartYear)
	if internalError != nil {
		return nil, e.NewDaoDbError()
	}
	return &internalEntity, nil
}

type CourseDao struct {
	repo              *repository.CourseRepository
	moduleDao         *ModuleDao
	userDao           *UserDao
	selectedCourseDao *SelectedCourseDao
	examDao           *ExamDao
}

// Create new dao with required repositories
func NewCourseDao(courseRepository *repository.CourseRepository, moduleRepository *repository.ModuleRepository, userRepository *repository.UserRepository, selectedCourseRepository *repository.SelectedCourseRepository, examRepository *repository.ExamRepository, roleRepository *repository.RoleRepository) *CourseDao {
	return &CourseDao{
		repo:              courseRepository,
		moduleDao:         NewModuleDao(moduleRepository),
		userDao:           NewUserDao(userRepository, roleRepository),
		selectedCourseDao: NewSelectedCourseDao(selectedCourseRepository),
		examDao:           NewExamDao(examRepository, courseRepository),
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

func (c *CourseDao) convertRefCourseToCourse(ent requestmodels.RefCourse) (retEnt models.Course, err *e.ApiError) {
	retEnt = models.Course{
		Description: ent.Description,
		Number:      ent.Number,
	}
	retEnt.ID = ent.ID
	retEnt.Version = ent.Version

	// get linked modules
	var resolvedModuleList []*models.Module
	for _, mod := range ent.Modules {
		resMod, internalError := c.moduleDao.Get(mod.ID, mod.Version)
		if internalError != nil {
			err = e.NewDaoReferenceVersionedError("module", mod.ID, mod.Version)
			return
		}
		resolvedModuleList = append(resolvedModuleList, resMod)
	}
	retEnt.Modules = resolvedModuleList

	// get linked teachers
	var resolvedTecherList []*models.User
	for _, teacher := range ent.TeachedBy {
		resTeacher, internalError := c.userDao.Get(teacher.ID)
		if internalError != nil {
			err = e.NewDaoReferenceIdError("teachedBy User", teacher.ID)
			return
		}
		resolvedTecherList = append(resolvedTecherList, resTeacher)
	}
	retEnt.TeachedBy = resolvedTecherList

	// get linked selected courses
	var resolvedSelectedCoursesList []models.SelectedCourse
	for _, selCourse := range ent.SelectedCourses {
		resSelCourse, internalError := c.selectedCourseDao.Get(selCourse.UserID, selCourse.CourseID, selCourse.CourseVersion, selCourse.ClassStartyear)
		if internalError != nil {
			err = e.NewDaoReferenceError("selected course", fmt.Sprintf("user id: %v, course id: %v, course version: %v, class start year: %v", selCourse.UserID, selCourse.CourseID, selCourse.CourseVersion, selCourse.ClassStartyear))
			return
		}
		resolvedSelectedCoursesList = append(resolvedSelectedCoursesList, *resSelCourse)
	}
	retEnt.SelectedCourses = resolvedSelectedCoursesList

	// get linked exams
	var resolvedExamList []*models.Exam
	for _, exm := range ent.Exams {
		resExam, internalError := c.examDao.Get(exm.ID)
		if internalError != nil {
			err = e.NewDaoReferenceIdError("exam", exm.ID)
			return
		}
		resolvedExamList = append(resolvedExamList, resExam)
	}
	retEnt.Exams = resolvedExamList

	return
}

// Will create a new course if neither id nor version are set
// Will create a new course version if only id is set
func (c *CourseDao) Create(referenceEntity requestmodels.RefCourse) (returnEntity *models.Course, err *e.ApiError) {
	var internalError error

	// check if a selected course is set already. this should not be possible
	if len(referenceEntity.SelectedCourses) > 0 {
		return nil, e.NewDaoValidationError("selected courses", "empty", "filled")
	}

	entity, err := c.convertRefCourseToCourse(referenceEntity)
	if err != nil {
		return nil, err
	}

	// First check if the id is zero, if yes generate it
	if entity.ID == 0 {
		entity.ID, internalError = c.repo.GetNextId()
		if internalError != nil {
			return nil, e.NewDaoDbError()
		}
		entity.Version = 1 // set version to one on new id
	} else {
		if entity.Version == 0 { // generate new version if it is initial on existing id
			entity.Version, internalError = c.repo.GetNextVersion(entity.ID)
			if internalError != nil {
				return nil, e.NewDaoDbError()
			}
		}
	}

	internalEntity, internalError := c.repo.Create(&entity)

	if internalError != nil {
		return nil, e.NewDaoDbError()
	}

	return internalEntity.(*models.Course), nil
}

// updates an existing course via id and version
func (c *CourseDao) Update(referenceEntity requestmodels.RefCourse) *e.ApiError {
	existingEntity, err := c.Get(referenceEntity.ID, referenceEntity.Version)
	if err != nil {
		return e.NewDaoNotExistingError("course", fmt.Sprintf("id: %v, version: %v", referenceEntity.ID, referenceEntity.Version))
	}

	entity, err := c.convertRefCourseToCourse(referenceEntity)

	// add in existing id and version
	entity.ID = existingEntity.ID
	entity.Version = existingEntity.Version

	if err != nil {
		return err
	}

	internalError := c.repo.Update(entity)
	if internalError != nil {
		return e.NewDaoDbError()
	}

	return nil
}

// Deletes a course by id and version
func (c *CourseDao) Delete(id, version uint) *e.ApiError {
	internalError := c.repo.DeleteVersioned(id, version)
	if internalError != nil {
		return e.NewDaoDbError()
	}

	return nil
}

type ExamDao struct {
	examRepo   *repository.ExamRepository
	courseRepo *repository.CourseRepository
}

// Create new exam dao with all used providers
func NewExamDao(examRepository *repository.ExamRepository, courseRepository *repository.CourseRepository) *ExamDao {
	return &ExamDao{
		examRepo:   examRepository,
		courseRepo: courseRepository,
	}
}

// Returns a list of exams for a specific course
func (ex *ExamDao) GetForCourse(courseId, courseVersion uint) (entities []models.Exam, err *e.ApiError) {
	// courseEnt, err := ex.courseRepo.GetVersioned(courseId, courseVersion) // TODO: waiting for uuid fix
	return nil, e.NewDaoUnimplementedError()
}

// Returns a single exam by id
func (ex *ExamDao) Get(id uint) (entity *models.Exam, err *e.ApiError) {
	return getOrError[models.Exam](ex.examRepo, id)
}

// Returns all examn
func (ex *ExamDao) GetAll() (entities []models.Exam, err *e.ApiError) {
	return getAllOrError[models.Exam](ex.examRepo)
}

// Creates a new exam for a course
func (ex *ExamDao) Create(entity models.Exam) (returnEntity *models.Exam, err *e.ApiError) {
	return createOrError(ex.examRepo, entity)
}

// Updates an existing exam for a course
func (ex *ExamDao) Update(entity models.Exam) *e.ApiError {
	internalError := ex.examRepo.Update(entity)
	if internalError != nil {
		return e.NewDaoDbError()
	}

	return nil
}

// Deletes exam
func (ex *ExamDao) Delete(id uint) *e.ApiError {
	internalError := ex.examRepo.DeleteId(id)
	if internalError != nil {
		return e.NewDaoDbError()
	}

	return nil
}

type UserDao struct {
	repo     *repository.UserRepository
	roleRepo *repository.RoleRepository
}

// Creates new dao from required repositories
func NewUserDao(userRepository *repository.UserRepository, roleRepository *repository.RoleRepository) *UserDao {
	return &UserDao{
		repo:     userRepository,
		roleRepo: roleRepository,
	}
}

func (u *UserDao) Get(uid uint) (entity *models.User, err *e.ApiError) {
	return getOrError[models.User](u.repo, uid)
}

// generic by role filter used in other functions
func (u *UserDao) GetByRole(roleId uint) (entities []models.User, err *e.ApiError) {
	roleEnt, internalError := u.roleRepo.GetId(roleId)
	if internalError != nil {
		return nil, e.NewDaoDbError()
	}

	role := roleEnt.(*models.Role)
	users := make([]models.User, 0)
	for i := range role.Users {
		users = append(users, *role.Users[i])
	}
	return users, nil
}

// Returns all Teachers as User types as slice
func (u *UserDao) GetTeachers() (entities []models.User, err *e.ApiError) {
	return u.GetByRole(config.ROLE_TEACHER)
}

// Returns all Students as User types as slice
func (u *UserDao) GetStudents() (entities []models.User, err *e.ApiError) {
	return u.GetByRole(config.ROLE_STUDENT)
}

// Returns all CourseAdmins as User types as slice
func (u *UserDao) GetCourseAdmins() (entities []models.User, err *e.ApiError) {
	return u.GetByRole(config.ROLE_COURSEADMIN)
}

// Returns all FieldManagers as User types as slice
func (u *UserDao) GetFieldManagers() (entities []models.User, err *e.ApiError) {
	return u.GetByRole(config.ROLE_FIELDMANAGER)
}

// Returns a role by claim name
func (u *UserDao) GetRoleByClaim(claimName string) (entity *models.Role, err *e.ApiError) {
	entities, internalError := u.roleRepo.Find(&models.Role{Claim: claimName})
	if internalError != nil {
		return nil, e.NewDaoDbError()
	}

	roleEntities := entities.([]models.Role)
	if len(roleEntities) < 1 {
		return nil, e.NewDaoNotExistingError(claimName, "claim")
	}

	return &roleEntities[0], nil
}

// Create default values for roles
func (u *UserDao) CreateDefaults() *e.ApiError {
	existingEntities, err := u.roleRepo.GetAll()
	existingRoles := existingEntities.([]models.Role)
	if err != nil {
		return e.NewDaoDbError()
	}

	for _, v := range u.repo.Provider.Config().Roles {

		existingFound := false
		for _, existing := range existingRoles {
			if v.Id == existing.ID {
				existingFound = true

				// update description and claim of existing role
				updatedRole := &models.Role{
					Description: v.Name,
					Claim:       v.ClaimName,
					Users:       existing.Users,
				}
				updatedRole.ID = existing.ID
				err = u.roleRepo.Update(updatedRole)
				if err != nil {
					return e.NewDaoDbError()
				}

				break
			}
		}
		if existingFound {
			continue
		}

		// create new entry
		newRole := &models.Role{
			Description: v.Name,
			Claim:       v.ClaimName,
		}
		newRole.ID = v.Id

		_, err := u.roleRepo.Create(newRole)
		if err != nil {
			return e.NewDaoDbError()
		}
	}

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

// Creates a new user
func (u *UserDao) Create(entity models.User) (returnEntity *models.User, err *e.ApiError) {
	return createOrError(u.repo, entity)
}

// Updates an existing user
func (u *UserDao) Update(entity models.User) *e.ApiError {
	internalError := u.repo.Update(entity)
	if internalError != nil {
		return e.NewDaoDbError()
	}

	return nil
}

// This function is used to update or create a user in the database based on the given user object
// if a user with the supplied email exists: update data in db, return with id and references filled
// if no user with email found: create new and return with id and references filled
func (u *UserDao) CreateOrUpdateByEmail(entity models.User) (returnEntity *models.User, err *e.ApiError) {
	existingUser, err := u.GetByEmail(entity.Email)
	if err != nil { // TODO: close error inspection with errors.Is needed
		// user does not exist, create new return obj
		return u.Create(entity)
	} else {
		// user exists, update in db return updated version
		entity.ID = existingUser.ID
		err = u.Update(entity)
		if err != nil {
			return
		}

		return u.Get(entity.ID)
	}
}

// returns first match for an email
func (u *UserDao) GetByEmail(email string) (entity *models.User, err *e.ApiError) {
	entities, internalError := u.repo.Find(&models.User{Email: email})
	if internalError != nil {
		return nil, e.NewDaoDbError()
	}

	userEntities, assertionOk := entities.([]models.User)
	if !assertionOk {
		return nil, e.NewDaoDbError()
	}

	if len(userEntities) < 1 {
		return nil, e.NewDaoDbError()
	}

	return &userEntities[0], nil
}
