package dao_test

import (
	"testing"

	"github.com/SmashGrade/backend/app/dao"
	"github.com/SmashGrade/backend/app/provider"
)

func TestSelectFocusByName(t *testing.T) {
	prov := &provider.SqliteProvider{}
	prov.Connect()
	db := dao.Database{Db: prov.Db}

	focusDesc := "abc"

	focus, err := db.CreateFocus(focusDesc, 1)
	if err != nil {
		t.Fatalf("Creation threw error %v\n", err.Error())
	}

	if focus.ID < 1 {
		t.Fatalf("Impossible record id returned %v\n", focus.ID)
	}

	selectedField, err := db.GetFocusByDescription(focusDesc)
	if err != nil {
		t.Fatalf("Selection threw error %v\n", err.Error())
	}

	if selectedField.Description != focusDesc {
		t.Fatalf("Wrong description returned %v\n", focus.Description)
	}

	if selectedField.ID != focus.ID {
		t.Logf("IDs differ from created %v and returned %v probably old data in DB", selectedField.ID, focus.ID)
	}
}
