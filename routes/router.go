package routes

import (
	"fictional-notes-api-go/controllers"
	"fictional-notes-api-go/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// Public
	router.POST("/login", controllers.Login)

	// Protected group
	auth := router.Group("/")
	auth.Use(middleware.JWTAuthMiddleware(), middleware.RateLimiter())
	{
		auth.POST("/notes", controllers.CreateNote)
		auth.GET("/notes", controllers.GetNotes)
		auth.DELETE("/notes/:id", controllers.DeleteNote)
	}
}
