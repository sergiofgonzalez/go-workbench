package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const port = ":8080"

// Greet writes a personalized message in the given Writer
func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello to %s!", name)
}

// GreetHandler is an HTTP handler that adapts the Greet function
func GreetHandler(w http.ResponseWriter, req *http.Request) {
	Greet(w, "Jason Isaacs")
}

func main() {
	log.Println("Starting HTTP server on port", port)
	log.Println("Type CTRL+C to exit")

	log.Fatal(http.ListenAndServe(port, http.HandlerFunc(GreetHandler)))
}
