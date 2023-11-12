package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	// TODO
	return nil
}

func PostUser(c echo.Context) error {
	// TODO
	return nil
}

func GetUser(c echo.Context) error {
	// Parameters
	id := c.Param("id")
	return c.String(http.StatusOK, id)

	// TODO
	return nil
}

func PutUser(c echo.Context) error {
	// Parameters
	id := c.Param("id")
	return c.String(http.StatusOK, id)

	// TODO
	return nil
}

func GetUserCourses(c echo.Context) error {
	// Parameters
	id := c.Param("id")
	return c.String(http.StatusOK, id)

	// TODO
	return nil
}

func GetUserExams(c echo.Context) error {
	// Parameters
	id := c.Param("id")
	return c.String(http.StatusOK, id)

	// TODO
	return nil
}

func GetUserExam(c echo.Context) error {
	// Parameters
	id := c.Param("id")
	return c.String(http.StatusOK, id)

	// TODO
	return nil
}
