# Mini-project: Concurrency: timeout and interval in Go
> this example illustrates how to schedule the execution in functions in Go.


## v0: initial implementation

Create a naive implementation of `setInterval` and `setTimeout` as functions that receive as arguments a function and a delay in milliseconds and execute them after the specified time.

Confirm that they work as expected. Notice that the program can finish without having completed the goroutine. (HINT: you can use `time.Sleep` to give some time to the main goroutine to complete).

## v1: introduce channels

Introduce channels, so that you can get notified when the operation is completed. Confirm that you can get rid of the `time.Sleep` when doing so.

You will need to change the signature of the functions so that they receive a channel.