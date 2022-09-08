package entry

import (
	"template-api-gateway/internal/controller"

	"github.com/labstack/echo/v4"
)

func NewHandler(router *echo.Echo) {
	router.GET("/ping", controller.Ping)
	router.GET("/rest", controller.Rest)
	router.GET("/grpc", controller.Grpc)
}
