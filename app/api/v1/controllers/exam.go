package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetExams(c echo.Context) error {
	// TODO
	return nil
}

func PostExam(c echo.Context) error {
	// TODO
	return nil
}

func GetExam(c echo.Context) error {
	// Parameters
	id := c.Param("id")
	return c.String(http.StatusOK, id)

	// TODO
	return nil
}

func PutExam(c echo.Context) error {
	// Parameters
	id := c.Param("id")
	return c.String(http.StatusOK, id)

	// TODO
	return nil
}

func DeleteExam(c echo.Context) error {
	// Parameters
	id := c.Param("id")
	return c.String(http.StatusOK, id)

	// TODO
	return nil
}

func PostExamGradeStudent(c echo.Context) error {
	// Parameters
	id := c.Param("id")
	return c.String(http.StatusOK, id)

	// TODO
	return nil
}

func PostExamGradeTeacher(c echo.Context) error {
	// Parameters
	id := c.Param("id")
	return c.String(http.StatusOK, id)

	// TODO
	return nil
}
