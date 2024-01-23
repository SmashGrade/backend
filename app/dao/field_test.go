package dao_test

import (
	"testing"

	"github.com/SmashGrade/backend/app/dao"
	"github.com/SmashGrade/backend/app/provider"
)

func TestSelectFieldByName(t *testing.T) {
	prov := &provider.SqliteProvider{}
	prov.Connect()
	db := dao.Database{Db: prov.Db}

	fieldDesc := "abc"

	field, err := db.CreateField(fieldDesc)
	if err != nil {
		t.Fatalf("Creation threw error %v\n", err.Error())
	}

	if field.ID < 1 {
		t.Fatalf("Impossible record id returned %v\n", field.ID)
	}

	selectedField, err := db.GetFieldByDescription(fieldDesc)
	if err != nil {
		t.Fatalf("Selection threw error %v\n", err.Error())
	}

	if selectedField.Description != fieldDesc {
		t.Fatalf("Wrong description returned %v\n", field.Description)
	}

	if selectedField.ID != field.ID {
		t.Logf("IDs differ from created %v and returned %v probably old data in DB", selectedField.ID, field.ID)
	}
}
