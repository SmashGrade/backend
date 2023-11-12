package controllers

import (
	"fmt"

	s "github.com/SmashGrade/backend/app/api/v1/schemas"
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

	// todelete
	fmt.Printf(`%v`, req)

	// TODO
	return nil
}

func GetCourse(c echo.Context) error {
	// Parameters
	id := c.Param("id")

	var res s.CourseRes

	// todelete
	fmt.Printf(`%v %v`, id, res)

	// TODO
	return nil
}

func PutCourse(c echo.Context) error {
	// Parameters
	id := c.Param("id")

	var req s.CourseReqPut

	// todelete
	fmt.Printf(`%v %v`, id, req)

	// TODO
	return nil
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
