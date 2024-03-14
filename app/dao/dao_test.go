package dao

import (
	"fmt"
	"testing"
	"time"

	"github.com/SmashGrade/backend/app/db"
	_ "github.com/SmashGrade/backend/app/docs"
	"github.com/SmashGrade/backend/app/models"
	"github.com/SmashGrade/backend/app/repository"
	"github.com/SmashGrade/backend/app/requestmodels"
	"github.com/stretchr/testify/require"
	_ "gorm.io/gorm"
)

// Smoketest
func TestMagicSmoke(t *testing.T) {
	provider := db.NewMockProvider()

	repo := repository.NewCourseRepository(provider)

	dao := NewCourseDao(repo, repository.NewModuleRepository(provider), repository.NewUserRepository(provider), repository.NewSelectedCourseRepository(provider), repository.NewExamRepository(provider), repository.NewRoleRepository(provider), repository.NewExamtypeRepository(provider))

	courseEnt := requestmodels.RefCourse{Description: "Lol"}

	retEnt, err := dao.Create(courseEnt)
	if err != nil {
		t.Fatalf("Got db error")
	}

	t.Logf("Got '%v'", retEnt)
}

// GetAll should give a slice of ents
func TestGetAll(t *testing.T) {
	provider := db.NewMockProvider()

	repo := repository.NewCourseRepository(provider)

	dao := NewCourseDao(repo, repository.NewModuleRepository(provider), repository.NewUserRepository(provider), repository.NewSelectedCourseRepository(provider), repository.NewExamRepository(provider), repository.NewRoleRepository(provider), repository.NewExamtypeRepository(provider))

	courseEnt := requestmodels.RefCourse{Description: "Lol"}

	retEnt, err := dao.Create(courseEnt)
	if err != nil {
		t.Fatalf("Got db error")
	}

	// original entity must not be modified
	require.Equal(t, uint(0), courseEnt.ID)

	_, err = dao.Create(courseEnt)
	if err != nil {
		t.Fatalf("Got db error")
	}

	entities, err := dao.GetAll()
	if err != nil {
		t.Fatalf("Got db error at getAll")
	}

	t.Logf("Return %v entities", len(entities))

	found := false
	for _, v := range entities {
		if v.Description == retEnt.Description {
			found = true
		}
	}

	if found == false {
		t.Fatalf("Inserted course not found in getAll")
	}
}

// Check if a slice can be asserted correctly and keep all data intact
func TestAssertSlice(t *testing.T) {
	inputSlice := make([]any, 0)

	inputSlice = append(inputSlice, models.Module{
		Description: "Test01",
	})

	inputSlice = append(inputSlice, models.Module{
		Description: "Supertest",
	})

	outputSlice := assertSlice[models.Module](inputSlice)

	if len(outputSlice) != len(inputSlice) {
		t.Fatalf("expected slice len %v got %v", len(inputSlice), len(outputSlice))
	}

	for i := range outputSlice {
		inputModule, assertionOk := inputSlice[i].(models.Module)
		if assertionOk == false {
			t.Fatalf("can not assert input slice as module")
		}
		if outputSlice[i].Description != inputModule.Description {
			t.Fatalf("input and output description differ")
		}
	}
}

// Create default values and check for double insertion and existing
func TestCreateDefaults(t *testing.T) {
	provider := db.NewMockProvider()

	repo := repository.NewStateRepository(provider)

	dao := NewStateDao(repo)

	err := dao.CreateDefaults()
	if err != nil {
		t.Fatalf("Got db error at first")
	}

	err = dao.CreateDefaults()
	if err != nil {
		t.Fatalf("Got db error at second")
	}

	entities, err := dao.GetAll()
	if err != nil {
		t.Fatalf("Got error in getAll")
	}

	checkDescription := ""
	checkId := -1
	for i, v := range entities {
		if checkDescription == "" {
			checkDescription = v.Description
			checkId = i
		} else {
			if checkDescription == v.Description {
				t.Fatalf("Got same description '%v' on id '%v' and '%v'", v.Description, checkId, i)
			}
		}
	}

	if len(entities) != len(provider.Config().States) {
		t.Fatalf("Expected '%v' entries, got '%v'", len(entities), len(provider.Config().States))
	}
}

// check if field and focus can be matched
func TestCreateFieldAndFocus(t *testing.T) {
	provider := db.NewMockProvider()

	//provider := db.NewProvider(config.NewAPIConfig())

	fieldRepo := repository.NewFieldRepository(provider)

	fielDao := NewFieldDao(fieldRepo)

	focusRepo := repository.NewFocusRepository(provider)

	focusDao := NewFocusDao(focusRepo)

	retField, err := fielDao.Create(models.Field{
		Description: "TestField",
	})
	if err != nil {
		t.Fatal("Error at creating field")
	}

	_, err = focusDao.Create(models.Focus{
		Description: "TestFocus1",
		Field:       *retField,
	})
	if err != nil {
		t.Fatal("Error at creating first focus")
	}

	_, err = focusDao.Create(models.Focus{
		Description: "TestFocus2",
		Field:       *retField,
	})
	if err != nil {
		t.Fatal("Error at creating second focus")
	}

	focuses, err := focusDao.GetAll()
	if err != nil {
		t.Fatal("Error at getAll focus")
	}

	for _, f := range focuses {
		if f.Field.ID != retField.ID {
			t.Fatalf("on focus %v expected fieldID %v got %v", f.ID, retField.ID, f.Field.ID)
		}
	}
}

func TestGetByEmail(t *testing.T) {
	provider := db.NewPrefilledMockProvider()

	dao := NewUserDao(repository.NewUserRepository(provider), repository.NewRoleRepository(provider))

	r1, err := dao.GetByEmail("rafael.stauffer@hftm.ch")
	if err != nil {
		t.Fatal("no user found")
	}

	if r1.Name != "Rafael Stauffer" {
		t.Fatalf("username %v does not match expected", r1.Name)
	}

	r2, _ := dao.GetByEmail("does.not@exist.ch")
	if r2 != nil {
		t.Fatal("we found a non existant user")
	}
}

