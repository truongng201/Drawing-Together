package main

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	config "server/pkg/config"
	controller "server/pkg/controller"
	routes "server/pkg/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)

func main() {
	config.LoadConfig()

	e := echo.New()

	controller := controller.AppController{}

	e = routes.Routes(e, controller)

	switch config.Con.Environment {
	case "development":
		log.SetFormatter(&CustomTextFormatter{})
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format:           "timestamp=${time_rfc3339_nano} method=${method} remote_ip=${remote_ip} uri=${uri} status=${status} error=${error}\n",
			CustomTimeFormat: "2006-01-02 15:04:05",
		})) // timestamp=2023-08-23T15:49:32.565481601Z method=GET remote_ip=172.19.0.1 uri=/health status=200 error=
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: config.Con.AllowedOrigins,
		}))

		e.Logger.Info("Server is running on port 8080")
		e.Logger.Fatal(e.Start(":8080"))

	case "production":
		CustomLogConfig()
		e.Use(middleware.Logger())
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: config.Con.AllowedOrigins,
		}))

		e.Logger.Info("Server is running on port 8080")
		e.Logger.Fatal(e.Start(":8080"))
	default:
		CustomLogConfig()
		e.Use(middleware.Logger())
		e.Logger.Fatal("Environment not set")
		return
	}
}

type CustomTextFormatter struct{}

func (f *CustomTextFormatter) Format(entry *log.Entry) ([]byte, error) {
	// Get the file and line number where the log was called
	_, filename, line, _ := runtime.Caller(7)

	// Get the script name from the full file path
	scriptName := filepath.Base(filename)

	// Format the log message
	message := fmt.Sprintf("[%s] [%s] [%s:%d] %s\n",
		entry.Time.Format("2006-01-02 15:04:05"), // Date-time
		entry.Level.String(),                     // Log level
		scriptName,                               // Script name
		line,                                     // Line number
		entry.Message,                            // Log message
	)

	return []byte(message), nil // [2023-08-23 14:58:28] [info] [room.go:66] Room ID: pgb6w37rx1
}

// Logrus configuration for JSON format
func CustomLogConfig() {
	log.SetReportCaller(true)
	formatter := &log.JSONFormatter{
		FieldMap: log.FieldMap{
			log.FieldKeyTime:  "timestamp",
			log.FieldKeyLevel: "level",
			log.FieldKeyMsg:   "message",
		},

		TimestampFormat: "02-01-2006 15:04:05", // the "time" field configuratiom
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			pathname := strings.Split(f.File, "/")
			return "", fmt.Sprintf("%s:%d", pathname[len(pathname)-1], f.Line)
		},
	}
	log.SetFormatter(formatter) // {"file":"exported.go:109","level":"info","message":"Room ID: 57zzb3z13g","timestamp":"23-08-2023 15:01:41"}
}
