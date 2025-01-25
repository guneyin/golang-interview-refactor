package main

import (
	"interview/api"
	"interview/config"
	"log"
)

func main() {
	cfg := config.Get().App
	app := api.New(cfg)
	log.Fatal(app.Start(cfg.Port))
}
