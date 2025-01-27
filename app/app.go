package app

import (
	"fmt"
	"interview/config"
	"interview/controllers"
	"interview/router"

	"github.com/gin-gonic/gin"
)

type API struct {
	port       string
	router     *gin.Engine
	controller *controllers.Controller
}

func New(cfg config.AppConfig) *API {
	r := router.NewRouter(cfg)

	return &API{
		port:       cfg.Port,
		router:     r,
		controller: controllers.New(r),
	}
}

func (a *API) Start() error {
	return a.router.Run(fmt.Sprintf(":%s", a.port))
}
