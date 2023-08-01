package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
  "runtime/debug"
)
  
type HealthCheck struct {
  status  string `json:"status"`
  version string `json:"version"`
}

var commit_hash = func() string {
    if info, ok := debug.ReadBuildInfo(); ok {
        for _, setting := range info.Settings {
            if setting.Key == "vcs.revision" {
                return setting.Value
            }
        }
    }
    return ""
}()

func Routes(e *echo.Echo) *echo.Echo {
	e.Use(middleware.Logger())

	e.GET("/health", func(c echo.Context) error {
		return c.String(200, &HealthCheck{
      status: "OK",
      version: commit_hash 
    })
	})

	return e
}
