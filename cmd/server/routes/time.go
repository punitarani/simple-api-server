package routes

import (
	"net/http"
	"path/filepath"
	"simple-api-server/pkg"
	"time"
)

// TimeHandler handles the "/time/" path
func TimeHandler(w http.ResponseWriter, r *http.Request) {
	// Remove the "/time" prefix
	path := r.URL.Path[len("/time/"):]

	// Clean the path and use forward slashes
	path = filepath.ToSlash(filepath.Clean(path))

	// Handle the path
	if path == "." {
		// Handle the "/time" path
		timePage(w)
	} else if path == "now" {
		// Redirect "/time/now" to the "/time" path
		http.Redirect(w, r, "/time", http.StatusFound)
	} else {
		// Handle the "/time/{zone}" path
		timeZonePage(w, path)
	}
}

// timePage handles the "/time" path
func timePage(w http.ResponseWriter) {
	// Get the server time
	timeStr := "Server time: " + time.Now().Format("15:04:05")

	// Print th server time
	if _, err := w.Write([]byte(timeStr)); err != nil {
		return
	}

	// Print the available timezones
	if _, err := w.Write(getTimezones()); err != nil {
		return
	}
}

// timeZonePage handles the "/time/{zone}" path
func timeZonePage(w http.ResponseWriter, tz string) {
	// Construct the Location object
	location, err := time.LoadLocation(tz)

	// Handle invalid timezone
	if err != nil {
		timeStr := "Invalid timezone: " + tz
		if _, err := w.Write([]byte(timeStr)); err != nil {
			return
		}

		// Print the timezones
		_, _ = w.Write(getTimezones())
		return
	}

	// Get the time in the location
	tzTime := time.Now().In(location)

	// Print the time in the location
	timeStr := "Time in " + tz + ": " + tzTime.Format("15:04:05")
	if _, err = w.Write([]byte(timeStr)); err != nil {
		return
	}
}

// getTimezones returns the available timezones as a byte array
func getTimezones() []byte {
	// Create a byte slice
	var tzs []byte = make([]byte, 0)

	tzs = append(tzs, []byte("\n\nAvailable Timezones:\n--------------------\n")...)

	// Iterate through the TimezonesList
	for _, tz := range pkg.TimezonesList {
		tzs = append(tzs, []byte(tz+"\n")...)
	}

	return tzs
}
