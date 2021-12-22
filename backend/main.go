package main

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Title string `json:"Title" binding:"required"`
	Body  string `json:"Body" binding:"required"`
}

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
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

func main() {

	r := gin.Default()
	// load env variables
	err := godotenv.Load()

	if err != nil {
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

	r.POST("/notes/new", func(c *gin.Context) {
		var input_json Note
		c.BindJSON(&input_json)
		var note_title = input_json.Title
		var note_body = input_json.Body

		if note_title == "" || note_body == "" {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"Result": "Bad Parameter"})
		} else {
			note := Note{Title: note_title, Body: note_body}
			db.Create(&note)
			c.JSON(http.StatusOK, gin.H{"ID": note.ID})
		}

	})

	r.GET("/notes/:note_id", func(c *gin.Context) {
		note_id := c.Param("note_id")
		var note Note
		err := db.First(&note, note_id)
		if err.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"Error": "Item not found"})
		} else {
			c.JSON(http.StatusOK, gin.H{"Title": note.Title, "Body": note.Body})
		}
	})

	r.DELETE("/notes/:note_id", func(c *gin.Context) {
		note_id := c.Param("note_id")
		var note Note
		err := db.Delete(&note, note_id)
		if err.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"Error": "Item not found"})
		} else {
			c.JSON(http.StatusOK, gin.H{"Success": "Item deleted"})
		}
	})

	r.PUT("/notes/:note_id", func(c *gin.Context) {
		note_id := c.Param("note_id")
		var input_json Note
		c.BindJSON(&input_json)
		var new_title = input_json.Title
		var new_body = input_json.Body

		var note Note
		object := db.First(&note, note_id)

		if object.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"Error": "Item not found"})
		} else {
			if new_title != "" {
				object.Update("Title", new_title)
			}
			if new_body != "" {
				object.Update("Body", new_body)
			}
			c.JSON(http.StatusOK, gin.H{"Title": note.Title, "Body": note.Body})
		}
	})

	r.POST("/login", func(c *gin.Context) {
		var u User
		if err := c.ShouldBindJSON(&u); err != nil {
			c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
			return
		}
		var user User
		e := db.Where("username = ?", u.Username).First(&user)

		if e.Error != nil || user.Password != user.Password {
			c.JSON(http.StatusUnauthorized, gin.H{"Result": "Please provide valid login details"})
			return
		}
		token, err := CreateToken(user.ID)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, err.Error())
			return
		}
		c.JSON(http.StatusOK, token)
	})

	r.POST("/signup", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"Result": "Bad Parameter"})
			return
		}
		db.Create(&user)
		c.JSON(http.StatusOK, gin.H{"ID": user.ID})
	})

	r.Run(":8080")
}
