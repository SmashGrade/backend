package dao_test

import (
	"testing"

	"github.com/SmashGrade/backend/app/api/v1/schemas"
	"github.com/SmashGrade/backend/app/dao"
	"github.com/SmashGrade/backend/app/provider"
)

// check if an exam can be created with minimal data
func TestCreateExam(t *testing.T) {
	prov := &provider.SqliteProvider{}
	prov.Connect()
	db := dao.Database{Db: prov.Db}

	courseReq := &schemas.CourseReqPost{}

	course, err := db.CreateCourse(courseReq)
	if err != nil {
		t.Fatalf("Creation exam preparation create course threw error %v", err.Error())
	}

	examReq := &schemas.ExamReq{
		Description:   "HelloExamn",
		CourseID:      course.ID,
		CourseVersion: course.Version,
		Weight:        2,
		Type:          "Schriftliche Arbeit",
	}

	exm, err := db.CreateExam(examReq)
	if err != nil {
		t.Fatalf("Creation exam threw error %v\n", err.Error())
	}

	if exm.ID < 1 {
		t.Fatalf("Creation exam; Impossible record id returned %v\n", exm.ID)
	}
}
