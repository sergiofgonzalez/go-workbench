package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// Request is a dummy struct with a label that wraps the task
// to be handled by `process`.
type Request struct {
	label string
}

// MaxOutstanding sets the max amount of tasks in the semaphore,
// that is, the number of simultaneous calls to `process`.
const MaxOutstanding = 10


func handle(queue chan *Request, id int) {
	log.Printf("handler[%d]: started and ready to process requests", id)
	for req := range queue {
		process(req, id)	// process the request (supposed to be heavy, long-running task)
	}
}

// Serve is in charge of processing a queue of requests implemented
// as a buffered channel
func Serve(queue chan *Request, quit chan bool) {
	log.Println("Serve: request server initialized")
	for i := 0; i < MaxOutstanding; i++ {
		log.Printf("Starting task handler #%d", i)
		go handle(queue, i)
	}
	<-quit		// This will block until quit signal is received
	log.Println("Serve: terminating request server")
}

// process is a dummy function that simulates the processing
// of a given request
func process(r *Request, id int) {
	delaySeconds := rand.Intn(5) + 1
	log.Printf("process[%d]: about to process %q which will take %d seconds to complete\n", id, r.label, delaySeconds)
	time.Sleep(time.Duration(delaySeconds) * time.Second)
	log.Printf("process[%d]: processing of %q complete\n", id, r.label)
}

// main initiates the processing of requests
func main() {

	reqQueue := make(chan *Request, 10)	// setting up the request queue as an unbuffered channel

	quit := make(chan bool)		// unbuffered channel for the quit signal
	go Serve(reqQueue, quit)  // initialize the Server to process the requests concurrently
	go timeout(quit)


	// creating a few dummy requests
	requests := make([]Request, 100)
	for i := 0; i < cap(requests); i++ {
		requests[i] = Request{fmt.Sprintf("req #%d", i)}
	}

	// i used for _, req := range requests but it causes req to always get the same address!
	for i := 0; i < len(requests); i++ {
		reqQueue <- &requests[i]
	}

}

func timeout(quit chan bool) {
	for i := 0; i < 3; i++ {
		log.Printf("timeout: quit signal will be sent in %d seconds", (3 - i)*10)
		time.Sleep(10 * time.Second)
	}
	log.Printf("timeout: sending signal to quit after timeout exahausted")
	quit <- true // send signal to notify timeout is done
}