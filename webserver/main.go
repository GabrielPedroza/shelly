package main

import (
	"log"
	"net/http"
)

func main() {
	// Serve static files from the "dist" directory
	fs := http.FileServer(http.Dir("../../frontend/build"))
	http.Handle("/", fs)

	log.Println("Listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}