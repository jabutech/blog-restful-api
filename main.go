package main

import (
	"jabutech/blog-restful-api/handler"
	"jabutech/blog-restful-api/helper"
	"jabutech/blog-restful-api/router"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env variable
	err := godotenv.Load(".env")
	// Check error with helper
	helper.FatalError(err)

	// // Database url
	// dsn := os.Getenv("DATABASE_URL")
	// // Open connection to database
	// _, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// // Check error with helper
	// helper.FatalError(err)

	// Handler
	appHealthHandler := handler.NewPingHandler()

	// Use router
	router := router.NewRouter(appHealthHandler)

	// Run router
	router.Run()

}
