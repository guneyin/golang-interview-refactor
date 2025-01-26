package cart

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"interview/dto"
	"interview/entity"
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
	ErrInvalidBody     = errors.New("invalid body")
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

	g.GET("/", h.index)
	g.POST("/add-item", h.addItem)
	g.GET("/remove-cart-item", h.removeCartItem)
}

func (h *Handler) index(c *gin.Context) {
	data := dto.CartResponse{Error: c.Query("error")}

	cartItems, err := h.getCart(c)
	if err != nil {
		mw.RenderError(c, err)
		return
	}

	data.FromEntity(cartItems)

	mw.RenderIndex(c, data)
}

func (h *Handler) getCart(c *gin.Context) (entity.CartItems, error) {
	sessionID, err := mw.GetSessionID(c)
	if err != nil {
		return nil, err
	}

	return h.service.GetCart(sessionID)
}

func (h *Handler) addItem(c *gin.Context) {
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

	c.Redirect(http.StatusFound, "/")
}

func (h *Handler) removeCartItem(c *gin.Context) {
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

	c.Redirect(http.StatusFound, "/")
}

func (h *Handler) parseCartItemForm(c *gin.Context) (*dto.CartItemForm, error) {
	if c.Request.Body == nil {
		return nil, ErrInvalidBody
	}

	form := &dto.CartItemForm{}
	if err := binding.FormPost.Bind(c.Request, form); err != nil {
		return nil, err
	}

	return form, nil
}
