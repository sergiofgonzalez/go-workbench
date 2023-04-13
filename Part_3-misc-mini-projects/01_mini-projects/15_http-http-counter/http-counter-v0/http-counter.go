package main

import (
	"fmt"
	"log"
	"net/http"
)

// Counter keeps track of the number of times a page has been visited
type Counter struct {
	n int
}


func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctr.n++
	fmt.Fprintf(w, "counter = %d\n", ctr.n)
}

func main() {
	const port = ":8080"
	ctr := new(Counter) // same as Counter{} but returns a pointer, which we need for http.Handle

	http.Handle("/counter", ctr)

	log.Println("Starting HTTP server on", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}