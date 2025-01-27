package app_test

import (
	"context"
	"interview/app"
	"interview/config"
	"os"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	t.Run("should return a new app", func(t *testing.T) {
		_ = os.Chdir("../")

		cfg := config.AppConfig{
			SessionSecret: "s€cR€t",
			Port:          "8088",
		}

		ice := app.New(cfg)
		if ice == nil {
			t.Fatal("app is nil")
		}

		ctx, done := context.WithTimeout(context.Background(), time.Second)
		defer done()

		var err error

		go func() {
			err = ice.Start()
		}()

		<-ctx.Done()

		if err != nil {
			t.Fatal(err)
		}
	})
}
