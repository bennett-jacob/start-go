package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

// Settings represent all application settings
type Settings struct {
	PORT         int    `env:"PORT" envDefault:"8000"`
	DatabaseHost string `env:"DB_HOST"`
}

// Load loads a new settings object from the environment
func Load() (*Settings, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Could not use .env file")
	}

	settings := &Settings{}
	if err := env.Parse(settings); err != nil {
		return nil, err
	}
	return settings, nil
}
