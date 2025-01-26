package main

import (
	"interview/app"
	"interview/config"
	"interview/database"
	"log"
)

func main() {
	cfg := config.Get()

	err := database.InitDB(database.DBMySQL)
	if err != nil {
		log.Fatal(err)
	}

	ice := app.New(cfg.App)

	log.Fatal(ice.Start())
}
