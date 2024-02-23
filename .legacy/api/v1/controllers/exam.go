package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	s "github.com/SmashGrade/backend/legacy/api/v1/schemas"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func GetExams(c echo.Context) error {
	var res []s.ExamRes

	err := db.ListExams(&res)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusAccepted, res)
}

// creates new exam post resulting entity back
func PostExam(c echo.Context) error {
	var req s.ExamReq

	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := validator.New().Struct(req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	exam, err := db.CreateExam(&req)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, exam)
}

func GetExam(c echo.Context) error {
	var res s.ExamRes

	// Parameters
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = db.GetExam(&res, uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func PutExam(c echo.Context) error {
	// Parameters
	id := c.Param("id")

	var req s.ExamReq

	// todelete
	fmt.Printf(`%v %v`, id, req)

	// TODO
	return nil
}

func DeleteExam(c echo.Context) error {
	// Parameters
	id := c.Param("id")

	// todelete
	fmt.Printf(`%v`, id)

	// TODO
	return nil
}

// converts a path param into an uint
func GetUintParam(c echo.Context, paramName string) (uint, error) {
	paramStr := c.Param(paramName)
	paramUint64, err := strconv.ParseUint(paramStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(paramUint64), nil
}

// converts a query param into an uint
func GetUintQueryParam(c echo.Context, paramName string) (uint, error) {
	paramStr := c.QueryParam(paramName)
	paramUint64, err := strconv.ParseUint(paramStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(paramUint64), nil
}

func PostExamGradeStudent(c echo.Context) error {
	// Parameters
	courseId, err := GetUintParam(c, "id")
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	courseVersion, err := GetUintQueryParam(c, "course-version")
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	startYear, err := time.Parse("2006", c.QueryParam("class-start-year"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	var req s.ExamReqStudent

	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := validator.New().Struct(req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	examEvaluation, err := db.CreateExamEvaluation(courseId, courseVersion, startYear, &req)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, examEvaluation)
}

func PostExamGradeTeacher(c echo.Context) error {
	// Parameters
	id := c.Param("id")

	var req s.ExamReqStudent

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
