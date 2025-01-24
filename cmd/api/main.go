package main

import (
	"github.com/gin-gonic/gin"
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
	return a.router.Run(":8088")
}

func main() {
	app := newApi()
	log.Fatal(app.Start())
}
