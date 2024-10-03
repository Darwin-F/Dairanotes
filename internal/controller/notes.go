package controller

import (
	"Dairanotes/internal/business"
	"Dairanotes/internal/entities"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type NotesController struct {
	bn business.NoteBusinessInterface
}

func NewNotesController(db *sqlx.DB) *NotesController {
	noteMethods := entities.NewNoteMethods(db)
	bn := business.NewNoteBusiness(noteMethods)
	return &NotesController{bn: bn}
}

func (nc *NotesController) CreateNote(c *gin.Context) {
	var newNote entities.Note

	if err := c.BindJSON(&newNote); err != nil {
		c.JSON(400, gin.H{
			"error": "Bad request",
		})
		return
	}

	err := nc.bn.CreateNote(c.Request.Context(), newNote)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "Note created",
	})

	return
}

func (nc *NotesController) GetNotes(c *gin.Context) {
	notes, err := nc.bn.GetNotes(c.Request.Context(), 1) //TODO : get user id from token
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(200, notes)
}

func (nc *NotesController) GetNoteByID(c *gin.Context) {
	noteID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Bad request",
		})
		return
	}

	//TODO : get user id from token
	//TODO : check if note belongs to user

	note, err := nc.bn.GetNoteByID(c.Request.Context(), noteID)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Internal server error",
		})
		return
	}

	if note == nil {
		c.JSON(404, gin.H{
			"error": "Note not found",
		})
		return
	}

	c.JSON(200, note)

	return

}

func (nc *NotesController) UpdateNoteByID(c *gin.Context) {
	noteID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Bad request",
		})
		return
	}

	var updatedNote entities.Note
	if err := c.BindJSON(&updatedNote); err != nil {
		c.JSON(400, gin.H{
			"error": "Bad request",
		})
		return
	}

	//TODO : get user id from token
	//TODO : check if note belongs to user

	err = nc.bn.UpdateNoteByID(c.Request.Context(), noteID, updatedNote)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Note updated",
	})

	return
}

func (nc *NotesController) DeleteNoteByID(c *gin.Context) {
	noteID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Bad request",
		})
		return
	}

	//TODO : get user id from token
	//TODO : check if note belongs to user

	err = nc.bn.DeleteNoteByID(c.Request.Context(), noteID)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Note deleted",
	})

	return
}
