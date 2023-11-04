package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Test struct {
	gorm.Model
}

func main() {
	fmt.Println("Hello world!2")

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Test{})
}

func Add(a, b int) int {
	return a + b
}
