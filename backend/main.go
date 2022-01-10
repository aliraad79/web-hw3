package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

const BEARER_SCHEMA = "Bearer "

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.JSON(http.StatusOK, "")
			return
		}
		c.Next()
	}
}

var db gorm.DB
var cacheClient CacherClient

func main() {
	// load env variables
	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file")
	}

	// rate limit redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: "",
		DB:       0,
	})
	//connect to db
	db = initDB()

	//connect to cache via gprc
	cacheClient = getCacheClient()

	router := gin.Default()
	router.Use(CORSMiddleware())
	router.POST("/login", login)
	router.POST("/signup", singup)

	note_router := router.Group("/notes")
	note_router.Use(JWTMiddleware())
	note_router.Use(RateLimiterMiddleware(redisClient))

	note_router.GET("/", getAllNotes)
	note_router.POST("/", createNote)
	note_router.GET("/:note_id", getNoteRoute)
	note_router.PUT("/:note_id", updateNote)
	note_router.DELETE("/:note_id", deleteNote)

	note_router.GET("/test", func(c *gin.Context) {
		test_rate_limit(c.GetHeader("Authorization"))
	})

	router.Run()
}
