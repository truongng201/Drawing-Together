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
                return setting.Value
            }
        }
    }
    return ""
}()

func Routes(e *echo.Echo) *echo.Echo {
	e.Use(middleware.Logger())

	e.GET("/health", func(c echo.Context) error {
    u := &HealthCheck{
      Status: "Ok",
      Version: commit_hash,
    }
		return c.JSON(200, u)
	})

	return e
}
