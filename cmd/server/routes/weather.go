package routes

import (
	"log"
	"net/http"
	"path/filepath"
)

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	// Remove the "/weather" prefix
	path := r.URL.Path[len("/weather/"):]

	// Clean the path and use forward slashes
	path = filepath.ToSlash(filepath.Clean(path))

	// Handle the path
	if path == "." {
		// Write the usage notes
		if _, err := w.Write(writeUsage()); err != nil {
			log.Default().Println(err)
		}
	} else {
		// Redirect "/weather/{latitude}/{/longitude}" to the "/weather" path
		http.Redirect(w, r, "/weather", http.StatusTemporaryRedirect)
	}
}

func writeUsage() []byte {
	return []byte(
		"Usage: weather/{latitude}/{longitude}\n" +
			"Example: weather/40.7128/-74.0060")
}
