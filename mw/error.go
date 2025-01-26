package mw

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FatalError error

func Fatal(err error) FatalError {
	return err
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			errType := "error"
			err := c.Errors.Last().Err

			var fatalError FatalError
			if errors.As(err, &fatalError) {
				errType = "fatal"
			}

			c.Redirect(http.StatusFound, fmt.Sprintf("/?%s=%s", errType, err))
		}
	}
}
