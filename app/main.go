package main

import (
	"fmt"

	"github.com/SmashGrade/backend/app/api"
	v1 "github.com/SmashGrade/backend/app/api/v1"
	"github.com/SmashGrade/backend/app/api/v1/schemas"
	"github.com/SmashGrade/backend/app/dao"
	"github.com/SmashGrade/backend/app/provider"
)

// generates minimal testdata at startup for checking
func generateTestdata(db *dao.Database) {
	field, _ := db.CreateField("xField")

	focus, _ := db.CreateFocus("focus", field.ID)

	curriculumType, _ := db.CreateCurriculumType("type", 3) // number of years not relevant

	curriculumRef := &schemas.CurriculumReq{
		Focus:           focus.Description,
		Field:           field.Description,
		CurriculumType:  curriculumType.Description,
		IsActive:        true,
		StartDate:       "01.01.2025",
		EndDate:         "01.01.2028",
		FieldmanagerRef: []uint{},
		ModulesRef:      []uint{},
	}

	_, _ = db.CreateCurriculum(curriculumRef)
}

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

	generateTestdata(stuff)

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
