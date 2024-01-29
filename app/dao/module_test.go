package dao_test

import (
	"testing"

	"github.com/SmashGrade/backend/app/api/v1/schemas"
	"github.com/SmashGrade/backend/app/dao"
	"github.com/SmashGrade/backend/app/provider"
)

// creates a new module
func TestCreateModule(t *testing.T) {
	prov := &provider.SqliteProvider{}
	prov.Connect()
	db := dao.Database{Db: prov.Db}

	testModule := &schemas.ModuleReq{
		Description: "SuperTest",
	}

	module, err := db.CreateModule(testModule)

	if err != nil {
		t.Fatalf("Could not create module, err: %v\n", err.Error())
	}

	if module.ID <= 0 {
		t.Fatalf("Module id '%v' wrong", module.ID)
	}

	if module.Description != testModule.Description {
		t.Fatalf("Returned module description '%v' does not match example '%v'", module.Description, testModule.Description)
	}
}
