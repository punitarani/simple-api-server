package routes

import (
	"log"
	"net/http"
)

// HomeHandler handles the home ("" or "/") route
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	responseLength, err := w.Write([]byte("Simple API Server"))

	if err != nil {
		log.Printf("Error writing response: %s", err)
		return
	}

	log.Printf("Endpoint: %s, Response status: %d. Response length: %d.\n", r.URL.Path, http.StatusOK, responseLength)
}
