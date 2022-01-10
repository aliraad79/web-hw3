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

type M map[string]interface{}

type SignUpInfo struct {
	Username     string `json:"username" binding:"required"`
	Password     string `json:"password" binding:"required"`
	SecretPhrase string `json:"secret_phrase"`
}

func NoteToJSON(note Note) map[string]interface{} {
	return gin.H{"ID": note.ID, "title": note.Title, "body": note.Body, "userID": note.UserID}
}

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

func singup(c *gin.Context) {
	var input SignUpInfo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"Result": "Bad Parameters"})
		return
	}
	user := User{Username: input.Username, Password: input.Password, Is_admin: input.SecretPhrase == os.Getenv("SECRET_ADMIN_PHRASE")}
	db.Create(&user)
	addUserToCache(user)
	c.JSON(http.StatusOK, gin.H{"ID": user.ID})
}

func login(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	user, err := getUser(u.Username)

	if err != nil || u.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"Result": "Please provide valid login details"})
		return
	}
	token, err := CreateToken(user.ID, user.Is_admin)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": BEARER_SCHEMA + token})
}

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

	router.GET("/test", func(c *gin.Context) {
		test_rate_limit()
	})

	note_router := router.Group("/notes")
	note_router.Use(JWTMiddleware())
	note_router.Use(RateLimiterMiddleware(redisClient))

	note_router.GET("/", getAllNotes)
	note_router.POST("/", createNote)
	note_router.GET("/:note_id", getNoteRoute)
	note_router.PUT("/:note_id", updateNote)
	note_router.DELETE("/:note_id", deleteNote)

	router.Run()
}
