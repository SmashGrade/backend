package dao

import (
	"testing"

	"github.com/SmashGrade/backend/app/db"
	_ "github.com/SmashGrade/backend/app/docs"
	"github.com/SmashGrade/backend/app/models"
	"github.com/SmashGrade/backend/app/repository"
	"github.com/google/uuid"
	_ "gorm.io/gorm"
)

// Smoketest
func TestMagicSmoke(t *testing.T) {
	provider := db.NewMockProvider()

	repo := repository.NewCourseRepository(provider)

	dao := NewCourseDao(repo)

	courseEnt := &models.Course{Description: "Lol"}

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

	courseEnt := &models.Course{Description: "Lol"}

	retEnt, err := dao.Create(courseEnt)
	if err != nil {
		t.Fatalf("Got db error")
	}

	courseEnt.ID = uuid.UUID{}

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
