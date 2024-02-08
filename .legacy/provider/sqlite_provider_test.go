package provider

import (
	"fmt"
	"testing"

	"github.com/SmashGrade/backend/legacy/entity"
)

func TestProvider(t *testing.T) {
	prov := &SqliteProvider{}
	prov.Connect()
}

// Test if a course can be saved with the same id but a higher version
func TestAutoUpdate(t *testing.T) {
	prov := &SqliteProvider{}
	prov.Connect()

	courseAutoincrement := &entity.Course{Description: "Testcourse", Version: 1}
	var courseUpdate *entity.Course

	courseReturn := prov.Db.Save(courseAutoincrement)

	if courseReturn.RowsAffected != 1 {
		t.Fatalf("Nothing inserted at autoincrement")
	}

	prov.Db.Model(&entity.Course{}).Where(courseAutoincrement).First(&courseUpdate)

	if courseUpdate.Description != courseAutoincrement.Description {
		t.Fatalf("Did not find the autoincrement course")
	}

	courseAutoincrement.ID = courseUpdate.ID
	courseAutoincrement.Version = courseUpdate.Version

	fmt.Printf("Initial ID %v and Version %v\n", courseAutoincrement.ID, courseAutoincrement.Version)

	courseUpdate.Version += 1
	courseUpdate.Description = "Supercourse"

	fmt.Printf("Update ID %v and Version %v\n", courseUpdate.ID, courseUpdate.Version)

	courseReturn = prov.Db.Save(courseUpdate)

	if courseReturn.RowsAffected != 1 {
		t.Fatalf("Nothing inserted at update")
	}

	var courseUpdateResult entity.Course
	prov.Db.Model(&entity.Course{}).Where(courseUpdate).First(&courseUpdateResult)

	if courseUpdateResult.Description != courseUpdate.Description {
		t.Fatalf("Did not find the updated course")
	}

	if courseAutoincrement.ID != courseUpdateResult.ID {
		t.Fatalf("ID of initial course %v and updated course %v do not match", courseAutoincrement.ID, courseUpdateResult.ID)
	}

	if courseAutoincrement.Version >= courseUpdateResult.Version {
		t.Fatalf("Version of initial course %v and updated course %v in wrong order", courseAutoincrement.Version, courseUpdateResult.Version)
	}
}
