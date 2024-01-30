package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	s "github.com/SmashGrade/backend/app/api/v1/schemas"
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

	// TODO
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

func PostExamGradeStudent(c echo.Context) error {
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
