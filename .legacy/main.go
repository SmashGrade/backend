package main

import (
	"fmt"

	"github.com/SmashGrade/backend/legacy/api"
	v1 "github.com/SmashGrade/backend/legacy/api/v1"
	"github.com/SmashGrade/backend/legacy/dao"
	"github.com/SmashGrade/backend/legacy/provider"
)

func main() {

	ctx := api.NewServer()

	// add all versions
	v1.RoutesV1(ctx)
	// v2...
	// v3...
	// v4...

	prov := &provider.SqliteProvider{}
	prov.Connect()
	stuff := &dao.Database{}
	stuff.Db = prov.Db

	err := ctx.Run(9000)
	if err != nil {
		fmt.Println(err) // TODO: handle error
	}

}
