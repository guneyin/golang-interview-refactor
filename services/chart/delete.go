package chart

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"interview/database"
	"interview/entity"
	"interview/mw"
	"net/http"
	"strconv"
)

func DeleteCartItem(c *gin.Context) {
	sessionID, err := mw.GetSessionID(c)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusFound, "/")
		return
	}

	cartItemIDString := c.Query("cart_item_id")
	if cartItemIDString == "" {
		c.Redirect(302, "/")
		return
	}

	db := database.Get()

	var cartEntity entity.CartEntity
	result := db.Where(fmt.Sprintf("status = '%s' AND session_id = '%s'", entity.CartOpen, sessionID)).First(&cartEntity)
	if result.Error != nil {
		c.Redirect(302, "/")
		return
	}

	if cartEntity.Status == entity.CartClosed {
		c.Redirect(302, "/")
		return
	}

	cartItemID, err := strconv.Atoi(cartItemIDString)
	if err != nil {
		c.Redirect(302, "/")
		return
	}

	var cartItemEntity entity.CartItem

	result = db.Where(" ID  = ?", cartItemID).First(&cartItemEntity)
	if result.Error != nil {
		c.Redirect(302, "/")
		return
	}

	db.Delete(&cartItemEntity)
	c.Redirect(302, "/")
}
