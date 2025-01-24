package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"interview/config"
	"interview/controllers"
	"log"
)

type api struct {
	router     *gin.Engine
	controller *controllers.Controller
}

func newApi() *api {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	return &api{
		router:     router,
		controller: controllers.New(router),
	}
}

func (a *api) Start() error {
	cfg := config.Get().App
	return a.router.Run(fmt.Sprintf(":%s", cfg.Port))
}

func main() {
	app := newApi()
	log.Fatal(app.Start())
}
