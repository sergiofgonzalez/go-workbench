package webracer

import (
	"fmt"
	"net/http"
	"time"
)

// WebRacer sends the same HTTP GET request to the given URLs and return the one
// that responds faster. If none responds in less than 10 seconds and error is
// returned.
func WebRacer(url1, url2 string) (fasterURL string) {
	start1 := time.Now()
	_, _ = http.Get(url1)
	duration1 := time.Since(start1)

	start2 := time.Now()
	_, _ = http.Get(url2)
	duration2 := time.Since(start2)

	if duration1 < duration2 {
		return url1
	}

	return url2
}