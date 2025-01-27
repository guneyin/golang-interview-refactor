package main

import (
	"interview/app"
	"interview/config"
	"interview/database"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	cfg := config.Get()

	err := database.InitDB(database.DBMySQL)
	if err != nil {
		log.Fatal(err)
	}

	ice := app.New(cfg.App)

	log.Fatal(ice.Start())
}
