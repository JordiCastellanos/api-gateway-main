package controller

import (
	"net/http"
	"template-api-gateway/internal/domain/service"

	"github.com/labstack/echo/v4"
)

func Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"ping": "pong",
	})
}

func Rest(c echo.Context) error {
	serviceUser := service.NewServiceUser()
	return c.String(http.StatusOK, serviceUser.GetData())
}

func Grpc(c echo.Context) error {
	serviceUser := service.NewServiceUser()
	return c.JSON(http.StatusOK, serviceUser.GetDataGrpc())
}
