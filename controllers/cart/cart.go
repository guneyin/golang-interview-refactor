package cart

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"interview/dto"
	"interview/mw"
	"interview/services/cart"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var (
	once    sync.Once
	handler *Handler
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
	sessionID, err := mw.GetSessionID(c)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusFound, "/")
		return
	}

	data, err := h.service.Get(sessionID)
	if err != nil {
		return
	}

	c.HTML(http.StatusOK, "add_item_form.html", data)
}

func (h *Handler) AddItem(c *gin.Context) {
	sessionID, err := mw.GetSessionID(c)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusFound, "/")
		return
	}

	form, err := h.parseCartItemForm(c)
	if err != nil {
		c.Redirect(302, "/?error="+err.Error())
		return
	}

	qty, err := strconv.ParseInt(form.Quantity, 10, 0)
	if err != nil {
		c.Redirect(302, "/?error=invalid quantity")
		return
	}

	err = h.service.Add(sessionID, form.Product, uint(qty))
	if err != nil {
		c.Redirect(302, "/?error="+err.Error())
		return
	}

	c.Redirect(302, "/")
}

func (h *Handler) DeleteCartItem(c *gin.Context) {
	sessionID, err := mw.GetSessionID(c)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusFound, "/")
		return
	}

	itemID, err := strconv.ParseInt(c.Query("cart_item_id"), 10, 0)
	if err != nil {
		c.Redirect(302, "/")
		return
	}

	err = h.service.Delete(sessionID, uint(itemID))
	if err != nil {
		c.Redirect(302, "/?error="+err.Error())
	}

	c.Redirect(302, "/")
}

func (h *Handler) parseCartItemForm(c *gin.Context) (*dto.CartItemForm, error) {
	if c.Request.Body == nil {
		return nil, fmt.Errorf("body cannot be nil")
	}

	form := &dto.CartItemForm{}
	if err := binding.FormPost.Bind(c.Request, form); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return form, nil
}
