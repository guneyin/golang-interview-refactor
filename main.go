package main

import (
	"interview/app"
	"interview/config"
	"log"
)

func main() {
	cfg := config.Get().App
	ice := app.New(cfg)

	log.Fatal(ice.Start(cfg.Port))
}
