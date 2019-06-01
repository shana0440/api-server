package config

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

// Config include all configuration, include database, app, or other else.
type Config struct {
	App *App
}

// App is the configuration about the app
type App struct {
	Port     string
	LogLevel log.Level
}

// LoadAndGetConfig use to load and get the configuration from .env file
func LoadAndGetConfig() (*Config, error) {
	err := LoadConfig()
	if err != nil {
		return nil, err
	}
	config, err := GetConfig()
	return config, err
}

// LoadConfig use to load the .env file
func LoadConfig() error {
	err := godotenv.Load()
	if err != nil {
		// TODO: should use custom logger
		log.Fatal("Error loading .env file")
	}
	return nil
}

// GetConfig use to get the config from the environment
func GetConfig() (*Config, error) {
	level, err := parseLogLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		return nil, err
	}
	return &Config{
		App: &App{
			Port:     os.Getenv("APP_PORT"),
			LogLevel: level,
		},
	}, nil
}
