package middlewares

import (
	"github.com/gin-gonic/gin"
	ratelimt2 "github.com/juju/ratelimit"
	ratelimt1 "go.uber.org/ratelimit"
	"net/http"
	"time"
)

// 漏桶算法
func Ratelimit1(rate int) func(c *gin.Context) {
	rl := ratelimt1.New(rate) //生成一个限流器
	return func(c *gin.Context) {
		// 去水滴
		if rl.Take().Sub(time.Now()) > 0 {
			c.String(http.StatusOK, "reta limit ...")
			c.Abort()
			return
		}
		c.Next()
	}
}

// 令牌算法
func Ratelimit2(rate time.Duration, capital int64) func(c *gin.Context) {
	// NewBucketWithQuantum 创建指定填充速率、容量大小和每次填充的令牌数的令牌桶
	// NewBucketWithRate 创建填充速度为指定速率和容量大小的令牌桶
	rl := ratelimt2.NewBucket(rate, capital) // 填充速率和容量
	return func(c *gin.Context) {
		// Take 可以赊账
		if rl.TakeAvailable(1) == 1 { // 有就能拿
			c.Next()
			return
		}
		c.String(http.StatusOK, "reta limit ...")
		c.Abort()
	}
}
