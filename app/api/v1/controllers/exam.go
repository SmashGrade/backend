package controllers

import (
	"fmt"

	s "github.com/SmashGrade/backend/app/api/v1/schemas"
	"github.com/labstack/echo/v4"
)

func GetExams(c echo.Context) error {
	var res []s.ExamRes
	// todelete
	fmt.Printf(`%v`, res)

	// TODO
	return nil
}

func PostExam(c echo.Context) error {
	var req s.ExamReq
	// todelete
	fmt.Printf(`%v`, req)

	// TODO
	return nil
}

func GetExam(c echo.Context) error {
	// Parameters
	id := c.Param("id")

	var res s.ExamRes

	// todelete
	fmt.Printf(`%v %v`, id, res)

	// TODO
	return nil
}

func PutExam(c echo.Context) error {
	// Parameters
	id := c.Param("id")

	var req s.ExamReq

	// todelete
	fmt.Printf(`%v %v`, id, req)

	// TODO
	return nil
}

func DeleteExam(c echo.Context) error {
	// Parameters
	id := c.Param("id")

	// todelete
	fmt.Printf(`%v`, id)

	// TODO
	return nil
}

func PostExamGradeStudent(c echo.Context) error {
	// Parameters
	id := c.Param("id")

	var req s.ExamReqStudent

	// todelete
	fmt.Printf(`%v %v`, id, req)

	// TODO
	return nil
}

func PostExamGradeTeacher(c echo.Context) error {
	// Parameters
	id := c.Param("id")

	var req s.ExamReqStudent

	// todelete
	fmt.Printf(`%v %v`, id, req)

	// TODO
	return nil
}
