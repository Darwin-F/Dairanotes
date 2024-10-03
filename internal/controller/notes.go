package controller

import (
	"dairanotes/internal/business"
	"dairanotes/internal/entities"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type NotesController struct {
	bn business.NoteBusinessInterface
}

func NewNotesController(db *sqlx.DB) *NotesController {
	noteMethods := entities.NewMethods(db)
	bn := business.NewNoteBusiness(noteMethods)
	return &NotesController{bn: bn}
}

func (nc *NotesController) Store(c *gin.Context) {
	var newNote entities.Note

	if err := c.BindJSON(&newNote); err != nil {
		c.JSON(400, gin.H{
			"error": "Bad request",
		})
		return
	}

	err := nc.bn.Store(c.Request.Context(), newNote)
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

func (nc *NotesController) Index(c *gin.Context) {
	notes, err := nc.bn.Index(c.Request.Context(), 1) //TODO : get user id from token
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(200, notes)
}

func (nc *NotesController) Show(c *gin.Context) {
	noteID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Bad request",
		})
		return
	}

	//TODO : get user id from token
	//TODO : check if note belongs to user

	note, err := nc.bn.Show(c.Request.Context(), noteID)
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

func (nc *NotesController) Update(c *gin.Context) {
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

	err = nc.bn.Update(c.Request.Context(), noteID, updatedNote)
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

func (nc *NotesController) Destroy(c *gin.Context) {
	noteID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Bad request",
		})
		return
	}

	//TODO : get user id from token
	//TODO : check if note belongs to user

	err = nc.bn.Destroy(c.Request.Context(), noteID)
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
