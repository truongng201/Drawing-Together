package main

import (
	config "server/pkg/config"
	controller "server/pkg/controller"
	database "server/pkg/lib/database"
	routes "server/pkg/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadConfig()
	db := database.DB.Connect()

	e := echo.New()
	
	controller := controller.AppController{}
	controller.Database = db

	e = routes.Routes(e, controller)
	
	e.Logger.Info("Server is running on port 8080")
	e.Logger.Fatal(e.Start(":8080"))
}