package mw

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			c.Redirect(302, fmt.Sprintf("/?error=%s", c.Errors.Last().Error()))
		}
	}
}
