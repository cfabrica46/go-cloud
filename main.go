package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("unread .env")
	}

	port := os.Getenv("PORT")

	http.HandleFunc("/create", create)
	http.HandleFunc("/read", read)

	log.Printf("ListenAndServe on localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func create(w http.ResponseWriter, r *http.Request) {
	ioutil.WriteFile("texto.txt", []byte("hola que tal"), 0644)
	w.Write([]byte("File Created"))
}

func read(w http.ResponseWriter, r *http.Request) {
	log.Println("Read")
	b, err := ioutil.ReadFile("texto.txt")
	if err != nil {
		w.Write([]byte("Error"))
	}

	result := fmt.Sprintf("%s\n", b)

	w.Write([]byte(result))
}
