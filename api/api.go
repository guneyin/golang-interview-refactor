package api

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"interview/config"
	"interview/controllers"
)

type Api struct {
	router     *gin.Engine
	controller *controllers.Controller
}

func New(cfg config.AppConfig) *Api {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	store := cookie.NewStore([]byte(cfg.SessionSecret))
	store.Options(sessions.Options{MaxAge: 60 * 60 * 24})
	router.Use(sessions.Sessions("session", store))

	return &Api{
		router:     router,
		controller: controllers.New(router),
	}
}

func (a *Api) Start(port string) error {
	return a.router.Run(fmt.Sprintf(":%s", port))
}
