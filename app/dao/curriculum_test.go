package dao_test

import (
	"testing"

	"github.com/SmashGrade/backend/app/api/v1/schemas"
	"github.com/SmashGrade/backend/app/dao"
	"github.com/SmashGrade/backend/app/entity"
	"github.com/SmashGrade/backend/app/provider"
)

// creates a curriculumtype, field and focus
func PrepareMinimalDataCurriculum(typeDesc, fieldDesc, focusDesc string, t *testing.T, db *dao.Database) (field *entity.Field, focus *entity.Focus, curriculumType *entity.Curriculumtype) {
	field, err := db.CreateField(fieldDesc)
	if err != nil {
		t.Fatalf("Field creation threw error %v\n", err.Error())
	}

	focus, err = db.CreateFocus(focusDesc, field.ID)
	if err != nil {
		t.Fatalf("Focus creation threw error %v\n", err.Error())
	}

	curriculumType, err = db.CreateCurriculumType(typeDesc, 3) // number of years not relevant
	if err != nil {
		t.Fatalf("CurriculumType creation threw error %v\n", err.Error())
	}
	return
}

// check if a curriculum can be created with minimal data
func TestCreateCurriculum(t *testing.T) {
	prov := &provider.SqliteProvider{}
	prov.Connect()
	db := dao.Database{Db: prov.Db}

	field, focus, curType := PrepareMinimalDataCurriculum("x", "a", "A", t, &db)

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

	ent, err := db.CreateCurriculum(curriculumRef)
	if err != nil {
		t.Fatalf("Creation threw error %v\n", err.Error())
	}

	if ent.ID < 1 {
		t.Fatalf("Impossible record id returned %v\n", ent.ID)
	}
}
