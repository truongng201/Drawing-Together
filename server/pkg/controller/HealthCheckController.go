package controller

import (
	"runtime/debug"

	"github.com/labstack/echo/v4"
)

type HealthCheckController struct {}

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


func (controller HealthCheckController) Execute(c echo.Context) error {
	return c.JSON(200, &HealthCheck{
		Status: "Oke",
		Version: commit_hash,
	})
}