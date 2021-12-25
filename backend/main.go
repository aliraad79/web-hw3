package main

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const BEARER_SCHEMA = "Bearer "

type M map[string]interface{}

type Note struct {
	gorm.Model
	Title  string `json:"Title" binding:"required"`
	Body   string `json:"Body" binding:"required"`
	UserID int
	User   User //`json:"owner" binding:"required"`
}

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.Contains(authHeader, BEARER_SCHEMA) || len(strings.Split(authHeader, " ")) != 2 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA):]
		claims := jwt.MapClaims{}
		tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("ACCESS_SECRET")), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}

func NoteToJSON(note Note) map[string]interface{} {
	return gin.H{"ID": note.ID, "title": note.Title, "body": note.Body, "owner": note.UserID}
}

func CreateToken(userid uint) (string, error) {
	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = 10
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		c.Next()
	}
}

func main() {

	router := gin.Default()
	router.Use(CORSMiddleware())
	note_router := router.Group("/notes")
	// note_router.Use(JWTMiddleware())

	// load env variables
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
	// connect to db
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Note{})
	db.AutoMigrate(&User{})

	note_router.GET("/", func(c *gin.Context) {
		var notes []Note
		db.Find(&notes)
		var response []M

		for _, u := range notes {
			response = append(response, M{"Body": u.Body, "Title": u.Title})
		}
		c.JSON(http.StatusOK, response)
	})

	note_router.POST("/new", func(c *gin.Context) {
		var note Note
		if err := c.ShouldBindJSON(&note); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"Result": "Bad Parameter"})
			return
		}
		db.Create(&note)
		c.JSON(http.StatusOK, NoteToJSON(note))
	})

	note_router.GET("/:note_id", func(c *gin.Context) {
		note_id := c.Param("note_id")
		var note Note
		err := db.First(&note, note_id)
		if err.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"Error": "Item not found"})
		} else {
			c.JSON(http.StatusOK, NoteToJSON(note))
		}
	})

	note_router.DELETE("/:note_id", func(c *gin.Context) {
		note_id := c.Param("note_id")
		var note Note
		err := db.Delete(&note, note_id)
		if err.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"Error": "Item not found"})
		} else {
			c.JSON(http.StatusOK, gin.H{"Success": "Item deleted"})
		}
	})

	note_router.PUT("/:note_id", func(c *gin.Context) {
		var note Note
		if err := c.ShouldBindJSON(&note); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"Result": "Bad Parameter"})
			return
		}

		note_id := c.Param("note_id")
		object := db.First(&note, note_id)

		if object.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"Error": "Item not found"})
		} else {
			if note.Title != "" {
				object.Update("Title", note.Title)
			}
			if note.Body != "" {
				object.Update("Body", note.Body)
			}
			c.JSON(http.StatusOK, NoteToJSON(note))
		}
	})

	router.POST("/login", func(c *gin.Context) {
		var u User
		if err := c.ShouldBindJSON(&u); err != nil {
			c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
			return
		}
		var user User
		e := db.Where("username = ?", u.Username).First(&user)

		if e.Error != nil || u.Password != user.Password {
			c.JSON(http.StatusUnauthorized, gin.H{"Result": "Please provide valid login details"})
			return
		}
		token, err := CreateToken(user.ID)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": BEARER_SCHEMA + token})
	})

	router.POST("/signup", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"Result": "Bad Parameter"})
			return
		}
		db.Create(&user)
		c.JSON(http.StatusOK, gin.H{"ID": user.ID})
	})

	router.Run(":8080")
}
