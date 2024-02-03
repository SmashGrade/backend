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

// check if autoincrement reacts correctly on multiple module updates with versions
func TestUpdateModule(t *testing.T) {
	prov := &provider.SqliteProvider{}
	prov.Connect()
	db := dao.Database{Db: prov.Db}

	testModule := &schemas.ModuleReq{
		Description: "Special",
	}

	retCreate1, err := db.CreateModule(testModule)
	if err != nil {
		t.Fatalf("Error first creation %v", err.Error())
	}
	if retCreate1.Version != 1 {
		t.Fatalf("Error first version expected '1' got '%v'", retCreate1.Version)
	}

	retCreate2, err := db.CreateModule(testModule)
	if err != nil {
		t.Fatalf("Error second creation %v", err.Error())
	}
	if retCreate2.Version != 1 {
		t.Fatalf("Error second version expected '1' got '%v'", retCreate2.Version)
	}

	if retCreate1.ID == retCreate2.ID {
		t.Fatalf("Error both created modules have the same id %v", retCreate1.ID)
	}

	updateTest := &schemas.ModuleReq{}
	updateTest.Description = "We updated that one inplace"

	updateModInplace, err := db.UpdateModule(retCreate1.ID, retCreate1.Version, updateTest)
	if err != nil {
		t.Fatalf("Error update inplace creation %v", err.Error())
	}
	if retCreate1.ID != updateModInplace.ID {
		t.Fatalf("Error update inplace; id changed expected '%v' got '%v'", retCreate1.ID, updateModInplace.ID)
	}
	if retCreate1.Version != updateModInplace.Version {
		t.Fatalf("Error update inplace; version changed expected '%v' got '%v'", retCreate1.Version, updateModInplace.Version)
	}

	updateTest.Description = "Now we add a new version"

	updateModNewVersion, err := db.UpdateModule(retCreate1.ID, retCreate1.Version+1, updateTest)
	if err != nil {
		t.Fatalf("Error update new version creation %v", err.Error())
	}
	if retCreate1.ID != updateModNewVersion.ID {
		t.Fatalf("Error update new version; id changed, expected '%v' got '%v'", retCreate1.ID, updateModNewVersion.ID)
	}
	if retCreate1.Version == updateModNewVersion.Version {
		t.Fatalf("Error update new version; version unchanged, expected '%v' got '%v'", retCreate1.Version+1, updateModNewVersion.Version)
	}
}
