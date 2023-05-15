package main

import (
	"fmt"
	"log"
	"net/http"
)

// Store is an interface that includes a single method Fetch
type Store interface {
	Fetch() string
}

// Server is a function that receives a store and returns a function you can use
// as a handler
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, store.Fetch())
	}
}

// Simulates the real implementation
type sportsStore struct {}

func (s *sportsStore) Fetch() string {
	return "List of items in the store retrieved from the inventory DB"
}

func main() {
	const port = ":8080"

	http.HandleFunc("/items", Server(&sportsStore{}))

	fmt.Println("Starting HTTP server on port:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}