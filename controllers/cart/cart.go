package cart

import (
	"errors"
	"github.com/gin-gonic/gin"
	"interview/services/chart"
	"net/http"
	"sync"
	"time"
)

var (
	once    sync.Once
	handler *Handler
)

type Handler struct {
}

func Register(router *gin.Engine) {
	once.Do(func() {
		handler = &Handler{}
		handler.SetRoutes(router)
	})
}

func (h *Handler) SetRoutes(router *gin.Engine) {
	router.GET("/", h.ShowAddItemForm)
	router.POST("/add-item", h.AddItem)
	router.GET("/remove-cart-item", h.DeleteCartItem)
}

func (*Handler) ShowAddItemForm(c *gin.Context) {
	_, err := c.Request.Cookie("ice_session_id")
	if errors.Is(err, http.ErrNoCookie) {
		c.SetCookie("ice_session_id", time.Now().String(), 3600, "/", "localhost", false, true)
	}

	chart.GetCartData(c)
}

func (*Handler) AddItem(c *gin.Context) {
	cookie, err := c.Request.Cookie("ice_session_id")

	if err != nil || errors.Is(err, http.ErrNoCookie) || (cookie != nil && cookie.Value == "") {
		c.Redirect(302, "/")
		return
	}

	chart.AddItemToCart(c)
}

func (*Handler) DeleteCartItem(c *gin.Context) {
	cookie, err := c.Request.Cookie("ice_session_id")

	if err != nil || errors.Is(err, http.ErrNoCookie) || (cookie != nil && cookie.Value == "") {
		c.Redirect(302, "/")
		return
	}

	chart.DeleteCartItem(c)
}
