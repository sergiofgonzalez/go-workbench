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

## v1: improving the solution with a bug! 😛

In the previous version, tasks are admitted even when there is no capacity to process them, which leads to poor resource utilization.

Try to fix the previous solution by gating the task admission in the `Serve` function itself.

1. Confirm the problem in the task admission by creating 1000 request to process and sending them to the queue.

    You should see in the logs that tasks get to the `Serve` function but are not processed.

2. Refactor the logic of the `handle` function into `Serve`, so that the creation of the goroutine is now handled in the latter.

    Implement `Serve` as a for-loop that iterates over the entries of the queue channel using `range`. Within the loop, increase the semaphore, and start a go routine which is charge of processing the request and freeing the semaphore.

| NOTE: |
| :---- |
| This solution is buggy, and you can appreciate it by looking at the logs and seeing multiple occurrences for the processing of a given request. |

```
2023/04/13 15:10:57 process: about to process "req #9" which will take 2 seconds to complete
2023/04/13 15:10:58 process: processing of "req #5" complete
2023/04/13 15:10:58 process: processing of "req #7" complete
2023/04/13 15:10:59 process: processing of "req #9" complete
2023/04/13 15:11:01 process: processing of "req #9" complete
2023/04/13 15:11:02 process: processing of "req #9" complete
```

Can you spot where the bug is?

## v2: Fixing the bug

The previous solution was more efficient but buggy. The reason was that the variable used to loop over the queue was reused in the different goroutines. As a result, not only was buggy in the sense that you saw repetition of processed tasks, but also the output changed depending on certain conditions.

Implement a solution using either:
+ passing the request as an argument to the goroutine closure
+ forcing the creation of a new variable in each iteration

## v3: An alternative implementation

An alternative implementation for the same problem will consist in creating a `maxAmount` of handlers consuming from the channel. Also, in order to finalize the program cleanly, the `Serve` function can include a `quit` channel to receive the signal to finalize the program.

| NOTE: |
| :---- |
| This version illustrates how you can leverage channels in a different way to obtain a similar solution (no need for semaphore this time). |

## v4: a more realistic Task and process using channels on channels

With the building blocks in place, we can focus on the shape of the `Request` by creating a struct that includes:
  + an `args` field as `[]int`
  + an `f` field defined as a function that receive those `args` and return an `int`
  + a `resultChan` that will be used for demultiplexing, that is, to allow to send a response from that channel.

You will also need to update the `handle` function so that it uses the `resultChan` to provide the result of the function execution, and remove the `process` function.