package main

import (
	"github.com/gin-gonic/gin"
	"interview/controllers/cart"
	"interview/database"
	"net/http"
)

func main() {
	database.Migrate()

	ginEngine := gin.Default()
	ginEngine.LoadHTMLGlob("templates/*")

	cnt := cart.NewController()
	ginEngine.GET("/", cnt.ShowAddItemForm)
	ginEngine.POST("/add-item", cnt.AddItem)
	ginEngine.GET("/remove-cart-item", cnt.DeleteCartItem)
	srv := &http.Server{
		Addr:    ":8088",
		Handler: ginEngine,
	}

	srv.ListenAndServe()
}
