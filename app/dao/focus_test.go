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

	// create a dummy field first
	dummyField, err := db.CreateField("dummy")
	if err != nil {
		t.Fatalf("Dummy field creation threw error %v\n", err.Error())
	}
	defer db.Db.Delete(dummyField)

	focusDesc := "abc"

	focus, err := db.CreateFocus(focusDesc, dummyField.ID)
	if err != nil {
		t.Fatalf("Creation threw error %v\n", err.Error())
	}
	defer db.Db.Delete(focus)

	if focus.ID < 1 {
		t.Fatalf("Impossible record id returned %v\n", focus.ID)
	}

	selectedFocus, err := db.GetFocusByDescription(focusDesc)
	if err != nil {
		t.Fatalf("Selection threw error %v\n", err.Error())
	}

	if selectedFocus.Description != focusDesc {
		t.Fatalf("Wrong description returned %v\n", focus.Description)
	}

	if selectedFocus.ID != focus.ID {
		t.Logf("IDs differ from created %v and returned %v probably old data in DB", selectedFocus.ID, focus.ID)
	}
}
