package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	s "github.com/SmashGrade/backend/app/api/v1/schemas"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func GetModules(c echo.Context) error {
	var res []s.ModuleRes

	err := db.ListModule(&res)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

// create new module
func PostModule(c echo.Context) error {
	var req s.ModuleReq

	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := validator.New().Struct(req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	module, err := db.CreateModule(&req)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, module)
}

func GetModule(c echo.Context) error {
	var res s.ModuleRes

	// Parameters
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	versionStr := c.QueryParam("version")
	var version uint64 = 0
	if versionStr != "" {
		version, err = strconv.ParseUint(versionStr, 10, 32)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
	}

	err = db.GetModule(&res, uint(id), uint(version))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func PutModule(c echo.Context) error {
	// Parameters
	id := c.Param("id")

	var req s.ModuleReq

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

func DeleteModule(c echo.Context) error {
	// Parameters
	id := c.Param("id")

	// todelete
	fmt.Printf(`%v`, id)

	// TODO
	return nil
}

func GetModuleStudent(c echo.Context) error {
	var res []s.ModuleRes

	// Parameters
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	studyStageIdStr := c.QueryParam("studyStageId")
	var studyStageId uint64 = 0
	if studyStageIdStr != "" {
		studyStageId, err = strconv.ParseUint(studyStageIdStr, 10, 32)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
	}

	// TODO: where do i get the UserId
	err = db.ListCoursesModuleStudent(&res, uint(id), uint(studyStageId))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func GetModuleTeacher(c echo.Context) error {
	var res []s.ModuleRes

	// Parameters
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	studyStageIdStr := c.QueryParam("studyStageId")
	var studyStageId uint64 = 0
	if studyStageIdStr != "" {
		studyStageId, err = strconv.ParseUint(studyStageIdStr, 10, 32)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
	}

	// TODO: where do i get the UserId
	err = db.ListModuleTeacher(&res, uint(id), uint(studyStageId))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func GetModuleFilter(c echo.Context) error {
	var res s.ModuleFilter

	err := db.GetModuleFilter(&res)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
