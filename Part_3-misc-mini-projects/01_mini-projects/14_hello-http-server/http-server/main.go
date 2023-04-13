package main

import (
	"fmt"
	// "io"
	"log"
	"net/http"
)

const port = ":8080"

// Handler functions

// Home is the http handler function for the homepage
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage")
	// Alternatively:
	// io.WriteString(w, "Homepage")
}

// Info is the http handler function for the info page
func Info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Info page")
	// Alternatively:
	// io.WriteString(w, "Info page")
}


func main() {
	log.Println("Starting HTTP server on port", port)
	log.Println("Type Ctrl-C to exit")

	// Associating handlers to paths
	http.HandleFunc("/", Home)
	http.HandleFunc("/info", Info)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("This will never be logged :p")
}