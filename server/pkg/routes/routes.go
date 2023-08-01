package routes

import (
	"runtime/debug"
    "html/template"
    "io"
    
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

type TemplateRenderer struct {
	templates *template.Template
}
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func Routes(e *echo.Echo) *echo.Echo {
	e.Use(middleware.Logger())

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, &HealthCheck{
            Status: "Ok",
            Version: commit_hash,
        })
	})
    renderer := &TemplateRenderer{
        templates: template.Must(template.ParseGlob("../../../client/out/*.html")),
    }
	e.Renderer = renderer
	e.GET("/", func(c echo.Context) error {
        return c.Render(200, "../../../client/out/index.html", "")
    })
	return e
}
