package dao_test

import (
	"testing"

	"github.com/SmashGrade/backend/app/api/v1/schemas"
	"github.com/SmashGrade/backend/app/dao"
	"github.com/SmashGrade/backend/app/provider"
)

func TestCreateCurriculum(t *testing.T) {
	prov := &provider.SqliteProvider{}
	prov.Connect()
	db := dao.Database{Db: prov.Db}

	field, err := db.CreateField("a")
	if err != nil {
		t.Fatalf("Field creation threw error %v\n", err.Error())
	}

	focus, err := db.CreateFocus("A", field.ID)
	if err != nil {
		t.Fatalf("Focus creation threw error %v\n", err.Error())
	}

	curType, err := db.CreateCurriculumType("x", 3)
	if err != nil {
		t.Fatalf("CurriculumType creation threw error %v\n", err.Error())
	}

	curriculumRef := &schemas.CurriculumReq{
		Focus:           focus.Description,
		Field:           field.Description,
		CurriculumType:  curType.Description,
		IsActive:        true,
		StartDate:       "01.01.2025",
		EndDate:         "01.01.2028",
		FieldmanagerRef: []uint{},
		ModulesRef:      []uint{},
	}

	id, err := db.CreateCurriculum(curriculumRef)
	if err != nil {
		t.Fatalf("Creation threw error %v\n", err.Error())
	}

	if id < 1 {
		t.Fatalf("Impossible record id returned %v\n", id)
	}
}
