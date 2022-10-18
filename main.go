package main

import (
	"fmt"
	"net/http"
)

const port = 8080

func main() {
	fmt.Println("Simple API Server")
	fmt.Printf("Running on http://localhost:%d.\n", port)

	// Start server on port 8080
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Printf("Failed to start server on http://localhost:%d.\n", port)
		fmt.Printf("Error: %s\n", err)
		return
	}
}
