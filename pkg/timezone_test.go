package pkg

import (
	"testing"
	"time"
)

// TestTimezone checks if the Timezone map has valid timezones
func TestTimezone(t *testing.T) {
	// Iterate over the map
	for tz := range Timezone {
		// Construct the Location object
		_, err := time.LoadLocation(tz)

		// Fail if the timezone is invalid
		if err != nil {
			t.Errorf("Invalid timezone: %s", tz)
		}
	}
}
