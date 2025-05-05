package main

import (
	"fictional-notes-api-go/routes"
	"github.com/gin-gonic/gin"
	_ "fictional-notes-api-go/docs" // swaggo docs

	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

// @title Fictional Notes API
// @version 1.0
// @description This is a containerized CRUD API with login and rate limiting.
// @host localhost:8000
// @BasePath /
// @schemes http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	r := gin.Default()

	// Register API routes
	routes.RegisterRoutes(r)

	// Swagger docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8000")
}
