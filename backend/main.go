package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Title string `json:"Title" binding:"required"`
	Body  string `json:"Body" binding:"required"`
}

func main() {

	r := gin.Default()
	// connect to db
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Note{})

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

	r.Run(":8080")
}
