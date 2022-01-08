package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"

	vegeta "github.com/tsenart/vegeta/v12/lib"
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

func test_rate_limit() {
	rate := vegeta.Rate{Freq: 20, Per: time.Second}
	duration := 60 * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    "http://localhost:8080/notes/5",
		Header: http.Header{"Authorization": {"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NDE2MzE2NzEsImlzX2FkbWluIjpmYWxzZSwidXNlcl9pZCI6MTB9.IVe-c7Mxfen2NDL7UM61VyB1hojWHMzJ897nGGWT8yw"}},
	})
	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()

	fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}
