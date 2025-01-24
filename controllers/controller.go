package controllers

import (
	"github.com/gin-gonic/gin"
	"interview/controllers/cart"
)

type IHandler interface {
	SetRoutes(*gin.Engine)
}

var (
	_ IHandler = (*cart.Handler)(nil)
)

type Controller struct {
	router *gin.Engine
}

func New(router *gin.Engine) *Controller {
	cnt := &Controller{
		router: router,
	}
	cnt.registerHandlers()
	return cnt
}

func (c *Controller) registerHandlers() {
	cart.Register(c.router)
}
