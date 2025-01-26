package mw

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RenderError(c *gin.Context, err error) {
	renderPage(c, "error.html", err)
}

func RenderIndex(c *gin.Context, data any) {
	renderPage(c, "index.html", data)
}

func renderPage(c *gin.Context, page string, data ...any) {
	var pageData any
	if len(data) > 0 {
		pageData = data[0]
	}

	c.HTML(http.StatusOK, page, pageData)
}
