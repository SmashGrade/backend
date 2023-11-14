package controllers

import (
	"fmt"
	"net/http"

	s "github.com/SmashGrade/backend/app/api/v1/schemas"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func PostOnboarding(c echo.Context) error {
	var req s.OnboardingReq

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

func PutOnboarding(c echo.Context) error {
	var req s.OnboardingReq

	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := validator.New().Struct(req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Parameters
	id := c.Param("id")

	// todelete
	fmt.Printf(`%v %v`, id, req)

	// TODO
	return nil
}

func GetOnboardingFilter(c echo.Context) error {
	var res s.OnboardingFilter

	// todelete
	fmt.Printf(`%v`, res)

	// TODO
	return nil
}
