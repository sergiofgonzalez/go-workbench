# Concurrency in Go using TDD: Select
> this example illustrates how to use `select` in concurrency scenarios in Go.

You have been asked to make a function called `WebsiteRacer` which takes two URLs and hits them in parallel with an HTTP GET request and returns the URL that returned first.

To prevent long timeout situations, if none of them returns within 10 seconds, an error should be returned.

## v0: first naive implementation

As good development practices recommend, you should:

> Make it work, make it right, make it fast

As a result, start by creating the test function: