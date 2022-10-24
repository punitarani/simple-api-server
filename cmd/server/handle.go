package server

import (
	"net/http"
	"simple-api-server/cmd/server/routes"
)

// Handle registers all handlers to the server
func Handle() {
	http.HandleFunc("/", routes.HomeHandler)
	http.HandleFunc("/time/", routes.TimeHandler)
}
