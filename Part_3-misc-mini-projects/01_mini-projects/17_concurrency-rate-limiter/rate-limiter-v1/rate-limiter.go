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
const MaxOutstanding = 5

// sem is a semaphore implemented as a buffered channel
var sem = make(chan int, MaxOutstanding)

// Serve is in charge of processing a queue of requests implemented
// as a buffered channel
func Serve(queue chan *Request) {
	log.Println("Serve: request server initialized")
	for req := range queue {
		sem <- 1  			// send signal to semaphore to indicate "about to start processing", will block if buffer is full
		go func() {			// concurrent processing
			process(req)	// process the request (supposed to be heavy, long-running task)
			<-sem					// send signal to indicate "processing done", will enable next request to run
		}()
	}
	log.Println("Serve: finished (the channel was closed)")
}

// process is a dummy function that simulates the processing
// of a given request
func process(r *Request) {
	delaySeconds := rand.Intn(5) + 1
	log.Printf("process: about to process %q which will take %d seconds to complete\n", r.label, delaySeconds)
	time.Sleep(time.Duration(delaySeconds) * time.Second)
	log.Printf("process: processing of %q complete\n", r.label)
}

// main initiates the processing of requests
func main() {
	// setting up the request queue as an unbuffered channel
	reqQueue := make(chan *Request, 10)

	// initialize the Server to process the requests as a goroutine
	go Serve(reqQueue)

	// creating a few dummy requests
	requests := make([]Request, 10)
	for i := 0; i < cap(requests); i++ {
		requests[i] = Request{fmt.Sprintf("req #%d", i)}
	}

	// i used for _, req := range requests but it causes req to always get the same address!
	for i := 0; i < len(requests); i++ {
		reqQueue <- &requests[i]
	}

	c := make(chan int)
	go timeout(c)
	<-c 									// wait for the timeout completion signal
	close(reqQueue)
}

func timeout(c chan int) {
	for i := 0; i < 3; i++ {
		log.Printf("timeout: channel will be closed in %d seconds", (3 - i)*10)
		time.Sleep(10 * time.Second)
	}
	log.Printf("timeout: closing channel")
	c <- 1 // send signal to notify timeout is done
}