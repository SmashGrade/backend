package controllers

import (
	"net/http"
	"strconv"

	s "github.com/SmashGrade/backend/app/api/v1/schemas"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func GetCourses(c echo.Context) error {
	var res []s.CoursesRes

	if err := db.ListCourses(&res); err != nil {
		return c.JSON(http.StatusBadGateway, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func PostCourse(c echo.Context) error {
	var req s.CourseReqPost

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := validator.New().Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	course, err := db.CreateCourse(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, course)
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

/*
func DeleteCourse(c echo.Context) error {
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

	if err := db.DeleteCourse(uint(id), uint(version)); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusNoContent, nil)
}
*/

func GetCourseStudent(c echo.Context) error {
	var res s.CourseResStudent
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

	// TODO: where do i get the UserId
	err = db.GetCourseResStudent(&res, uint(id), uint(version), 8)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusAccepted, res)
}

func GetCourseTeacher(c echo.Context) error {
	var res s.CourseResTeacher

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

	// TODO: where do i get the UserId
	err = db.GetCourseResTeacher(&res, uint(id), uint(version), 8)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusAccepted, res)
}

func GetCourseFilter(c echo.Context) error {
	var res s.CourseFilter

	if err := db.GetCourseFilter(&res); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
