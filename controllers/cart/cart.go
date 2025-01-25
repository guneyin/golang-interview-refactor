package cart

import (
	"github.com/gin-gonic/gin"
	"interview/mw"
	"interview/services/chart"
	"sync"
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
	g := router.Group("/").Use(mw.UseSession())
	g.GET("/", h.ShowAddItemForm)
	g.POST("/add-item", h.AddItem)
	g.GET("/remove-cart-item", h.DeleteCartItem)
}

func (*Handler) ShowAddItemForm(c *gin.Context) {
	chart.GetCartData(c)
}

func (*Handler) AddItem(c *gin.Context) {
	chart.AddItemToCart(c)
}

func (*Handler) DeleteCartItem(c *gin.Context) {
	chart.DeleteCartItem(c)
}
