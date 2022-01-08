package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
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

func main() {
	// load env variables
	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file")
	}

	router := gin.Default()
	router.Use(CORSMiddleware())
	note_router := router.Group("/notes")
	note_router.Use(JWTMiddleware())

	// rate limit redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: "",
		DB:       0,
	})

	note_router.Use(RateLimiterMiddleware(redisClient))

	//connect to db
	db, err := initDB()
	if err != nil {
		panic("failed to connect to database")
	}

	//connect to cache via gprc
	cacheClient := getCacheClient()

	note_router.GET("/", func(c *gin.Context) {
		var notes []Note
		db.Find(&notes)
		var response []M

		user_id, _ := c.Get("user_id")
		is_admin, _ := c.Get("is_admin")

		for _, u := range notes {
			if int(user_id.(float64)) == u.UserID || is_admin.(bool) {
				response = append(response, NoteToJSON(u))
			}
		}
		c.JSON(http.StatusOK, response)
	})

	note_router.POST("/", func(c *gin.Context) {
		var note Note
		if err := c.ShouldBindJSON(&note); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"Result": "Bad json"})
			return
		}
		user_id, _ := c.Get("user_id")
		note.UserID = int(user_id.(float64))
		db.Create(&note)
		addNoteToCache(cacheClient, note)
		c.JSON(http.StatusOK, NoteToJSON(note))
	})

	note_router.GET("/:note_id", func(c *gin.Context) {
		note_id, _ := strconv.Atoi(c.Param("note_id"))
		note, err := getNote(note_id, db, cacheClient)

		user_id, _ := c.Get("user_id")
		is_admin, _ := c.Get("is_admin")

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"Error": "Item not found"})
		} else if note.UserID != int(user_id.(float64)) && !is_admin.(bool) {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "You can't see someone else note"})
		} else {
			c.JSON(http.StatusOK, NoteToJSON(note))
		}
	})

	note_router.DELETE("/:note_id", func(c *gin.Context) {
		note_id, _ := strconv.Atoi(c.Param("note_id"))
		var note Note

		user_id, _ := c.Get("user_id")
		is_admin, _ := c.Get("is_admin")

		verified_note, err := getNote(note_id, db, cacheClient)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"Error": "Item not found"})
		} else if !(verified_note.UserID == int(user_id.(float64)) || is_admin.(bool)) {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "You can't delete someone else note"})
		} else {
			db.Delete(&note, note_id)
			c.JSON(http.StatusOK, gin.H{"Success": "Item deleted"})
		}
	})

	note_router.PUT("/:note_id", func(c *gin.Context) {
		var new_note Note
		var note Note
		if err := c.ShouldBindJSON(&new_note); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"Result": "Bad Parameter"})
			return
		}

		user_id, _ := c.Get("user_id")
		is_admin, _ := c.Get("is_admin")
		note_id := c.Param("note_id")
		object := db.First(&note, note_id)

		if object.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"Error": "Item not found"})
		} else if note.UserID != int(user_id.(float64)) && !is_admin.(bool) {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "You can't update someone else note"})
		} else {
			if new_note.Title != "" {
				object.Update("Title", new_note.Title)
			}
			if new_note.Body != "" {
				object.Update("Body", new_note.Body)
			}
			c.JSON(http.StatusOK, NoteToJSON(new_note))
		}
	})

	router.POST("/login", func(c *gin.Context) {
		var u User
		if err := c.ShouldBindJSON(&u); err != nil {
			c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
			return
		}
		user, err := getUser(u.Username, db, cacheClient)

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
	})

	router.POST("/signup", func(c *gin.Context) {
		var input SignUpInfo
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"Result": "Bad Parameters"})
			return
		}
		user := User{Username: input.Username, Password: input.Password, Is_admin: input.SecretPhrase == os.Getenv("SECRET_ADMIN_PHRASE")}
		db.Create(&user)
		addUserToCache(cacheClient, user)
		c.JSON(http.StatusOK, gin.H{"ID": user.ID})
	})

	router.GET("/test", func(c *gin.Context) {
		test_rate_limit()
	})

	router.Run(":8080")
}