func TestCreateCourseVersion(t *testing.T) {
	provider := db.NewMockProvider()

	//provider := db.NewProvider(config.NewAPIConfig())

	courseDao := NewCourseDao(repository.NewCourseRepository(provider), repository.NewModuleRepository(provider), repository.NewUserRepository(provider), repository.NewSelectedCourseRepository(provider), repository.NewExamRepository(provider), repository.NewRoleRepository(provider), repository.NewExamtypeRepository(provider))

	testCourse := requestmodels.RefCourse{
		Description: "testcourse",
	}

	createdCourse, err := courseDao.Create(testCourse) // no id, no version, generate id and version 1

	require.Nil(t, err)
	require.NotEqual(t, testCourse.ID, createdCourse.ID)
	require.Equal(t, uint(1), createdCourse.Version)

	testCourse.ID = 66                                //uuid.New()
	createdCourse, err = courseDao.Create(testCourse) // set id, no version, set version 1

	require.Nil(t, err)
	require.Equal(t, testCourse.ID, createdCourse.ID)
	require.Equal(t, uint(1), createdCourse.Version)

	testCourse.Version = 2
	createdCourse, err = courseDao.Create(testCourse) // set id, set version, just create

	require.Nil(t, err)
	require.Equal(t, testCourse.ID, createdCourse.ID)
	require.Equal(t, uint(2), createdCourse.Version)

	testCourse.ID = 0 //uuid.UUID{}
	testCourse.Version = 5
	createdCourse, err = courseDao.Create(testCourse) // initial id, set version, generate random id and version 1

	require.Nil(t, err)
	require.NotEqual(t, testCourse.ID, createdCourse.ID)
	require.NotEqual(t, testCourse.Version, createdCourse.Version)
	require.Equal(t, uint(1), createdCourse.Version)
}

func TestModuleCRUD(t *testing.T) {
	provider := db.NewPrefilledMockProvider()

	dao := NewModuleDao(repository.NewModuleRepository(provider))

	testDescr := "testdescr"

	// Create
	retModule, err := dao.Create(requestmodels.RefModule{
		Description: testDescr,
	})
	require.Nil(t, err)
	require.NotEqual(t, uint(0), retModule.ID)
	require.Equal(t, testDescr, retModule.Description)

	// Read
	readModuleGet, err := dao.Get(retModule.ID, retModule.Version)
	require.Nil(t, err)
	require.Equal(t, retModule.Version, readModuleGet.Version)
	require.Equal(t, retModule.ID, readModuleGet.ID)

	readModuleGetLatest, err := dao.GetLatest(retModule.ID)
	require.Nil(t, err)
	require.Equal(t, retModule.Version, readModuleGetLatest.Version)
	require.Equal(t, retModule.ID, readModuleGetLatest.ID)

	// Update
	newTestDescr := "wowee"
	newModule := requestmodels.RefModule{
		Description: newTestDescr,
	}
	newModule.ID = retModule.ID
	newModule.Version = retModule.Version

	err = dao.Update(newModule)
	require.Nil(t, err)

	updatedModule, err := dao.Get(retModule.ID, retModule.Version)
	require.Nil(t, err)
	require.Equal(t, retModule.Version, readModuleGet.Version)
	require.Equal(t, retModule.ID, readModuleGet.ID)
	require.Equal(t, newTestDescr, updatedModule.Description)

	// Delete
	err = dao.Delete(retModule.ID, retModule.Version)
	require.Nil(t, err)
	_, err = dao.Get(retModule.ID, retModule.Version)
	require.NotNil(t, err)
}

/*
func TestExamCRUD(t *testing.T) {
	provider := db.NewPrefilledMockProvider()

	dao := NewExamDao(repository.NewExamRepository(provider))

	testDescr := "testdescr"

	// Create
	retExam, err := dao.Create(models.Exam{
		Description: testDescr,
	})
	require.Nil(t, err)
	require.NotEqual(t, uint(0), retExam.ID)
	require.Equal(t, testDescr, retExam.Description)

	// Read
	readExamGet, err := dao.Get(retExam.ID)
	require.Nil(t, err)
	require.Equal(t, retExam.ID, readExamGet.ID)

	// Update
	newTestDescr := "wowee"
	newExam := models.Exam{
		Description: newTestDescr,
	}
	newExam.ID = retExam.ID

	err = dao.Update(newExam)
	require.Nil(t, err)

	updatedModule, err := dao.Get(retExam.ID)
	require.Nil(t, err)
	require.Equal(t, retExam.ID, readExamGet.ID)
	require.Equal(t, newTestDescr, updatedModule.Description)

	// Delete
	err = dao.Delete(retExam.ID)
	require.Nil(t, err)
	_, err = dao.Get(retExam.ID)
	require.NotNil(t, err)
}
*/

