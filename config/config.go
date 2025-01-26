package config

import (
	"errors"
	"fmt"
	"os"
	"reflect"
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

type AppConfig struct {
	SessionSecret string `env:"SESSION_SECRET"`
	Port          string `env:"HTTP_PORT"`
}

type DatabaseConfig struct {
	Host     string `env:"MYSQL_HOST"`
	Port     string `env:"MYSQL_PORT"`
	Database string `env:"MYSQL_DATABASE"`
	Username string `env:"MYSQL_USER"`
	Password string `env:"MYSQL_PASSWORD"`
}

func Get() *Config {
	var err error
	once.Do(func() {
		config, err = newConfig()
		if err != nil {
			panic(err)
		}
	})

	return config
}

func newConfig() (*Config, error) {
	cfg := &Config{
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

	err := cfg.validate()
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) validate() error {
	return validate(c)
}

func validate(s any) error {
	var errs error

	val := reflect.ValueOf(s).Elem()
	for i := range val.NumField() {
		field := val.Field(i)

		kind := field.Kind()
		if kind == reflect.Struct {
			err := validate(field.Addr().Interface())
			if err != nil {
				errs = errors.Join(errs, err)
			}
		}

		if field.IsZero() {
			param := val.Type().Field(i).Tag.Get("env")
			if param == "" {
				continue
			}

			errs = errors.Join(errs, fmt.Errorf("param %s is required", param))
		}
	}

	return errors.Join(errs)
}
