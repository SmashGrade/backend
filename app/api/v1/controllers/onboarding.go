package controllers

import (
	"fmt"

	s "github.com/SmashGrade/backend/app/api/v1/schemas"
	"github.com/labstack/echo/v4"
)

func PostOnboarding(c echo.Context) error {
	var req s.OnboardingReq

	// todelete
	fmt.Printf(`%v`, req)

	// TODO
	return nil
}

func PutOnboarding(c echo.Context) error {
	var req s.OnboardingReq

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
