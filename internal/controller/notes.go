package controller

import (
	"Dairanotes/internal/business"
	"github.com/gin-gonic/gin"
)

type NotesController struct {
	bn business.NoteBusinessInterface
}

func NewNotesController(bn business.NoteBusinessInterface) *NotesController {
	return &NotesController{bn: bn}
}

func (nc *NotesController) CreateNote(c *gin.Context) {
	// TODO
}

func (nc *NotesController) GetNotes(c *gin.Context) {
	// TODO
}

func (nc *NotesController) GetNoteByID(c *gin.Context) {
	// TODO
}

func (nc *NotesController) UpdateNoteByID(c *gin.Context) {
	// TODO
}

func (nc *NotesController) DeleteNoteByID(c *gin.Context) {
	// TODO
}
