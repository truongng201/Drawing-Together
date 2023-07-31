package main

import (
	config "server/pkg/config"
	routes "server/pkg/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadConfig()

	e := echo.New()

	e = routes.Routes(e)
	e.Logger.Info("Server is running on port 8080")
	e.Logger.Fatal(e.Start(":8080"))
}