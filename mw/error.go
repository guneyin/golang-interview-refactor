package mw

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			c.Redirect(http.StatusFound, fmt.Sprintf("/?error=%s", err))
		}
	}
}
