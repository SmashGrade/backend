package api

import "github.com/labstack/echo/v4"

type Router struct {
	echo *echo.Echo
}

func NewRouter(e *echo.Echo) *Router {
	return &Router{echo: e}
}
