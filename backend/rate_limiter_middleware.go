package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func RateLimiterMiddleware(redisClient *redis.Client) gin.HandlerFunc {
	_, err := redisClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprint("error init redis", err.Error()))
	}

	return func(c *gin.Context) {
		now := time.Now().UnixNano()
		user_id, _ := c.Get("user_id")
		userCntKey := fmt.Sprint(user_id)
		limit, _ := strconv.Atoi(os.Getenv("MAX_REQ_PER_MINUTE"))
		slidingWindow := time.Duration(60 * time.Second)

		redisClient.ZRemRangeByScore(userCntKey,
			"0",
			fmt.Sprint(now-(slidingWindow.Nanoseconds()))).Result()

		reqs, _ := redisClient.ZRange(userCntKey, 0, -1).Result()

		if len(reqs) >= limit {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"status":  http.StatusTooManyRequests,
				"message": "too many request",
			})
			return
		}

		c.Next()
		redisClient.ZAddNX(userCntKey, redis.Z{Score: float64(now), Member: float64(now)})
		redisClient.Expire(userCntKey, slidingWindow)
	}

}
