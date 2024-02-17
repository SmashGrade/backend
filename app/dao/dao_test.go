package dao

import (
	"testing"

	"github.com/SmashGrade/backend/app/db"
	_ "github.com/SmashGrade/backend/app/docs"
	"github.com/SmashGrade/backend/app/models"
	"github.com/SmashGrade/backend/app/repository"
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

	courseEnt.ID = 0

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
