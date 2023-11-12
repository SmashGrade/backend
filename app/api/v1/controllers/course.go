package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCourses(c echo.Context) error {
	// TODO
	return nil
}

func PostCourse(c echo.Context) error {
	// TODO
	return nil
}

func GetCourse(c echo.Context) error {
	// Parameters
	id := c.Param("id")
	return c.String(http.StatusOK, id)

	// TODO
	return nil
}

func PutCourse(c echo.Context) error {
	// Parameters
	id := c.Param("id")
	return c.String(http.StatusOK, id)

	// TODO
	return nil
}

func DeleteCourse(c echo.Context) error {
	// Parameters
	id := c.Param("id")
	return c.String(http.StatusOK, id)

	// TODO
	return nil
}

func GetCourseStudent(c echo.Context) error {
	// Parameters
	id := c.Param("id")
	return c.String(http.StatusOK, id)

	// TODO
	return nil
}

func GetCourseTeacher(c echo.Context) error {
	// Parameters
	id := c.Param("id")
	return c.String(http.StatusOK, id)

	// TODO
	return nil
}

func GetCourseFilter(c echo.Context) error {
	// TODO
	return nil
}
