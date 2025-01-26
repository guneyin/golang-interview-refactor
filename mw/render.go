package mw

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RenderIndex(c *gin.Context, data any) {
	renderPage(c, "index.html", data)
}

func RenderError(c *gin.Context, err string) {
	data := gin.H{"error": err}

	renderPage(c, "error.html", data)
}

func renderPage(c *gin.Context, page string, data ...any) {
	var pageData any
	if len(data) > 0 {
		pageData = data[0]
	}

	c.HTML(http.StatusOK, page, pageData)
}
