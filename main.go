package main

import (
	"fmt"
	"jabutech/blog-restful-api/helper"
	"os"

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

	fmt.Println("ok")
}
