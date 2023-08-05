package routes

import (
	"runtime/debug"

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


func Routes(e *echo.Echo) *echo.Echo {
	e.Use(middleware.Logger())

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, &HealthCheck{
            Status: "Oke",
            Version: commit_hash,
        })
	})

    return e
}
