package dao

import (
	"testing"

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

	dao := NewCourseDao(repo, repository.NewModuleRepository(provider), repository.NewUserRepository(provider), repository.NewSelectedCourseRepository(provider), repository.NewExamRepository(provider), repository.NewRoleRepository(provider))

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

	dao := NewCourseDao(repo, repository.NewModuleRepository(provider), repository.NewUserRepository(provider), repository.NewSelectedCourseRepository(provider), repository.NewExamRepository(provider), repository.NewRoleRepository(provider))

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

	courseDao := NewCourseDao(repository.NewCourseRepository(provider), repository.NewModuleRepository(provider), repository.NewUserRepository(provider), repository.NewSelectedCourseRepository(provider), repository.NewExamRepository(provider), repository.NewRoleRepository(provider))

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
	retModule, err := dao.Create(models.Module{
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
	newModule := models.Module{
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

func TestExamCRUD(t *testing.T) {
	provider := db.NewPrefilledMockProvider()

	dao := NewExamDao(repository.NewExamRepository(provider), repository.NewCourseRepository(provider))

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

// Check if a full object is linked by gorm if only the key is set in it
func TestLinkCourseObjectsByKey(t *testing.T) {
	provider := db.NewMockProvider()

	//provider := db.NewProvider(config.NewAPIConfig())

	courseDao := NewCourseDao(repository.NewCourseRepository(provider), repository.NewModuleRepository(provider), repository.NewUserRepository(provider), repository.NewSelectedCourseRepository(provider), repository.NewExamRepository(provider), repository.NewRoleRepository(provider))

	// create a module and link it indirectly with the course
	moduleDao := NewModuleDao(repository.NewModuleRepository(provider))

	testModule := models.Module{
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

	courseDao := NewCourseDao(repository.NewCourseRepository(provider), repository.NewModuleRepository(provider), repository.NewUserRepository(provider), repository.NewSelectedCourseRepository(provider), repository.NewExamRepository(provider), repository.NewRoleRepository(provider))

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

	courseDao := NewCourseDao(repository.NewCourseRepository(provider), repository.NewModuleRepository(provider), repository.NewUserRepository(provider), repository.NewSelectedCourseRepository(provider), repository.NewExamRepository(provider), repository.NewRoleRepository(provider))

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

	courseDao := NewCourseDao(repository.NewCourseRepository(provider), repository.NewModuleRepository(provider), repository.NewUserRepository(provider), repository.NewSelectedCourseRepository(provider), repository.NewExamRepository(provider), repository.NewRoleRepository(provider))

	// create a module and link it indirectly with the course
	moduleDao := NewModuleDao(repository.NewModuleRepository(provider))

	testModule := models.Module{
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

// check if creating and deleting of exams is reflected in the course
func TestCourseExamLink(t *testing.T) {
	provider := db.NewMockProvider()

	courseDao := NewCourseDao(repository.NewCourseRepository(provider), repository.NewModuleRepository(provider), repository.NewUserRepository(provider), repository.NewSelectedCourseRepository(provider), repository.NewExamRepository(provider), repository.NewRoleRepository(provider))

	examDao := NewExamDao(repository.NewExamRepository(provider), repository.NewCourseRepository(provider))

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
