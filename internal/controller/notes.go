package controller

import (
	"dairanotes/internal/auth"
	"dairanotes/internal/business"
	"dairanotes/internal/entities"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type NotesController struct {
	bn business.NoteBusinessInterface
	us business.UserBusinessInterface
}

func NewNotesController(db *sqlx.DB) *NotesController {
	methods := entities.NewNotesMethods(db)
	bn := business.NewNoteBusiness(methods)
	us := business.NewUserBusiness(entities.NewUserMethods(db))

	return &NotesController{bn: bn, us: us}
}

func (nc *NotesController) Store(c *gin.Context) {
	var newNote entities.Note

	var err error

	newNote.UserID, err = nc.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Internal server error",
		})
		return
	}

	if err = c.BindJSON(&newNote); err != nil {
		c.JSON(400, gin.H{
			"error": "Bad request",
		})
		return
	}

	err = nc.bn.Store(c.Request.Context(), newNote)
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
	userID, err := nc.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Internal server error",
		})
		return
	}

	notes, err := nc.bn.Index(c.Request.Context(), userID)
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

	userID, err := nc.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Internal server error",
		})
		return
	}

	note, err := nc.bn.Show(c.Request.Context(), noteID, userID)
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
	if err = c.BindJSON(&updatedNote); err != nil {
		c.JSON(400, gin.H{
			"error": "Bad request",
		})
		return
	}

	userID, err := nc.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Internal server error",
		})
		return
	}

	ok, err := nc.bn.CheckUserNote(c.Request.Context(), noteID, userID)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Internal server error",
		})
		return
	}

	if !ok {
		c.JSON(403, gin.H{
			"error": "Forbidden",
		})
		return
	}

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

	userID, err := nc.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Internal server error",
		})
		return
	}

	ok, err := nc.bn.CheckUserNote(c.Request.Context(), noteID, userID)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Internal server error",
		})
		return
	}

	if !ok {
		c.JSON(403, gin.H{
			"error": "Forbidden",
		})
		return
	}

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

func (nc *NotesController) GetUserIDFromContext(c *gin.Context) (int64, error) {
	rawClaims, exists := c.Get("claims")
	if !exists {
		return 0, nil
	}

	claims, ok := rawClaims.(*auth.Claims)
	if !ok {
		return 0, nil
	}

	userID, err := nc.us.GetUserID(c.Request.Context(), claims.Username)
	if err != nil {
		return 0, err
	}

	return userID, nil
}
