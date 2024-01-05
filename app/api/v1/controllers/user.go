package controllers

import (
	"fmt"
	"net/http"

	s "github.com/SmashGrade/backend/app/api/v1/schemas"
	"github.com/SmashGrade/backend/app/dao"
	"github.com/SmashGrade/backend/app/provider"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func getDb() *gorm.DB {
	prov := &provider.SqliteProvider{}
	prov.Connect()
	return prov.Db
}

var db = &dao.Database{Db: getDb()}

func GetUsers(c echo.Context) error {
	var res []s.User

	err := db.ListUsers(&res)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusAccepted, res)
}

func PostUser(c echo.Context) error {
	var req s.User

	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := validator.New().Struct(req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// todelete
	fmt.Printf(`%v`, req)

	// TODO
	return nil
}

func GetUser(c echo.Context) error {
	// Parameters
	id := c.Param("id")

	var res s.User

	// todelete
	fmt.Printf(`%v %v`, id, res)

	// TODO
	return nil
}

func PutUser(c echo.Context) error {
	// Parameters
	id := c.Param("id")

	var req s.User

	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := validator.New().Struct(req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// todelete
	fmt.Printf(`%v %v`, id, req)

	// TODO
	return nil
}

func GetUserCourses(c echo.Context) error {
	// Parameters
	id := c.Param("id")

	var res []s.CourseRes

	// todelete
	fmt.Printf(`%v %v`, id, res)

	// TODO
	return nil
}

func GetUserExams(c echo.Context) error {
	// Parameters
	id := c.Param("id")

	var res []s.ExamRes

	// todelete
	fmt.Printf(`%v %v`, id, res)

	// TODO
	return nil
}

func GetUserExam(c echo.Context) error {
	// Parameters
	id := c.Param("id")
	examId := c.Param("examId")

	var res s.ExamRes

	// todelete
	fmt.Printf(`%v %v %v`, id, res, examId)

	// TODO
	return nil
}
