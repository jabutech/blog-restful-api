package main

import (
	"jabutech/blog-restful-api/handler"
	"jabutech/blog-restful-api/helper"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Load .env variable
	err := godotenv.Load(".env")
	// Check error with helper
	helper.FatalError(err)

	// Database url
	dsn := os.Getenv("DATABASE_URL")
	// Open connection to database
	_, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// Check error with helper
	helper.FatalError(err)

	// Handler
	appHealthHandler := handler.NewAppHealthHandler()

	// Create router with gin
	router := gin.Default()
	// Router group api
	api := router.Group("/api")

	// Endpoint app-health
	api.GET("/app-health", appHealthHandler.Check)

	// Router run
	router.Run()
}
