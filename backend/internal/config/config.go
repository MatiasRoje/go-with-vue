package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	// Server
	AppHost string
	AppPort string

	// database
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string

	// JWT
	JwtSecret string
}

func LoadConfig() (*Config, error) {
	c := &Config{}
	c.AppHost = os.Getenv("APP_HOST")
	c.AppPort = os.Getenv("APP_PORT")
	c.DbHost = os.Getenv("DB_HOST")
	c.DbPort = os.Getenv("DB_PORT")
	c.DbUser = os.Getenv("DB_USER")
	c.DbPassword = os.Getenv("DB_PASSWORD")
	c.DbName = os.Getenv("DB_NAME")
	c.JwtSecret = os.Getenv("JWT_SECRET")

	if err := validateConfig(); err != nil {
		return nil, fmt.Errorf("error validating config: %s", err)
	}

	return c, nil
}

func validateConfig() error {
	requiredVars := []string{
		"APP_HOST",
		"APP_PORT",
		"DB_HOST",
		"DB_PORT",
		"DB_USER",
		"DB_PASSWORD",
		"DB_NAME",
		"JWT_SECRET",
	}
	var errors []string
	for _, varName := range requiredVars {
		if value := os.Getenv(varName); value == "" {
			errors = append(errors, varName)
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("missing required environment variables: %s", strings.Join(errors, ", "))
	}
	return nil
}
