package dao_test

import (
	"testing"

	"github.com/SmashGrade/backend/legacy/api/v1/schemas"
	"github.com/SmashGrade/backend/legacy/dao"
	"github.com/SmashGrade/backend/legacy/provider"
)

// check if an exam evaluation can be created with minimal data
func TestCreateExamEvaluation(t *testing.T) {
	prov := &provider.SqliteProvider{}
	prov.Connect()
	db := dao.Database{Db: prov.Db}

	courseReq := &schemas.CourseReqPost{}

	course, err := db.CreateCourse(courseReq)
	if err != nil {
		t.Fatalf("Creation exam evaluation preparation; create course threw error %v", err.Error())
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
		t.Fatalf("Creation exam evaluation preparation; create exam threw error %v\n", err.Error())
	}

	if exm.ID < 1 {
		t.Fatalf("Creation exam evaluation preparation; create exam Impossible record id returned %v\n", exm.ID)
	}
}