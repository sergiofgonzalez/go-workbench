package webracer

import (
	"net/http"
	"time"
)

// WebRacer sends the same HTTP GET request to the given URLs and return the one
// that responds faster. If none responds in less than 10 seconds and error is
// returned.
func WebRacer(url1, url2 string) (fasterURL string) {
	duration1 := measureResponseTime(url1)
	duration2 := measureResponseTime(url2)

	if duration1 < duration2 {
		return url1
	}

	return url2
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	_, _ = http.Get(url)
	return time.Since(start)
}
