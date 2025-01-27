package cart

import "github.com/gin-gonic/gin"

func NewHandler() *Handler {
	return newHandler()
}

func (h *Handler) Index(c *gin.Context) {
	h.index(c)
}

func (h *Handler) AddItem(c *gin.Context) {
	h.addItem(c)
}

func (h *Handler) RemoveCartItem(c *gin.Context) {
	h.removeCartItem(c)
}
