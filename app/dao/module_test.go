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

	id, err := db.CreateModule(testModule)

	if err != nil {
		t.Fatalf("Could not create module, err: %v\n", err.Error())
	}

	db.GetModule(, id, )
}
