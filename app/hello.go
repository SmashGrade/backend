package main

import (
	"fmt"

	"github.com/SmashGrade/backend/app/api"
	v1 "github.com/SmashGrade/backend/app/api/v1"
	"gorm.io/gorm"
)

type Test struct {
	gorm.Model
}

func main() {
	ctx := api.Setup()

	// add all versions
	v1.RoutesV1(&ctx)
	// v2...
	// v3...
	// v4...

	// start server at Port :9000
	err := ctx.Run(9000)
	if err != nil {
		fmt.Println(err) // TODO: handle error
	}

	// http://localhost:9000/api/v1/apples response: "Lululu i've got some Apples"

	/*
		fmt.Println("Hello world!")

		db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		// Migrate the schema
		db.AutoMigrate(&Test{})
	*/
}

func Add(a, b int) int {
	return a + b
}
