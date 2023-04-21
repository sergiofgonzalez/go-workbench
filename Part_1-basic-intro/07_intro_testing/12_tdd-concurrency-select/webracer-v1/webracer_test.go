package webracer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestWebRacer(t *testing.T) {
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := WebRacer(slowURL, fastURL)

	slowServer.Close()
	fastServer.Close()

	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}
