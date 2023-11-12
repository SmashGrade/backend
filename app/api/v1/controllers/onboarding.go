package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func PostOnboarding(c echo.Context) error {
	// TODO
	return nil
}

func PutOnboarding(c echo.Context) error {
	// Parameters
	id := c.Param("id")
	return c.String(http.StatusOK, id)

	// TODO
	return nil
}

func GetOnboardingFilter(c echo.Context) error {
	// TODO
	return nil
}