// Check if a full object is linked by gorm if only the key is set in it
func TestLinkCourseObjectsByKey(t *testing.T) {
	provider := db.NewMockProvider()

	//provider := db.NewProvider(config.NewAPIConfig())

	courseDao := NewCourseDao(repository.NewCourseRepository(provider), repository.NewModuleRepository(provider), repository.NewUserRepository(provider), repository.NewSelectedCourseRepository(provider), repository.NewExamRepository(provider), repository.NewRoleRepository(provider), repository.NewExamtypeRepository(provider))

	// create a module and link it indirectly with the course
	moduleDao := NewModuleDao(repository.NewModuleRepository(provider))

	testModule := requestmodels.RefModule{
		Description: "testmodule",
	}

	retEnt, err := moduleDao.Create(testModule)
	require.Nil(t, err)

	// create a user to link as teacher
	userDao := NewUserDao(repository.NewUserRepository(provider), repository.NewRoleRepository(provider))

	testUser := models.User{
		Name: "Rafael Stauffer",
	}

	retEntUser, err := userDao.Create(testUser)
	require.Nil(t, err)

	// create a course and add all indirect key only structs to it
	testCourse := requestmodels.RefCourse{
		Description: "testcourse",
	}

	testCourse.Modules = append(testCourse.Modules, requestmodels.RefVersioned{ID: retEnt.ID, Version: retEnt.Version})
	testCourse.TeachedBy = append(testCourse.TeachedBy, requestmodels.RefId{ID: retEntUser.ID})

	createdCourse, err := courseDao.Create(testCourse)
	require.Nil(t, err)

	// check has our main course object kept its data
	require.Equal(t, "testcourse", createdCourse.Description)

	// check is our linked module complete
	require.Equal(t, retEnt.ID, createdCourse.Modules[0].ID)
	require.Equal(t, retEnt.Version, createdCourse.Modules[0].Version)
	require.Equal(t, testModule.Description, createdCourse.Modules[0].Description)

	// check is our linked teacher complete
	require.Equal(t, retEntUser.ID, createdCourse.TeachedBy[0].ID)
	require.Equal(t, retEntUser.Name, createdCourse.TeachedBy[0].Name)
}

// Check error if a reference is missing
func TestErrorAtNonexistingLink(t *testing.T) {
	provider := db.NewMockProvider()

	//provider := db.NewProvider(config.NewAPIConfig())

	courseDao := NewCourseDao(repository.NewCourseRepository(provider), repository.NewModuleRepository(provider), repository.NewUserRepository(provider), repository.NewSelectedCourseRepository(provider), repository.NewExamRepository(provider), repository.NewRoleRepository(provider), repository.NewExamtypeRepository(provider))

	nonexistantOnlyIDModule := requestmodels.RefVersioned{}
	nonexistantOnlyIDModule.ID = 234
	nonexistantOnlyIDModule.Version = 12

	testCourse := requestmodels.RefCourse{
		Description: "testcourse",
	}

	testCourse.Modules = append(testCourse.Modules, nonexistantOnlyIDModule)

	retEnt, err := courseDao.Create(testCourse)
	require.NotEmpty(t, err)
	require.Empty(t, retEnt)
	t.Logf("Error returned %v", err.Msg)
}

// Check error if a validation failed
func TestErrorAtValidationError(t *testing.T) {
	provider := db.NewMockProvider()

	//provider := db.NewProvider(config.NewAPIConfig())

	courseDao := NewCourseDao(repository.NewCourseRepository(provider), repository.NewModuleRepository(provider), repository.NewUserRepository(provider), repository.NewSelectedCourseRepository(provider), repository.NewExamRepository(provider), repository.NewRoleRepository(provider), repository.NewExamtypeRepository(provider))

	testCourse := requestmodels.RefCourse{
		Description: "testcourse",
	}

	testCourse.SelectedCourses = append(testCourse.SelectedCourses, requestmodels.RefSelectedCourse{})

	retEnt, err := courseDao.Create(testCourse)
	require.NotEmpty(t, err)
	require.Empty(t, retEnt)
	t.Logf("Error returned %v", err.Msg)
}

// check if referenced objects are preserved if a course is deleted
func TestPreventCascadeDelete(t *testing.T) {
	provider := db.NewMockProvider()

	//provider := db.NewProvider(config.NewAPIConfig())

	courseDao := NewCourseDao(repository.NewCourseRepository(provider), repository.NewModuleRepository(provider), repository.NewUserRepository(provider), repository.NewSelectedCourseRepository(provider), repository.NewExamRepository(provider), repository.NewRoleRepository(provider), repository.NewExamtypeRepository(provider))

	// create a module and link it indirectly with the course
	moduleDao := NewModuleDao(repository.NewModuleRepository(provider))

	testModule := requestmodels.RefModule{
		Description: "testmodule",
	}

	retEnt, err := moduleDao.Create(testModule)
	require.Nil(t, err)

	testCourse := requestmodels.RefCourse{
		Description: "testcourse",
	}

	testCourse.Modules = append(testCourse.Modules, requestmodels.RefVersioned{ID: retEnt.ID, Version: retEnt.Version})

	retCourseEnt, err := courseDao.Create(testCourse)
	require.Nil(t, err)

	err = courseDao.Delete(retCourseEnt.ID, retCourseEnt.Version)
	require.Nil(t, err)

	retCheckModule, err := moduleDao.Get(retEnt.ID, retEnt.Version)
	require.Nil(t, err)
	require.Equal(t, testModule.Description, retCheckModule.Description)
}

