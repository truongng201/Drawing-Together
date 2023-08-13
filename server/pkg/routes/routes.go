package routes

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"server/pkg/controller"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)
  
type HealthCheck struct {
  Status  string `json:"status" xml:"status"`
  Version string `json:"version" xml:"version"`
}

var commit_hash = func() string {
    if info, ok := debug.ReadBuildInfo(); ok {
        for _, setting := range info.Settings {
            if setting.Key == "vcs.revision" {
                return setting.Value[:7]
            }
        }
    }
    return ""
}()

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    // We'll need to check the origin of our connection
    // this will allow us to make requests from our React
    // development server to here.
    // For now, we'll do no checking and just allow any connection
    CheckOrigin: func(r *http.Request) bool { return true },
}

func Routes(e *echo.Echo, controller controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, &HealthCheck{
            Status: "Oke",
            Version: commit_hash,
        })
	})

    e.GET("/ws", func(c echo.Context) error {
        ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
        if err != nil {
            return err
        }

        defer ws.Close()

        fmt.Println("Client connected:", ws.RemoteAddr())

        for {
            // Read
            _, msg, err := ws.ReadMessage()
            if err != nil {
                c.Logger().Error(err)
            }
            fmt.Printf("Message Received : %s\n", msg)

            // Write
            err = ws.WriteMessage(websocket.TextMessage, []byte(msg))
            if err != nil {
                c.Logger().Error(err)
            }

	    }
    })

    return e
}
