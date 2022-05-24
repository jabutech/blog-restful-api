package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env variable
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	appName := os.Getenv("APP_NAME")
	fmt.Println(appName)
}
