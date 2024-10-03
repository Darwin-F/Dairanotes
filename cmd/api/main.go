package main

import (
	"dairanotes/internal/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	noteController := controller.NewNotesController(db)
	noteGroup := r.Group("/notes")
	noteGroup.GET("/", noteController.GetNotes)
	noteGroup.POST("/", noteController.CreateNote)
	noteGroup.GET("/:id", noteController.GetNoteByID)
	noteGroup.PATCH("/:id", noteController.UpdateNoteByID)
	noteGroup.DELETE("/:id", noteController.DeleteNoteByID)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	// Start the server on port 8080
	r.Run(":8080")
}
