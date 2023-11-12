package controllers

import (
	"fmt"

	s "github.com/SmashGrade/backend/app/api/v1/schemas"
	"github.com/labstack/echo/v4"
)

func GetModules(c echo.Context) error {
	var res []s.ModuleRes

	// todelete
	fmt.Printf(`%v`, res)

	// TODO
	return nil
}

func PostModule(c echo.Context) error {
	var req s.ModuleReq

	// todelete
	fmt.Printf(`%v`, req)

	// TODO
	return nil
}

func GetModule(c echo.Context) error {
	// Parameters
	id := c.Param("id")

	var res s.ModuleRes

	// todelete
	fmt.Printf(`%v %v`, id, res)

	// TODO
	return nil
}

func PutModule(c echo.Context) error {
	// Parameters
	id := c.Param("id")

	var req s.ModuleReq

	// todelete
	fmt.Printf(`%v %v`, id, req)

	// TODO
	return nil
}

func DeleteModule(c echo.Context) error {
	// Parameters
	id := c.Param("id")

	// todelete
	fmt.Printf(`%v`, id)

	// TODO
	return nil
}

func GetModuleStudent(c echo.Context) error {
	// Parameters
	studyStageId := c.Param("studyStageId")

	var res []s.ModuleResStudent

	// todelete
	fmt.Printf(`%v %v`, studyStageId, res)

	// TODO
	return nil
}

func GetModuleTeacher(c echo.Context) error {
	// Parameters
	studyStageId := c.Param("studyStageId")

	var res []s.ModuleResTeacher

	// todelete
	fmt.Printf(`%v %v`, studyStageId, res)

	// TODO
	return nil
}

func GetModuleFilter(c echo.Context) error {
	var res s.ModuleFilter

	// todelete
	fmt.Printf(`%v`, res)

	// TODO
	return nil
}
