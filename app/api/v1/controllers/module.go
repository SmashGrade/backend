package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetModules(c echo.Context) error {
	// TODO
	return nil
}

func PostModule(c echo.Context) error {
	// TODO
	return nil
}

func GetModule(c echo.Context) error {
	// Parameters
	id := c.Param("id")
	return c.String(http.StatusOK, id)

	// TODO
	return nil
}

func PutModule(c echo.Context) error {
	// Parameters
	id := c.Param("id")
	return c.String(http.StatusOK, id)

	// TODO
	return nil
}

func DeleteModule(c echo.Context) error {
	// Parameters
	id := c.Param("id")
	return c.String(http.StatusOK, id)

	// TODO
	return nil
}

func GetModuleStudent(c echo.Context) error {
	// Parameters
	studyStageId := c.Param("studyStageId")
	return c.String(http.StatusOK, studyStageId)

	// TODO
	return nil
}

func GetModuleTeacher(c echo.Context) error {
	// Parameters
	studyStageId := c.Param("studyStageId")
	return c.String(http.StatusOK, studyStageId)

	// TODO
	return nil
}

func GetModuleFilter(c echo.Context) error {
	// TODO
	return nil
}
