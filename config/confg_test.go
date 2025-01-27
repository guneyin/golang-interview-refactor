package config_test

import (
	"interview/config"
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	_ = os.Chdir("../")

	t.Run("Get Config", func(t *testing.T) {
		cfg := config.Get()
		if cfg == nil {
			t.Fatal("config is nil")
		}
	})

	t.Run("Validate Invalid Config", func(t *testing.T) {
		cfg := config.Get()
		if cfg == nil {
			t.Fatal("config is nil")
		}

		cfg.App.Port = ""
		cfg.Database.Host = ""

		err := config.Validate(cfg)
		if err == nil {
			t.Fatal("error expected")
		}
	})
}
