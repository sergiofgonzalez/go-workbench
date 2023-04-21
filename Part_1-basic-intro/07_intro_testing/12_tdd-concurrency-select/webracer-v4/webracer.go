package webracer

import (
	"fmt"
	"net/http"
	"time"
)

// WebRacer sends the same HTTP GET request to the given URLs and return the one
// that responds faster. If none responds in less than 10 seconds and error is
// returned.
func WebRacer(url1, url2 string) (fasterURL string, err error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waiting for response from %q and %q", url1, url2)
	}

}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		_, _ = http.Get(url)
		close(ch)
	}()
	return ch
}
