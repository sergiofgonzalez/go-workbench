package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// StubStore is a mock implementation of a Store for testing
type StubStore struct {
	response string
}

// Fetch implementation of StubStore so that it implements the Store interface
func (s *StubStore) Fetch() string {
	return s.response	// return the stubbed response
}

func TestServer(t *testing.T) {
	// prepare the stubbed server value which returns a hardcoded string
	data := "Hello to Jason Isaacs"
	server := Server(&StubStore{data})

	request := httptest.NewRequest(http.MethodGet, "/", nil) // Prepare a canned request that can be fed to ServeHTTP
	response := httptest.NewRecorder() // new initialized response recorder for assertion purposes

	server.ServeHTTP(response, request) // call the underlying handler with req and response
	got := response.Body.String()

	if got != data {
		t.Errorf("got %q, but wanted %q", got, data)
	}
}
