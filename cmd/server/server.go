package server

import (
	"fmt"
	"log"
	"net/http"
)

// Main runs the server
func Main(port int) {
	log.Printf("Starting server on http://localhost:%d.\n", port)

	// Register all routes handlers to the server
	Handle()

	// Start server on port 8080
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
