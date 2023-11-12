package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCurriculums(c echo.Context) error {
	// TODO
	return nil
}

func PostCurriculum(c echo.Context) error {
	// TODO
	return nil
}

func GetCurriculum(c echo.Context) error {
	// Parameters
	id := c.Param("id")
	return c.String(http.StatusOK, id)

	// TODO
	return nil
}

func PutCurriculum(c echo.Context) error {
	// Parameters
	id := c.Param("id")
	return c.String(http.StatusOK, id)

	// TODO
	return nil
}

func DeleteCurriculum(c echo.Context) error {
	// Parameters
	id := c.Param("id")
	return c.String(http.StatusOK, id)

	// TODO
	return nil
}

func GetCurriculumFilter(c echo.Context) error {
	// TODO
	return nil
}
