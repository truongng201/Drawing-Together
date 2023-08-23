package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type config struct {
	// Database
	// Database struct {
	// 	POSTGRES_USER     string
	// 	POSTGRES_PASSWORD string
	// 	POSTGRES_DB       string
	// 	POSTGRES_HOST     string
	// 	POSTGRES_PORT     string
	// }
	// Server environment
	Environment string
	// Allowed origins
	AllowedOrigins []string
}

var Con config

func LoadConfig() (config config, err error) {
	log.Info("Loading config file")

	Config := &Con
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		log.Printf("Unable to decode into struct, %v", err)
	}
	return
}
