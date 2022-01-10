package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func updateNote(c *gin.Context) {
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
}

func deleteNote(c *gin.Context) {
	note_id, _ := strconv.Atoi(c.Param("note_id"))
	var note Note

	user_id, _ := c.Get("user_id")
	is_admin, _ := c.Get("is_admin")

	verified_note, err := getNote(note_id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Item not found"})
	} else if !(verified_note.UserID == int(user_id.(float64)) || is_admin.(bool)) {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "You can't delete someone else note"})
	} else {
		db.Delete(&note, note_id)
		c.JSON(http.StatusOK, gin.H{"Success": "Item deleted"})
	}
}

func getNoteRoute(c *gin.Context) {
	note_id, _ := strconv.Atoi(c.Param("note_id"))
	note, err := getNote(note_id)

	user_id, _ := c.Get("user_id")
	is_admin, _ := c.Get("is_admin")

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Item not found"})
	} else if note.UserID != int(user_id.(float64)) && !is_admin.(bool) {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "You can't see someone else note"})
	} else {
		c.JSON(http.StatusOK, NoteToJSON(note))
	}
}

func createNote(c *gin.Context) {
	var note Note
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"Result": "Bad json"})
		return
	}
	user_id, _ := c.Get("user_id")
	note.UserID = int(user_id.(float64))
	db.Create(&note)
	addNoteToCache(note)
	c.JSON(http.StatusOK, NoteToJSON(note))
}

func getAllNotes(c *gin.Context) {
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
}
