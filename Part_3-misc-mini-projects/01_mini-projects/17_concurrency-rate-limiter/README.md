# Mini-project: Rate limiter
> this example illustrates how to create a rate limiter with a buffered channel that acts as a semaphore.

This program implements a *Task Server* that is configured to process tasks in an asynchronous fashion. In order to control the degree of parallelism a task processing rate limiter is implemented using channels.

# v0: initial version

1. Create a data structure `Request` that wraps a string field `label`. This will represent the task to be processed.

2. Define a constant `MaxOutstanding` that will control the amount of tasks that can be processed in parallel. When set to 0 no task will be processed; when set to 1 tasks will be processed one-by-one, etc.

3. Define a `chan int` that will implement the semaphore, as a buffered channel size it with `MaxOutstanding`.

4. Define a function `handle(r *Request)` that will be used to handle requests.

    In the function, you will start by sending a value to the semaphore to indicate that a task is about to be processed. If the channel is full, that operation will be blocking, thus implementing the rate limiting capabilities.

    Then, you will invoke a function `process` passign the request.

    Finally, once process returns, you will consume one item from the channel to signal that there's availability to process one task.

5. Define a function `Serve` that is configured to receive a channel of requests `chan *Request`. Call this parameter queue, as that will be its function.

    In the function, start by implementing an infinite `for {}` loop.

    In the body of the loop you will get a request from the queue, and invoke `handle` with the retrieved request as a goroutine.

6. Define a function `process` as a dummy function that simulates the processing of a request. Include a random delay for each of the request received.

7. Implement the main function to orchestrate the process:

    Start by creating the request queue as `chan *Request` with buffer len 10.

    Then invoke the `Serve` function with the channel as a goroutine.

    After than, initialize a few dummy requests and send them to the channel.

8. As the go program will end as soon as main has completed its execution, implement a timeout of 30 seconds in main to make sure that you can see the execution of all the tasks.
