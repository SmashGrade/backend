package v1

import (
	"net/http"

	"github.com/SmashGrade/backend/app/api"
	"github.com/labstack/echo/v4"
)

func RoutesV1(ctx *api.SetupContext) {
	v1 := ctx.Echo.Group("/api/v1")
	v1.GET("/apples", apples)
}

func apples(c echo.Context) error {
	err := c.String(http.StatusOK, "Lululu i've got some Apples")
	if err != nil {
		return err
	}
	return nil
}
