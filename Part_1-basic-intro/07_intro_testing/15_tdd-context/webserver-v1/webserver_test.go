package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// SpyStore is a mock implementation of a Store for testing
type SpyStore struct {
	response string
	cancelled bool
}

// Fetch implementation of StubStore so that it implements the Store interface
func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response	// return the stubbed response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {
	t.Run("happy path: fetch runs through", func(t *testing.T) {
		// prepare the stubbed server value which returns a hardcoded string
		data := "Hello to Jason Isaacs"
		store := &SpyStore{response: data}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil) // Prepare a canned request that can be fed to ServeHTTP
		response := httptest.NewRecorder() // new initialized response recorder for assertion purposes

		server.ServeHTTP(response, request) // call the underlying handler with req and response
		got := response.Body.String()

		if got != data {
			t.Errorf("got %q, but wanted %q", got, data)
		}
		if store.cancelled {
			t.Errorf("store got cancelled when it shouldn't")
		}
	})

	t.Run("fetch gets cancelled", func(t *testing.T) {
		data := "Hello to Jasaon Isaacs"
		store := &SpyStore{response: data}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5 * time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if !store.cancelled {
			t.Errorf("store was not told to cancel when it should")
		}
	})
}
