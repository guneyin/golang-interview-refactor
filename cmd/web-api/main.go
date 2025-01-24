package main

import (
	"github.com/gin-gonic/gin"
	"interview/pkg/controllers"
	"interview/pkg/db"
	"net/http"
)

func main() {
	db.MigrateDatabase()

	ginEngine := gin.Default()
	ginEngine.LoadHTMLGlob("templates/*")

	cnt := controllers.New()
	ginEngine.GET("/", cnt.ShowAddItemForm)
	ginEngine.POST("/add-item", cnt.AddItem)
	ginEngine.GET("/remove-cart-item", cnt.DeleteCartItem)
	srv := &http.Server{
		Addr:    ":8088",
		Handler: ginEngine,
	}

	srv.ListenAndServe()
}
