package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

// Simple handler function
func handler(w http.ResponseWriter, r *http.Request) {
	responseLength, err := w.Write([]byte("Simple API Server"))
	if err != nil {
		log.Printf("Error writing response: %s", err)
		return
	}

	log.Printf("Endpoint: %s, Response status: %d. Response length: %d.\n", r.URL.Path, http.StatusOK, responseLength)
}

func main() {
	fmt.Println("Simple API Server")
	log.Printf("Starting server on http://localhost:%d.\n", port)

	// Register handler function to "/" endpoint
	http.HandleFunc("/", handler)

	// Start server on port 8080
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