/*
// check if creating and deleting of exams is reflected in the course
func TestCourseExamLink(t *testing.T) {
	provider := db.NewMockProvider()

	courseDao := NewCourseDao(repository.NewCourseRepository(provider), repository.NewModuleRepository(provider), repository.NewUserRepository(provider), repository.NewSelectedCourseRepository(provider), repository.NewExamRepository(provider), repository.NewRoleRepository(provider), repository.NewExamtypeRepository(provider))

	examDao := NewExamDao(repository.NewExamRepository(provider))

	testExam := models.Exam{
		Description: "testexam",
	}

	testExamSecondary := models.Exam{
		Description: "another one",
	}

	retExam, err := examDao.Create(testExam)
	require.Nil(t, err)

	retExamSecondary, err := examDao.Create(testExamSecondary)
	require.Nil(t, err)

	testRefCourse := requestmodels.RefCourse{
		Description: "testcourse",
		Exams:       []requestmodels.RefId{{ID: retExam.ID}, {ID: retExamSecondary.ID}},
	}

	retCourse, err := courseDao.Create(testRefCourse)
	require.Nil(t, err)

	t.Logf("len of created course exams is %v", len(retCourse.Exams))
	require.True(t, len(retCourse.Exams) == 2)
	require.Equal(t, testExam.Description, retCourse.Exams[0].Description) // TODO: refactor this fragile test

	// check if the exam now has the course registered in the orm
	retExamSecondaryCheck, err := examDao.Get(retExamSecondary.ID)
	require.Nil(t, err)

	require.Equal(t, retCourse.ID, retExamSecondaryCheck.Course.ID)
	require.Equal(t, retCourse.Version, retExamSecondaryCheck.Course.Version)

	// check if the exam is correctly removed from the course once deleted
	err = examDao.Delete(retExam.ID)
	require.Nil(t, err)

	retCourseCheck, err := courseDao.Get(retCourse.ID, retCourse.Version)
	require.Nil(t, err)

	t.Logf("len of created course exams is %v", len(retCourseCheck.Exams))
	require.True(t, len(retCourseCheck.Exams) == 1)
	require.Equal(t, testExamSecondary.Description, retCourseCheck.Exams[0].Description) // TODO: refactor this fragile test
}
*/

// check if roles get created and a duplicated creation is prevented
func TestCreateDefaultRoles(t *testing.T) {
	provider := db.NewMockProvider()
	userDao := NewUserDao(repository.NewUserRepository(provider), repository.NewRoleRepository(provider))

	expectedNumberOfRoles := len(provider.Config().Roles)

	err := userDao.CreateDefaults()
	require.Nil(t, err)

	// quick check do we have the same number of roles
	roleEnts, internalErr := userDao.roleRepo.GetAll()
	require.Nil(t, internalErr)
	require.Len(t, roleEnts, expectedNumberOfRoles)

	err = userDao.CreateDefaults()
	require.Nil(t, err)

	// quick check again do we have the same number of roles
	roleEnts, internalErr = userDao.roleRepo.GetAll()
	require.Nil(t, internalErr)
	require.Len(t, roleEnts, expectedNumberOfRoles)

	// now check in detail if the roles match
	for _, expectedRole := range provider.Config().Roles {
		roleEnt, internalErr := userDao.roleRepo.GetId(expectedRole.Id)
		require.Nil(t, internalErr)
		role := roleEnt.(*models.Role)
		require.Equal(t, expectedRole.ClaimName, role.Claim)
	}
}

// check if a user can created and updated
func TestCreateUpdateUser(t *testing.T) {
	provider := db.NewMockProvider()
	userDao := NewUserDao(repository.NewUserRepository(provider), repository.NewRoleRepository(provider))

	testUser := models.User{
		Name: "Test",
	}

	ret, err := userDao.Create(testUser)
	require.Nil(t, err)
	require.Equal(t, testUser.Name, ret.Name)

	testUserUpdate := models.User{
		Name: "TestAgain",
	}
	testUserUpdate.ID = ret.ID

	err = userDao.Update(testUserUpdate)
	require.Nil(t, err)

	retUpdated, err := userDao.Get(ret.ID)
	require.Nil(t, err)
	require.Equal(t, testUserUpdate.Name, retUpdated.Name)

}

// check if create and update user by email works
func TestCreateUpdateUserByEmail(t *testing.T) {
	provider := db.NewMockProvider()
	userDao := NewUserDao(repository.NewUserRepository(provider), repository.NewRoleRepository(provider))
	userDao.CreateDefaults()

	roles := make([]*models.Role, 0)

	dozentRole, err := userDao.GetRoleByClaim("Dozent")
	require.Nil(t, err)

	roles = append(roles, dozentRole)

	testUser := models.User{
		Name:  "Test",
		Email: "stafi@hftm.ch",
		Roles: roles,
	}

	retCreated, err := userDao.CreateOrUpdateByEmail(testUser)
	require.Nil(t, err)
	require.Equal(t, testUser.Email, retCreated.Email)
	require.Equal(t, testUser.Name, retCreated.Name)

	studentRole, err := userDao.GetRoleByClaim("Student")
	require.Nil(t, err)

	roles = append(roles, studentRole)

	testUserUpdate := models.User{
		Name:  "Woopsiedoopsie",
		Email: "stafi@hftm.ch",
		Roles: roles,
	}

	retUpdated, err := userDao.CreateOrUpdateByEmail(testUserUpdate)
	require.Nil(t, err)
	require.Equal(t, testUserUpdate.Email, retUpdated.Email)
	require.Equal(t, testUserUpdate.Name, retUpdated.Name)

	// be paranoid and check with a fresh get by id
	retGet, err := userDao.Get(retUpdated.ID)
	require.Nil(t, err)
	require.Equal(t, testUserUpdate.Email, retGet.Email)
	require.Equal(t, testUserUpdate.Name, retGet.Name)

	require.Len(t, retGet.Roles, len(roles)) // do the roles match up in length
	for i := range roles {
		require.Contains(t, retGet.Roles, roles[i]) // check if each role is in the role list
	}
}

