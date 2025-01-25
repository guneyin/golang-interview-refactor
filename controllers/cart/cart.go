package cart

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"interview/dto"
	"interview/mw"
	"interview/services/cart"
	"net/http"
	"strconv"
	"sync"
)

var (
	once    sync.Once
	handler *Handler

	ErrInvalidQuantity = errors.New("invalid quantity")
	ErrInvalidItemId   = errors.New("invalid item id")
)

type Handler struct {
	service *cart.Service
}

func Register(router *gin.Engine) {
	once.Do(func() {
		handler = &Handler{service: cart.NewService()}
		handler.SetRoutes(router)
	})
}

func (h *Handler) SetRoutes(router *gin.Engine) {
	g := router.Group("/").Use(mw.UseSession())
	g.GET("/", h.ShowAddItemForm)
	g.POST("/add-item", h.AddItem)
	g.GET("/remove-cart-item", h.DeleteCartItem)
}

func (h *Handler) ShowAddItemForm(c *gin.Context) {
	if c.Query("error") != "" {
		c.HTML(http.StatusOK, "add_item_form.html", map[string]any{"Error": c.Query("error")})
		return
	}

	sessionID, err := mw.GetSessionID(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	data, err := h.service.Get(sessionID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.HTML(http.StatusOK, "add_item_form.html", data)
}

func (h *Handler) AddItem(c *gin.Context) {
	sessionID, err := mw.GetSessionID(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	form, err := h.parseCartItemForm(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	qty, err := strconv.ParseInt(form.Quantity, 10, 0)
	if err != nil {
		_ = c.Error(ErrInvalidQuantity)
		return
	}

	if qty < 1 {
		_ = c.Error(ErrInvalidQuantity)
		return
	}

	err = h.service.Add(sessionID, form.Product, uint(qty))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.Redirect(302, "/")
}

func (h *Handler) DeleteCartItem(c *gin.Context) {
	sessionID, err := mw.GetSessionID(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	itemID, err := strconv.ParseInt(c.Query("cart_item_id"), 10, 0)
	if err != nil {
		_ = c.Error(ErrInvalidItemId)
		return
	}

	err = h.service.Delete(sessionID, uint(itemID))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.Redirect(302, "/")
}

func (h *Handler) parseCartItemForm(c *gin.Context) (*dto.CartItemForm, error) {
	if c.Request.Body == nil {
		return nil, fmt.Errorf("body cannot be nil")
	}

	form := &dto.CartItemForm{}
	if err := binding.FormPost.Bind(c.Request, form); err != nil {
		return nil, err
	}

	return form, nil
}
