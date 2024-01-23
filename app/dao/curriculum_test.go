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

	curriculumRef := &schemas.CurriculumReq{
		Focus:           "A",
		Field:           "a",
		CurriculumType:  "x",
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
