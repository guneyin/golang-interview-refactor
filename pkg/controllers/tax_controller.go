package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"interview/pkg/calculator"
	"net/http"
	"time"
)

type Controller struct {
}

func New() *Controller {
	return &Controller{}
}

func (*Controller) ShowAddItemForm(c *gin.Context) {
	_, err := c.Request.Cookie("ice_session_id")
	if errors.Is(err, http.ErrNoCookie) {
		c.SetCookie("ice_session_id", time.Now().String(), 3600, "/", "localhost", false, true)
	}

	calculator.GetCartData(c)
}

func (*Controller) AddItem(c *gin.Context) {
	cookie, err := c.Request.Cookie("ice_session_id")

	if err != nil || errors.Is(err, http.ErrNoCookie) || (cookie != nil && cookie.Value == "") {
		c.Redirect(302, "/")
		return
	}

	calculator.AddItemToCart(c)
}

func (*Controller) DeleteCartItem(c *gin.Context) {
	cookie, err := c.Request.Cookie("ice_session_id")

	if err != nil || errors.Is(err, http.ErrNoCookie) || (cookie != nil && cookie.Value == "") {
		c.Redirect(302, "/")
		return
	}

	calculator.DeleteCartItem(c)
}
