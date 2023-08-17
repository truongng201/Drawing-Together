package routes

import (
	"server/pkg/controller"
    "server/pkg/socket"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)



func Routes(e *echo.Echo, controller controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())

	e.GET("/health", func(c echo.Context) error {
        return controller.HealthCheckController.Execute(c)
    })

    // e.GET("/ws/messagesGuess", func(c echo.Context) error {
    //     return controller.WSMessageGuessController.Execute(c)
    // })
    pool := socket.NewPool()
	go pool.Start()

    e.GET("/ws/messagesChat", func(c echo.Context) error {
        return controller.WSMessageChatController.Execute(c, pool)
    })

    return e
}
