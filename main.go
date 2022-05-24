package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Load .env variable
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// Database url
	dsn := os.Getenv("DATABASE_URL")
	// Open connection to database
	_, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
}
