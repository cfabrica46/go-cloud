package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("unread .env")
	}

	port := os.Getenv("PORT")

	r := setupRouter()
	err = r.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
