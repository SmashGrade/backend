package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SetupContext struct {
	Echo echo.Echo
}

func Setup() SetupContext {
	return SetupContext{
		Echo: *echo.New(),
	}
}

func (ctx *SetupContext) Run(port int) error {
	portStr := fmt.Sprintf(":%v", port)
	s := http.Server{
		Addr:    portStr,
		Handler: &ctx.Echo,
	}
	err := s.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (ctx *SetupContext) Stop() error {
	err := ctx.Stop()
	if err != nil {
		return err
	}
	return nil
}
