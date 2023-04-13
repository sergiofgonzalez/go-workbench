package main

import (
	"fmt"
	"log"
	"net/http"
)

// Counter keeps track of the number of times a page has been visited
type Counter int


func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	*ctr++
	fmt.Fprintf(w, "counter = %d\n", *ctr)
}


// Chan is channel used to send notifications about the request received
type Chan chan *http.Request

func (ch Chan) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ch <- req
	fmt.Fprintf(w, "notification sent!")
}


func main() {
	const port = ":8080"
	ctr := new(Counter) // same as Counter{} but returns a pointer, which we need for http.Handle

	ch := make(Chan)
	go processNotification(ch)

	// Establishing handlers
	http.Handle("/counter", ctr)
	http.Handle("/notification", ch)

	log.Println("Starting HTTP server on", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func processNotification(ch Chan) {
	for {
		req := <- ch
		fmt.Printf("Notification received for path %q from user-agent %v\n", req.URL, req.UserAgent())
	}
}
