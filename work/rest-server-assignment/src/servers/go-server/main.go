package main

import (
	"log"
	"net/http"
)

func main() {
	handler := NewServer()
	log.Println("Go REST server listening on http://localhost:5003")
	if err := http.ListenAndServe(":5003", handler); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
