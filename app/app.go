package app

import (
	"errors"
	"fmt"
	"interview/config"
	"interview/controllers"
	"interview/mw"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const cookieMaxAge = 60 * 60 * 24

var ErrInternalServerError = errors.New("internal Server Error")

type API struct {
	port       string
	router     *gin.Engine
	controller *controllers.Controller
}

func New(cfg config.AppConfig) *API {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	store := cookie.NewStore([]byte(cfg.SessionSecret))
	store.Options(sessions.Options{MaxAge: cookieMaxAge, Path: "/"})
	router.Use(sessions.Sessions("session", store))
	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(error); ok {
			mw.RenderError(c, fmt.Errorf("panic recovered: %w", err))
		} else {
			mw.RenderError(c, ErrInternalServerError)
		}
	}))
	router.Use(mw.ErrorHandler())
	router.Use(mw.UseRateLimiter())

	return &API{
		port:       cfg.Port,
		router:     router,
		controller: controllers.New(router),
	}
}

func (a *API) Start() error {
	return a.router.Run(fmt.Sprintf(":%s", a.port))
}
