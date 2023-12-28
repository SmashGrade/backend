package controllers

import (
	"fmt"
	"net/http"

	s "github.com/SmashGrade/backend/app/api/v1/schemas"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func GetCurriculums(c echo.Context) error {
	var res []s.CurriculumRes

	// todelete
	fmt.Printf(`%v`, res)

	// TODO
	return nil
}

func PostCurriculum(c echo.Context) error {
	var req s.CurriculumReq

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

func GetCurriculum(c echo.Context) error {
	// Parameters
	id := c.Param("id")

	var res s.CurriculumRes

	// todelete
	fmt.Printf(`%v %v`, id, res)

	// TODO
	return nil
}

func PutCurriculum(c echo.Context) error {
	// Parameters
	id := c.Param("id")

	var req s.CurriculumReq

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

func DeleteCurriculum(c echo.Context) error {
	// Parameters
	id := c.Param("id")

	// todelete
	fmt.Printf(`%v`, id)

	// TODO
	return nil
}

func GetCurriculumFilter(c echo.Context) error {
	var res s.CurriculumFilter

	// todelete
	fmt.Printf(`%v`, res)

	// TODO
	return nil
}
