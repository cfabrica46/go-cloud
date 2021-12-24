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

	portHTTP := os.Getenv("PORT")
	portHTTPS := os.Getenv("PORTHTTPS")

	err = runServer(portHTTP, portHTTPS)
	if err != nil {
		log.Fatal(err)
	}
}
