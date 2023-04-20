package main

import (
	"fmt"
	"log"
	"time"
)

// Request is a wrapper for a task that will be processed by the Task Server.
// It includes a set of arguments to be given to the function that implements the
// task to be carried out, the function, and a channel in which the result will be
// communicated
type Request struct {
	label      string
	args       []int
	fn         func([]int) int
	resultChan chan int
}

// MaxOutstanding sets the max amount of tasks in the semaphore,
// that is, the number of simultaneous calls to `process`.
const MaxOutstanding = 10

func handle(queue chan *Request, id int) {
	log.Printf("handler[%d]: started and ready to process requests", id)
	for req := range queue {
		process(req, id) // process the request in the corresponding concurrent handler
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
	<-quit // This will block until quit signal is received
	log.Println("Serve: terminating request server")
}

// process is in charge of executing the task wrapped in the request
func process(r *Request, id int) {
	log.Printf("process[%d]: about to process %q", id, r.label)
	r.resultChan <- r.fn(r.args)
	log.Printf("process[%d]: processing of %q complete\n", id, r.label)
}

// main initiates the processing of requests
func main() {

	reqQueue := make(chan *Request, 10) // setting up the request queue as an unbuffered channel

	quit := make(chan bool)  // unbuffered channel for the quit signal
	go Serve(reqQueue, quit) // initialize the Server to process the requests concurrently
	go timeout(quit)

	// creating a single request
	req := Request{
		label: "sum 1, 2, 3",
		args: []int{1, 2, 3},
		fn: func(nums []int) int {
			result := 0
			for _, num := range nums {
				result += num
			}
			return result
		},
		resultChan: make(chan int),
	}

	// requests := make([]Request, 100)
	// for i := 0; i < cap(requests); i++ {
	// 	requests[i] = Request{fmt.Sprintf("req #%d", i)}
	// }

	// i used for _, req := range requests but it causes req to always get the same address!
	// for i := 0; i < len(requests); i++ {
	// 	reqQueue <- &requests[i]
	// }

	reqQueue <- &req
	fmt.Printf("The result of %q task is %d\n", req.label, <-req.resultChan)
}

func timeout(quit chan bool) {
	for i := 0; i < 3; i++ {
		log.Printf("timeout: quit signal will be sent in %d seconds", (3-i)*10)
		time.Sleep(10 * time.Second)
	}
	log.Printf("timeout: sending signal to quit after timeout exahausted")
	quit <- true // send signal to notify timeout is done
}
