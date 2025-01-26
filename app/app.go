package app

import (
	"errors"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"interview/config"
	"interview/controllers"
	"interview/mw"
)

var ErrInternalServerError = errors.New("internal Server Error")

type Api struct {
	port       string
	router     *gin.Engine
	controller *controllers.Controller
}

func New(cfg config.AppConfig) *Api {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	store := cookie.NewStore([]byte(cfg.SessionSecret))
	store.Options(sessions.Options{MaxAge: 60 * 60 * 24})
	router.Use(sessions.Sessions("session", store))
	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(error); ok {
			mw.RenderError(c, fmt.Errorf("panic recovered: %v", err))
		} else {
			mw.RenderError(c, ErrInternalServerError)
		}
	}))
	router.Use(mw.ErrorHandler())

	return &Api{
		port:       cfg.Port,
		router:     router,
		controller: controllers.New(router),
	}
}

func (a *Api) Start() error {
	return a.router.Run(fmt.Sprintf(":%s", a.port))
}

func recoveryHandler(c *gin.Context, err interface{}) {
	c.HTML(500, "error.tmpl", gin.H{
		"title": "Error",
		"err":   err,
	})
}
