package webracer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestWebRacer(t *testing.T) {
	t.Run("happy path, before timeout", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := WebRacer(slowURL, fastURL)

		if got != want {
			t.Errorf("want %q, got %q", want, got)
		}
		if err != nil {
			t.Errorf("expected no errors, but got %v", err)
		}
	})

	t.Run("timeout in both servers", func(t *testing.T) {
		server1 := makeDelayedServer(25 * time.Millisecond)
		server2 := makeDelayedServer(25 * time.Millisecond)

		defer server1.Close()
		defer server2.Close()

		_, err := ConfigurableWebRacer(server1.URL, server2.URL, 20 * time.Millisecond)

		if err == nil {
			t.Errorf("expected an error, but found none")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
