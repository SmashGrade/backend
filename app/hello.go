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

	ctx := api.NewServer()

	// add all versions
	v1.RoutesV1(ctx)
	// v2...
	// v3...
	// v4...

	// start server at Port :9000
	err := ctx.Run(9000)
	if err != nil {
		fmt.Println(err) // TODO: handle error
	}

}
