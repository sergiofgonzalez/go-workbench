package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
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

// ArgServer is a method on HandlerFunc that displays the arguments used to start the server
func ArgServer(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, os.Args)
}

func main() {
	const port = ":8080"
	ctr := new(Counter) // same as Counter{} but returns a pointer, which we need for http.Handle

	ch := make(Chan)
	go processNotification(ch)

	// Establishing handlers
	http.Handle("/counter", ctr)
	http.Handle("/notification", ch)
	http.Handle("/args", http.HandlerFunc(ArgServer)) // same as http.HandleFunc("/args", ArgServer)

	printArgs()

	log.Println("Starting HTTP server on", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func processNotification(ch Chan) {
	for {
		req := <-ch
		fmt.Printf("Notification received for path %q from user-agent %v\n", req.URL, req.UserAgent())
	}
}

func printArgs() {
	fmt.Println(os.Args)
}
