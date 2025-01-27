package mw

import (
	"time"

	"github.com/didip/tollbooth/v7"
	"github.com/didip/tollbooth/v7/limiter"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
)

func UseRateLimiter() gin.HandlerFunc {
	rl := tollbooth.NewLimiter(1, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour})

	return tollbooth_gin.LimitHandler(rl)
}
