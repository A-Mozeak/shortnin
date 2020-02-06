package main

import (
	"log"
	"net/http"
)

const PORT = ":8080"

func main() {
	router := NewRouter()
	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(PORT, router))
}
