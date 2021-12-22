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

func connectToDB() {
	// Create
	// db.Create(&Note{Code: "D42", Price: 100})

	// // Read
	// var product Note
	// db.First(&product, 1)                 // find product with integer primary key
	// db.First(&product, "code = ?", "D42") // find product with code D42

	// // Update - update product's price to 200
	// db.Model(&product).Update("Price", 200)
	// // Update - update multiple fields
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - delete product
	// db.Delete(&product, 1)
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
		db.First(&note, note_id)
		c.JSON(http.StatusOK, gin.H{"Title": note.Title, "Body": note.Body})
	})

	r.Run(":8080")
}
