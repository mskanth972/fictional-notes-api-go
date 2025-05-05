package middleware

import (
	"github.com/gin-gonic/gin"
	ginlimiter "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

func RateLimiter() gin.HandlerFunc {
	rate, err := limiter.NewRateFromFormatted("5-M") // 5 requests per minute
	if err != nil {
		panic(err)
	}
	store := memory.NewStore()
	middleware := ginlimiter.NewMiddleware(limiter.New(store, rate))
	return middleware
}
