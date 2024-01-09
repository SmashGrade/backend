package main

import (
	"fmt"

	"github.com/SmashGrade/backend/app/api"
	v1 "github.com/SmashGrade/backend/app/api/v1"
	"github.com/SmashGrade/backend/app/dao"
	"github.com/SmashGrade/backend/app/provider"
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
	/*
			user := schemas.User{}
			user.Name = "User5"
			user.Role = "schueler"
			stuff.PostUser(&user)

		users := []schemas.User{}
		stuff.ListUsers(&users)

			exam2 := entity.Exam{}
			exam2.Description = "more stuff"
			stuff.EditExam(&exam2, 2)

			exam3 := entity.Exam{}
			stuff.GetExam(&exam3, 2)
	*/
	// start server at Port :9000
	err := ctx.Run(9000)
	if err != nil {
		fmt.Println(err) // TODO: handle error
	}

}