func TestCreateCurriculum(t *testing.T) {
	provider := db.NewMockProvider()
	curriculumDao := NewCurriculumDao(
		repository.NewCurriculumRepository(provider),
		repository.NewFocusRepository(provider),
		repository.NewCurriculumtypeRepository(provider),
		repository.NewStateRepository(provider),
		repository.NewModuleRepository(provider),
	)

	// Prefill Data with mock data for the Curriculum
	focus := db.Focus1
	curriculumType := db.CurriculumTyp3
	state := db.State1
	module := models.Module{
		State:       state,
		Description: "Schnittstellen-Technologien",
		Number:      "IN123",
		VersionedBasemodel: models.VersionedBasemodel{
			Version: 2,
			ID:      2,
		},
	}
	provider.DB().Table("focus").Create(&focus)
	provider.DB().Table("curriculumtypes").Create(&curriculumType)
	provider.DB().Table("states").Create(&state)
	provider.DB().Table("modules").Create(&module)

	mustHaveId, _ := curriculumDao.repo.GetNextId()
	// Create a new Curriculum and add all Values
	curriculum := requestmodels.RefCurriculum{
		Description: "Softwareentwicklung",
		EndValidity: "01.02.2024",
	}
	curriculum.ID = 0
	curriculum.StartValidity = "30.12.2025"
	curriculum.FocusID.ID = focus.ID
	curriculum.CurriculumtypeID.ID = curriculumType.ID
	curriculum.StateID.ID = state.ID
	versRef := requestmodels.RefVersioned{
		ID:      module.ID,
		Version: module.Version,
	}
	curriculum.Modules = append(curriculum.Modules, versRef)

	// Create new Curriculum in Database
	createdCurriculum, err := curriculumDao.Create(curriculum)

	require.Nil(t, err)
	// Dependencies
	require.Equal(t, focus.Description, createdCurriculum.Focus.Description)
	require.Equal(t, curriculumType.Description, createdCurriculum.Curriculumtype.Description)
	require.Equal(t, state.Description, createdCurriculum.State.Description)
	require.Equal(t, module.ID, createdCurriculum.Modules[0].ID)
	require.Equal(t, module.Version, createdCurriculum.Modules[0].Version)

	// Correct Date
	require.Equal(t, time.Date(2024, time.February, 1, 0, 0, 0, 0, time.UTC), createdCurriculum.EndValidity)
	require.Equal(t, time.Date(2025, time.December, 30, 0, 0, 0, 0, time.UTC), createdCurriculum.StartValidity)

	// Correct ID
	require.Equal(t, mustHaveId, createdCurriculum.ID)
}

func TestGetAllCurriculum(t *testing.T) {
	provider := db.NewPrefilledMockProvider()
	curriculumDao := NewCurriculumDao(
		repository.NewCurriculumRepository(provider),
		repository.NewFocusRepository(provider),
		repository.NewCurriculumtypeRepository(provider),
		repository.NewStateRepository(provider),
		repository.NewModuleRepository(provider),
	)

	curriculums, err := curriculumDao.GetAll()

	require.Nil(t, err)

	// There are two curriculums in the prefild Mockdatabase
	require.Equal(t, 2, len(curriculums))

	// Dependencies
	// Curriculum 1
	require.Equal(t, db.Focus1.Description, curriculums[0].Focus.Description)
	require.Equal(t, db.CurriculumTyp3.Description, curriculums[0].Curriculumtype.Description)
	require.Equal(t, db.State1.Description, curriculums[0].State.Description)
	require.Equal(t, db.Module1.ID, curriculums[0].Modules[0].ID)
	require.Equal(t, db.Module2.ID, curriculums[0].Modules[1].ID)
	// Curriculum 2
	require.Equal(t, db.Focus2.Description, curriculums[1].Focus.Description)
	require.Equal(t, db.CurriculumTyp2.Description, curriculums[1].Curriculumtype.Description)
	require.Equal(t, db.State1.Description, curriculums[1].State.Description)
	require.Equal(t, db.Module1.ID, curriculums[1].Modules[0].ID)
	require.Equal(t, db.Module2.ID, curriculums[1].Modules[1].ID)

	// Correct Date
	// Curriculum 1
	require.Equal(t, db.Curriculum1.StartValidity, curriculums[0].StartValidity)
	require.Equal(t, db.Curriculum1.EndValidity, curriculums[0].EndValidity)
	// Curriculum 2
	require.Equal(t, db.Curriculum2.StartValidity, curriculums[1].StartValidity)
	require.Equal(t, db.Curriculum2.EndValidity, curriculums[1].EndValidity)

	// Correct ID
	require.Equal(t, db.Curriculum1.ID, curriculums[0].ID)
	require.Equal(t, db.Curriculum2.ID, curriculums[1].ID)
}

func TestGetOneCurriculum(t *testing.T) {
	provider := db.NewPrefilledMockProvider()
	curriculumDao := NewCurriculumDao(
		repository.NewCurriculumRepository(provider),
		repository.NewFocusRepository(provider),
		repository.NewCurriculumtypeRepository(provider),
		repository.NewStateRepository(provider),
		repository.NewModuleRepository(provider),
	)

	curriculum, err := curriculumDao.Get(2, time.Date(2021, time.April, 1, 0, 0, 0, 0, time.UTC))

	require.Nil(t, err)

	// Dependencies
	require.Equal(t, db.Focus2.Description, curriculum.Focus.Description)
	require.Equal(t, db.CurriculumTyp2.Description, curriculum.Curriculumtype.Description)
	require.Equal(t, db.State1.Description, curriculum.State.Description)
	require.Equal(t, db.Module1.ID, curriculum.Modules[0].ID)
	require.Equal(t, db.Module2.ID, curriculum.Modules[1].ID)

	// Correct Date
	require.Equal(t, db.Curriculum2.StartValidity, curriculum.StartValidity)
	require.Equal(t, db.Curriculum2.EndValidity, curriculum.EndValidity)

	// Correct ID
	require.Equal(t, db.Curriculum2.ID, curriculum.ID)
}

