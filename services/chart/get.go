package chart

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"interview/database"
	"interview/entity"
	"interview/mw"
	"net/http"
)

func GetCartData(c *gin.Context) {
	data := map[string]interface{}{
		"Error": c.Query("error"),
	}

	sessionID, err := mw.GetSessionID(c)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusFound, "/")
		return
	}

	data["CartItems"] = getCartItemData(sessionID)

	c.HTML(http.StatusOK, "add_item_form.html", data)
}

func getCartItemData(sessionID string) (items []map[string]interface{}) {
	db := database.Get()
	var cartEntity entity.CartEntity
	result := db.Where(fmt.Sprintf("status = '%s' AND session_id = '%s'", entity.CartOpen, sessionID)).First(&cartEntity)

	if result.Error != nil {
		return
	}

	var cartItems []entity.CartItem
	result = db.Where(fmt.Sprintf("cart_id = %d", cartEntity.ID)).Find(&cartItems)
	if result.Error != nil {
		return
	}

	for _, cartItem := range cartItems {
		item := map[string]interface{}{
			"ID":       cartItem.ID,
			"Quantity": cartItem.Quantity,
			"Price":    cartItem.Price,
			"Product":  cartItem.ProductName,
		}

		items = append(items, item)
	}
	return items
}
