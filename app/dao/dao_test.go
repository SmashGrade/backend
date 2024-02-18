package dao

import (
	"testing"

	"github.com/SmashGrade/backend/app/db"
	_ "github.com/SmashGrade/backend/app/docs"
	"github.com/SmashGrade/backend/app/models"
	"github.com/SmashGrade/backend/app/repository"
	"github.com/stretchr/testify/require"
	_ "gorm.io/gorm"
)

// Smoketest
func TestMagicSmoke(t *testing.T) {
	provider := db.NewMockProvider()

	repo := repository.NewCourseRepository(provider)

	dao := NewCourseDao(repo)

	courseEnt := models.Course{Description: "Lol"}

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

	dao := NewCourseDao(repo)

	courseEnt := models.Course{Description: "Lol"}

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

	dao := NewUserDao(repository.NewUserRepository(provider))

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

	courseDao := NewCourseDao(repository.NewCourseRepository(provider))

	testCourse := models.Course{
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