func TestUpdateCurriculum(t *testing.T) {
	provider := db.NewPrefilledMockProvider()
	curriculumDao := NewCurriculumDao(
		repository.NewCurriculumRepository(provider),
		repository.NewFocusRepository(provider),
		repository.NewCurriculumtypeRepository(provider),
		repository.NewStateRepository(provider),
		repository.NewModuleRepository(provider),
	)

	// Change a Curriculum
	refCurriculum := requestmodels.RefCurriculum{
		Description: "Elektroniker",
		EndValidity: "23.07.2020",
	}
	refCurriculum.FocusID.ID = db.Focus1.ID
	refCurriculum.CurriculumtypeID.ID = db.CurriculumTyp3.ID
	refCurriculum.StateID.ID = db.State2.ID
	refCurriculum.ID = 2
	refCurriculum.StartValidity = "01.04.2021"

	// Update Curriculum2 (from mock)
	err := curriculumDao.Update(refCurriculum)
	require.Nil(t, err)

	// Get Curriculum2
	curriculum, err := curriculumDao.Get(2, time.Date(2021, time.April, 1, 0, 0, 0, 0, time.UTC))

	require.Nil(t, err)

	// Dependencies
	require.Equal(t, db.Focus1.Description, curriculum.Focus.Description)
	require.Equal(t, db.CurriculumTyp3.Description, curriculum.Curriculumtype.Description)
	require.Equal(t, db.State2.Description, curriculum.State.Description)
	require.Equal(t, db.Module1.ID, curriculum.Modules[0].ID)
	require.Equal(t, db.Module2.ID, curriculum.Modules[1].ID)

	// Correct Date
	require.Equal(t, db.Curriculum2.StartValidity, curriculum.StartValidity)
	require.Equal(t, time.Date(2020, time.July, 23, 0, 0, 0, 0, time.UTC), curriculum.EndValidity)

	// Correct ID
	require.Equal(t, db.Curriculum2.ID, curriculum.ID)
}

func TestDeleteCurriculum(t *testing.T) {
	provider := db.NewPrefilledMockProvider()
	curriculumDao := NewCurriculumDao(
		repository.NewCurriculumRepository(provider),
		repository.NewFocusRepository(provider),
		repository.NewCurriculumtypeRepository(provider),
		repository.NewStateRepository(provider),
		repository.NewModuleRepository(provider),
	)

	// Get All Curriculums to get Length
	curriculums, err := curriculumDao.GetAll()
	require.Nil(t, err)

	lengthCurriculums := len(curriculums)

	// Delete Curriculums
	for _, c := range curriculums {
		err := curriculumDao.Delete(c.ID, c.StartValidity)
		require.Nil(t, err)
		lengthCurriculums = lengthCurriculums - 1
	}

	// Get New lenth of Curriculums
	curriculumsAfterDelete, err := curriculumDao.GetAll()
	require.Nil(t, err)

	require.Equal(t, lengthCurriculums, len(curriculumsAfterDelete))

}

// Check if users can be selected by role
func TestGetUsersByRole(t *testing.T) {
	provider := db.NewMockProvider()
	userDao := NewUserDao(repository.NewUserRepository(provider), repository.NewRoleRepository(provider))
	userDao.CreateDefaults()

	role, err := userDao.GetRoleByClaim("Kursadministrator")
	require.Nil(t, err)

	roles := []*models.Role{role}

	for i := 0; i < 10; i++ {
		userDao.Create(models.User{Roles: roles, Email: fmt.Sprintf("test%v@hftm.ch", i)})
	}

	ents, err := userDao.GetByRole(1)
	require.Nil(t, err)

	require.Len(t, ents, 10)
}

// check if we can select a curriculum by timepoint
func TestCurriculumValidityTimeSelection(t *testing.T) {
	provider := db.NewMockProvider()
	cDao := NewCurriculumDao(repository.NewCurriculumRepository(provider), repository.NewFocusRepository(provider), repository.NewCurriculumtypeRepository(provider), repository.NewStateRepository(provider), repository.NewModuleRepository(provider))

	cRef := requestmodels.RefCurriculum{}
	cRef.ID = 5
	cRef.StartValidity = "01.01.2022"
	cRef.EndValidity = "01.01.2024"

	retEnt, err := cDao.Create(cRef)
	require.Nil(t, err)
	require.Equal(t, cRef.ID, retEnt.ID)

	cRef.StartValidity = "02.01.2024"
	cRef.EndValidity = "05.01.2024"

	retEnt2, err := cDao.Create(cRef)
	require.Nil(t, err)
	require.Equal(t, cRef.ID, retEnt2.ID)

	cRef.StartValidity = "06.01.2024"
	cRef.EndValidity = ""

	retEnt3, err := cDao.Create(cRef)
	require.Nil(t, err)
	require.Equal(t, cRef.ID, retEnt3.ID)

	// get the valid entry for 03.01.24 (should be the one from 02.01.24 to 05.01.24)
	tPoint, intErr := ParseTime("03.01.2024", "02.01.2006")
	require.Nil(t, intErr)
	cRetValid, err := cDao.GetValidForTimepoint(cRef.ID, tPoint)
	require.Nil(t, err)
	require.Equal(t, cRef.ID, cRetValid.ID)
	require.Equal(t, retEnt2.StartValidity, cRetValid.StartValidity)
	require.Equal(t, retEnt2.EndValidity, cRetValid.EndValidity)

	// get the valid entry for 21.01.24 (should be the one from 06.01.24)
	tPoint, intErr = ParseTime("21.01.2024", "02.01.2006")
	require.Nil(t, intErr)
	cRetValid, err = cDao.GetValidForTimepoint(cRef.ID, tPoint)
	require.Nil(t, err)
	require.Equal(t, cRef.ID, cRetValid.ID)
	require.Equal(t, retEnt3.StartValidity, cRetValid.StartValidity)
	require.Equal(t, retEnt3.EndValidity, cRetValid.EndValidity)

	// get the valid entry for 01.01.20 (should be nil)
	tPoint, intErr = ParseTime("01.01.2020", "02.01.2006")
	require.Nil(t, intErr)
	_, err = cDao.GetValidForTimepoint(cRef.ID, tPoint)
	require.NotNil(t, err)
}

