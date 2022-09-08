package entry

import (
	"flag"
	"template-api-gateway/cmd/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var router *echo.Echo

func init() {
	var configPath = ""
	configPath = *flag.String("config", "", "")

	if configPath == "" {
		configPath = "../data/config.yml"
	}

	setConfiguration(configPath)
}

func setConfiguration(configPath string) {
	config.Setup(configPath)

}

func Run() {

	conf := config.GetConfig()
	router = echo.New()
	router.Use(middleware.Recover())
	router.Use(middleware.CORS())
	router.Use(middleware.CSRF())
	NewHandler(router)
	router.Start(":" + conf.Server.Port)
}
