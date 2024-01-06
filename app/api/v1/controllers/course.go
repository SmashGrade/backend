package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	s "github.com/SmashGrade/backend/app/api/v1/schemas"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func GetCourses(c echo.Context) error {
	var res []s.CourseRes

	// todelete
	fmt.Printf(`%v`, res)

	// TODO
	return nil
}

func PostCourse(c echo.Context) error {
	var req s.CourseReqPost

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := validator.New().Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := db.PostCourse(&req, 1, 0); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, "Success")
}

func GetCourse(c echo.Context) error {
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

	var res s.CourseRes

	if err := db.GetCourse(&res, uint(id), uint(version)); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// TODO
	return c.JSON(http.StatusAccepted, res)
}

func PutCourse(c echo.Context) error {
	// Parameters
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var req s.CourseReqPut

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := validator.New().Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := db.PutCourse(&req, uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "Success")
}

func DeleteCourse(c echo.Context) error {
	// Parameters
	id := c.Param("id")

	// todelete
	fmt.Printf(`%v`, id)

	// TODO
	return nil
}

func GetCourseStudent(c echo.Context) error {
	// Parameters
	id := c.Param("id")

	var res s.CourseResStudent

	// todelete
	fmt.Printf(`%v %v`, id, res)

	// TODO
	return nil
}

func GetCourseTeacher(c echo.Context) error {
	// Parameters
	id := c.Param("id")

	var res s.CourseResTeacher

	// todelete
	fmt.Printf(`%v %v`, id, res)

	// TODO
	return nil
}

func GetCourseFilter(c echo.Context) error {
	var res s.CourseFilter

	// todelete
	fmt.Printf(`%v`, res)

	// TODO
	return nil
}