// Check if courses and studends are returned from the teaches by user field
func TestTechedByTeacher(t *testing.T) {
	provider := db.NewMockProvider()

	//provider := db.NewProvider(config.NewAPIConfig())

	courseDao := NewCourseDao(repository.NewCourseRepository(provider), repository.NewModuleRepository(provider), repository.NewUserRepository(provider), repository.NewSelectedCourseRepository(provider), repository.NewExamRepository(provider), repository.NewRoleRepository(provider), repository.NewExamtypeRepository(provider))

	// create a module and link it indirectly with the course
	moduleDao := NewModuleDao(repository.NewModuleRepository(provider))

	testModule := requestmodels.RefModule{
		Description: "testmodule",
	}

	retEnt, err := moduleDao.Create(testModule)
	require.Nil(t, err)

	// create a user to link as teacher
	userDao := NewUserDao(repository.NewUserRepository(provider), repository.NewRoleRepository(provider))

	testUser := models.User{
		Name:  "Rafael Stauffer",
		Email: "rafael.stauffer@hftm.ch",
	}

	retEntUser, err := userDao.Create(testUser)
	require.Nil(t, err)

	// create a course and add all indirect key only structs to it
	testCourse := requestmodels.RefCourse{
		Description: "testcourse",
	}

	testCourse.Modules = append(testCourse.Modules, requestmodels.RefVersioned{ID: retEnt.ID, Version: retEnt.Version})
	testCourse.TeachedBy = append(testCourse.TeachedBy, requestmodels.RefId{ID: retEntUser.ID})

	createdCourse, err := courseDao.Create(testCourse)
	require.Nil(t, err)

	// check has our main course object kept its data
	require.Equal(t, "testcourse", createdCourse.Description)

	startToday := time.Now()

	testStudent := models.User{
		Name:           "Max Mustermann",
		Email:          "max.mustermann@hftm.ch",
		ClassStartyear: startToday,
	}

	retEntUserStudent, err := userDao.Create(testStudent) // create first to get an id
	require.Nil(t, err)

	// now we can create the selected course link
	retEntUserStudent.SelectedCourses = []models.SelectedCourse{{CourseID: createdCourse.ID, CourseVersion: createdCourse.Version, ClassStartyear: startToday, UserID: retEntUserStudent.ID}}

	err = userDao.Update(*retEntUserStudent)
	require.Nil(t, err)

	retEntUserStudent, err = userDao.Get(retEntUserStudent.ID)
	require.Nil(t, err)

	require.True(t, len(retEntUserStudent.SelectedCourses) > 0) // This works, the user has now assigned a selected course
	require.Equal(t, createdCourse.ID, retEntUserStudent.SelectedCourses[0].CourseID)

	retEntUser, err = userDao.Get(retEntUser.ID) // now select the teacher again to check if we can see the teached module and its selected course link
	require.Nil(t, err)

	require.True(t, len(retEntUser.TeachesCourses) > 0)
	// This fails, the selected courses does not autoload
	// require.True(t, len(retEntUser.TeachesCourses[0].SelectedCourses) > 0)
	//require.Equal(t, retEntUserStudent.ID, retEntUser.TeachesCourses[0].SelectedCourses[0].UserID)
}

func TestClassSelection(t *testing.T) {
	provider := db.NewPrefilledMockProvider()

	classDao := NewClassDao(repository.NewCourseRepository(provider), repository.NewUserRepository(provider), repository.NewSelectedCourseRepository(provider), repository.NewExamRepository(provider), repository.NewRoleRepository(provider), repository.NewExamEvaluationRepository(provider), repository.NewExamtypeRepository(provider))

	class, err := classDao.Get(db.SelectedCourse1.CourseID, db.SelectedCourse1.CourseVersion, db.SelectedCourse1.ClassStartyear)
	require.Nil(t, err)
	require.NotNil(t, class)
}

// check if a list of teachers can be extracted from a prefilled mock provider to check for role id bugs
func TestGetTeachersForMeta(t *testing.T) {
	provider := db.NewPrefilledMockProvider()

	userDao := NewUserDao(repository.NewUserRepository(provider), repository.NewRoleRepository(provider))

	err := userDao.CreateDefaults() // create the roles
	require.Nil(t, err)

	// Get all Teachers
	teachers, err := userDao.GetTeachers()
	require.Nil(t, err)
	require.NotEmpty(t, teachers)
}

func TestCreateModule(t *testing.T) {
	provider := db.NewMockProvider()
	moduleDao := NewModuleDao(
		repository.NewModuleRepository(provider),
	)

	// Prefill Data with mock data for the module
	state := db.State1
	studyStage := db.StudyStage2
	evaluationType := db.EvaluationType1
	curriculum1 := db.Curriculum1
	curriculum2 := db.Curriculum2
	provider.DB().Table("states").Create(&state)
	provider.DB().Table("study_stages").Create(&studyStage)
	provider.DB().Table("evaluationtypes").Create(&evaluationType)
	provider.DB().Table("curriculums").Create(&curriculum1)
	provider.DB().Table("curriculums").Create(&curriculum2)

	mustHaveId, _ := moduleDao.repo.GetNextId()
	// Create a new Module and add Values
	module := requestmodels.RefModule{
		Description: "Schnittstellen-Technologien",
		Number:      "IN123",
	}
	module.State.ID = state.ID
	module.StudyStage.ID = studyStage.ID
	module.EvaluationType.ID = evaluationType.ID
	module.Curriculums = []requestmodels.RefTimed{
		{
			ID:            curriculum1.ID,
			StartValidity: "01.04.2021",
		},
		{
			ID:            curriculum2.ID,
			StartValidity: "01.04.2021",
		},
	}

	// Create new Module in Database
	createdModule, err := moduleDao.Create(module)

	require.Nil(t, err)
	// Dependencies
	require.Equal(t, state.Description, createdModule.State.Description)
	require.Equal(t, studyStage.Description, createdModule.StudyStage.Description)
	require.Equal(t, evaluationType.Description, createdModule.EvaluationType.Description)
	require.Equal(t, curriculum1.ID, createdModule.Curriculums[0].ID)
	require.Equal(t, curriculum2.ID, createdModule.Curriculums[1].ID)
	require.Equal(t, curriculum1.StartValidity, createdModule.Curriculums[0].StartValidity)
	require.Equal(t, curriculum2.StartValidity, createdModule.Curriculums[1].StartValidity)

	// Correct ID
	require.Equal(t, mustHaveId, createdModule.ID)
	// The Version has to be 1
	require.Equal(t, uint(1), createdModule.Version)

	// Create another Module with same ID
	module.ID = createdModule.ID

	createVersionModule, err := moduleDao.Create(module)

	require.Nil(t, err)
	// Dependencies
	require.Equal(t, state.Description, createVersionModule.State.Description)
	require.Equal(t, studyStage.Description, createVersionModule.StudyStage.Description)
	require.Equal(t, evaluationType.Description, createVersionModule.EvaluationType.Description)
	require.Equal(t, curriculum1.ID, createVersionModule.Curriculums[0].ID)
	require.Equal(t, curriculum2.ID, createVersionModule.Curriculums[1].ID)
	require.Equal(t, curriculum1.StartValidity, createVersionModule.Curriculums[0].StartValidity)
	require.Equal(t, curriculum2.StartValidity, createVersionModule.Curriculums[1].StartValidity)

	// Correct ID
	require.Equal(t, mustHaveId, createVersionModule.ID)
	// The Version has to be 2
	require.Equal(t, uint(2), createVersionModule.Version)
}

