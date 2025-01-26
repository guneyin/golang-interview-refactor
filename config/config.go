package config

import (
	"os"
	"sync"

	// autoload env variable.
	_ "github.com/joho/godotenv/autoload"
)

var (
	once   sync.Once
	config *Config
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
}

func Get() *Config {
	once.Do(func() {
		config = newConfig()
	})
	return config
}

func newConfig() *Config {
	return &Config{
		App: AppConfig{
			SessionSecret: os.Getenv("SESSION_SECRET"),
			Port:          os.Getenv("HTTP_PORT"),
		},
		Database: DatabaseConfig{
			Host:     os.Getenv("MYSQL_HOST"),
			Port:     os.Getenv("MYSQL_PORT"),
			Database: os.Getenv("MYSQL_DATABASE"),
			Username: os.Getenv("MYSQL_USER"),
			Password: os.Getenv("MYSQL_PASSWORD"),
		},
	}
}

type AppConfig struct {
	SessionSecret string `env:"SESSION_SECRET"`
	Port          string `env:"HTTP_PORT" envDefault:"8088"`
}

type DatabaseConfig struct {
	Host     string `env:"MYSQL_HOST" envDefault:"localhost"`
	Port     string `env:"MYSQL_PORT" envDefault:"5432"`
	Database string `env:"MYSQL_DATABASE" envDefault:"ice_db"`
	Username string `env:"MYSQL_USER"`
	Password string `env:"MYSQL_PASSWORD"`
}
