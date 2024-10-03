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
	noteGroup.GET("/", noteController.Index)
	noteGroup.POST("/", noteController.Store)
	noteGroup.GET("/:id", noteController.Show)
	noteGroup.PATCH("/:id", noteController.Update)
	noteGroup.DELETE("/:id", noteController.Destroy)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	// Start the server on port 8080
	r.Run(":8080")
}