func TestGetAllModules(t *testing.T) {
	provider := db.NewPrefilledMockProvider()
	moduleDao := NewModuleDao(repository.NewModuleRepository(provider))

	modules, err := moduleDao.GetAll()

	require.Nil(t, err)

	// There are two modules in the prefild Mockdatabase
	require.Equal(t, 2, len(modules))

	// Dependencies
	// Module 1
	require.Equal(t, db.State1.Description, modules[0].State.Description)
	require.Equal(t, db.StudyStage1.Description, modules[0].StudyStage.Description)
	require.Equal(t, db.EvaluationType1.Description, modules[0].EvaluationType.Description)
	require.Equal(t, db.Curriculum1.ID, modules[0].Curriculums[0].ID)
	require.Equal(t, db.Curriculum2.ID, modules[0].Curriculums[1].ID)
	// Module 2
	require.Equal(t, db.State1.Description, modules[1].State.Description)
	require.Equal(t, db.StudyStage1.Description, modules[1].StudyStage.Description)
	require.Equal(t, db.EvaluationType1.Description, modules[1].EvaluationType.Description)
	require.Equal(t, db.Curriculum1.ID, modules[1].Curriculums[0].ID)
	require.Equal(t, db.Curriculum2.ID, modules[1].Curriculums[1].ID)

	// Correct ID and Version
	require.Equal(t, db.Module1.ID, modules[0].ID)
	require.Equal(t, db.Module1.Version, modules[0].Version)

	require.Equal(t, db.Module2.ID, modules[1].ID)
	require.Equal(t, db.Module2.Version, modules[1].Version)
}

func TestGetOneModule(t *testing.T) {
	provider := db.NewPrefilledMockProvider()
	moduleDao := NewModuleDao(repository.NewModuleRepository(provider))

	module, err := moduleDao.Get(2, 1)

	require.Nil(t, err)

	// Dependencies
	require.Equal(t, db.State1.Description, module.State.Description)
	require.Equal(t, db.StudyStage1.Description, module.StudyStage.Description)
	require.Equal(t, db.EvaluationType1.Description, module.EvaluationType.Description)
	require.Equal(t, db.Curriculum1.ID, module.Curriculums[0].ID)
	require.Equal(t, db.Curriculum2.ID, module.Curriculums[1].ID)

	// Correct ID and Version
	require.Equal(t, db.Module2.ID, module.ID)
	require.Equal(t, db.Module2.Version, module.Version)
}

func TestUpdateModule(t *testing.T) {
	provider := db.NewPrefilledMockProvider()
	moduleDao := NewModuleDao(repository.NewModuleRepository(provider))

	// Change a module
	refModule := requestmodels.RefModule{
		Description: "Kleines 1x1",
	}
	refModule.State.ID = db.State2.ID
	refModule.StudyStage.ID = db.StudyStage2.ID
	refModule.EvaluationType.ID = db.EvaluationType2.ID
	refModule.Curriculums = []requestmodels.RefTimed{
		{
			ID:            db.Curriculum1.ID,
			StartValidity: "01.04.2021",
		},
	}

	refModule.ID = 2
	refModule.Version = 1

	// Update Module2
	err := moduleDao.Update(refModule)
	require.Nil(t, err)

	module, err := moduleDao.Get(2, 1)

	require.Nil(t, err)

	// Dependencies
	require.Equal(t, db.State2.Description, module.State.Description)
	require.Equal(t, db.StudyStage2.Description, module.StudyStage.Description)
	require.Equal(t, db.EvaluationType2.Description, module.EvaluationType.Description)
	require.Equal(t, db.Curriculum1.ID, module.Curriculums[0].ID)
}

func TestDeleteModule(t *testing.T) {
	provider := db.NewPrefilledMockProvider()
	moduleDao := NewModuleDao(repository.NewModuleRepository(provider))

	// Get all Modules to get length
	modules, err := moduleDao.GetAll()
	require.Nil(t, err)

	lengthModules := len(modules)

	// Delete Modules
	for _, c := range modules {
		err := moduleDao.Delete(c.ID, c.Version)
		require.Nil(t, err)
		lengthModules = lengthModules - 1
	}

	// Get New length of modules
	modulesAfterDelete, err := moduleDao.GetAll()
	require.Nil(t, err)

	require.Equal(t, lengthModules, len(modulesAfterDelete))
}
