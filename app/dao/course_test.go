package dao_test

import (
	"testing"

	"github.com/SmashGrade/backend/app/api/v1/schemas"
	"github.com/SmashGrade/backend/app/dao"
	"github.com/SmashGrade/backend/app/provider"
)

func TestPostCourse(t *testing.T) {
	prov := &provider.SqliteProvider{}
	prov.Connect()
	db := dao.Database{Db: prov.Db}

	// Post Module
	var moduleReq schemas.ModuleReq
	moduleReq.Description = "Modulebeschreibung"
	moduleReq.IsActiv = true
	moduleReq.Number = "Test123"
	db.PostModule(&moduleReq)

	// Post User
	var userReq schemas.User
	userReq.Name = "Max Muster"
	userReq.Email = "Muster@hftm.ch"
	db.PostUser(&userReq)

	var courseReq schemas.CourseReqPost
	courseReq.Description = "Neuer Kurs"
	courseReq.Number = "Course123"
	courseReq.TeacherRef = []uint{1}
	courseReq.ModuleRef = []schemas.ModuleRef{
		{Id: 2, Version: 1},
		{Id: 3, Version: 1},
	}
	courseReq.Exams = []schemas.ExamRes{{Description: "Projektarbeit", Weight: 3, Type: "Schriftliche Arbeit"}}
	db.PostCourse(&courseReq, 1, 0)
}

func TestGetCourse(t *testing.T) {
	prov := &provider.SqliteProvider{}
	prov.Connect()
	db := dao.Database{Db: prov.Db}

	var courseRes schemas.CourseRes
	db.GetCourse(&courseRes, 4, 0)
}

func TestGetCourses(t *testing.T) {
	prov := &provider.SqliteProvider{}
	prov.Connect()
	db := dao.Database{Db: prov.Db}

	var courseRes []schemas.CoursesRes
	db.ListCourses(&courseRes)
}

func TestPutCourse(t *testing.T) {
	prov := &provider.SqliteProvider{}
	prov.Connect()
	db := dao.Database{Db: prov.Db}

	var courseReq schemas.CourseReqPut
	courseReq.Description = "Angepasster Kurs 2"
	courseReq.Number = "Course123_2"
	courseReq.Version = 1
	courseReq.TeacherRef = []uint{1}
	courseReq.ModuleRef = []schemas.ModuleRef{
		{Id: 2, Version: 1},
		{Id: 3, Version: 1},
	}

	db.PutCourse(&courseReq, 2)
}

func TestFilterCourse(t *testing.T) {
	prov := &provider.SqliteProvider{}
	prov.Connect()
	db := dao.Database{Db: prov.Db}

	var courseFilter schemas.CourseFilter
	db.GetCourseFilter(&courseFilter)
}
