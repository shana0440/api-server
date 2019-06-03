package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config include all configuration, include database, app, or other else.
type Config struct {
	App *App
}

// App is the configuration about the app
type App struct {
	Port string
}

// LoadAndGetConfig use to load and get the configuration from .env file
func LoadAndGetConfig() (*Config, error) {
	err := LoadConfig()
	if err != nil {
		return nil, err
	}
	config := GetConfig()
	return config, nil
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
func GetConfig() *Config {
	return &Config{
		App: &App{
			Port: os.Getenv("APP_PORT"),
		},
	}
}
