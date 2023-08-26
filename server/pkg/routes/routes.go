package routes

import (
	"net/http"
	"server/pkg/controller"
	"server/pkg/lib/socket"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func Routes(e *echo.Echo, controller controller.AppController) *echo.Echo {
	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/health", func(c echo.Context) error {
		return controller.HealthCheckController.Execute(c)
	})

	wsServer := socket.NewWsServer()
	go wsServer.Start()

	e.POST("/create-room", func(c echo.Context) error {
		return controller.CreateRoomController.Execute(c, wsServer)
	})

	e.POST("/check-room-existed", func(c echo.Context) error {
		return controller.CheckRoomExistedController.Execute(c, wsServer)
	})

	e.GET("/room/", func(c echo.Context) error {
		return controller.WsRoomController.Execute(c, wsServer)
	})

	return e
}
