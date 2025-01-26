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

type Controller struct{}

func New(router *gin.Engine) *Controller {
	cnt := &Controller{}
	cnt.registerHandlers(router)

	return cnt
}

func (c *Controller) registerHandlers(router *gin.Engine) {
	cart.Register(router)
}
