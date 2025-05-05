package controllers

import (
	"net/http"
	"strconv"
	"sync"

	"fictional-notes-api-go/models"
	"github.com/gin-gonic/gin"
)

var (
	notes      = make(map[int]models.Note)
	nextID     = 1
	notesMutex = sync.Mutex{}
)

// @Summary Create a new note
// @Tags notes
// @Accept json
// @Produce json
// @Param note body models.Note true "Note data"
// @Success 200 {object} models.Note
// @Failure 400 {object} map[string]string
// @Security BearerAuth
// @Router /notes [post]
func CreateNote(c *gin.Context) {
	var note models.Note
	if err := c.BindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	notesMutex.Lock()
	note.ID = nextID
	notes[nextID] = note
	nextID++
	notesMutex.Unlock()

	c.JSON(http.StatusOK, note)
}

// @Summary Get all notes
// @Tags notes
// @Produce json
// @Success 200 {array} models.Note
// @Security BearerAuth
// @Router /notes [get]
func GetNotes(c *gin.Context) {
	notesMutex.Lock()
	defer notesMutex.Unlock()

	var list []models.Note
	for _, n := range notes {
		list = append(list, n)
	}
	c.JSON(http.StatusOK, list)
}

// @Summary Delete a note
// @Tags notes
// @Param id path int true "Note ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Security BearerAuth
// @Router /notes/{id} [delete]
func DeleteNote(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	notesMutex.Lock()
	defer notesMutex.Unlock()

	if _, exists := notes[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	delete(notes, id)
	c.JSON(http.StatusOK, gin.H{"message": "Note deleted"})
}
