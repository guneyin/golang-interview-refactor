package router

import (
	"errors"
	"fmt"
	"interview/config"
	"interview/mw"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const cookieMaxAge = 60 * 60 * 24

var ErrInternalServerError = errors.New("internal Server Error")

func NewRouter(cfg config.AppConfig) *gin.Engine {
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

	return router
}
