package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/create", create)
	http.HandleFunc("/read", read)

	log.Println("ListenAndServe on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
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
