# Concurrency in Go using TDD: Select
> this example illustrates how to use `select` in concurrency scenarios in Go.

You have been asked to make a function called `WebsiteRacer` which takes two URLs and hits them in parallel with an HTTP GET request and returns the URL that returned first.

To prevent long timeout situations, if none of them returns within 10 seconds, an error should be returned.

## v0: first naive implementation

As good development practices recommend, you should:

> Make it work, make it right, make it fast

As a result, start by creating the test function, in accordance to the TDD guidelines.

This will mean creating a `TestWebRacer` function that invokes a function:

```go
func WebRacer(url1, url2 string) string
```

Then, you can make the first implementation of the `WebRacer` function in which you start by taking the time using `time.Now`, then invoke the first URL using `http.Get` and ignoring the returned values; and finally calculation the duration using `time.Since`.

You then should do the same thing for the second URL, to finally compare both durations and returning the URL that took less time.

Run the test and don't bother too much if the test fails as this won't be the final implementation.

## v1: mocking the servers

The previous version had multiple problems, one being that we will actually be contacting actual sites, which is not a good idea.

There's a package in the standard library `net/http/httptest` that lets you create mock servers very easily.

In this version, change the implementation of the `TestWebRacer` function using `httptest.NewServer` function.

In the same way that you can use `http.ListenAndServe` to set up a real http server, you can use `http.NewServer` to set up a mock server.

That functions require you to pass an `http.Handler` which is the well-known interface:

```go
type Handler interface {
  ServeHTTP(ResponseWriter, *request)
}
```

You might also recall that you can use the `http.HandlerFunc`

```go
type HandlerFunc func(ResponseWriter, *request)
```

which is an adapter that will let you use regular functions that satisfy the signature `func(ResponseWriter, *request)` as handlers for HTTP servers.

As a result, you will be able to create a mock server very succinctly as:

```go
myServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
  // ... mock behavior here
  w.WriteHeader(http.StatusOK)
}))
```

Additionally, you can get the URL of such a mock server using:

```go
myServerURL := myServer.URL
```

With all these pieces you will be able to create a much better `TestWebRacer` implementation, solving one of the problems.

Finally, make sure you invoke:

```go
myServer.Close()
```

## v2: doing some refactoring

While we still lack some basic capabilities (e.g., concurrency, timeout control, ...) we can start doing some refactoring in `WebRacer` to prepare it for concurrent execution.

In this version, you should extract the common parts in `WebRacer` and place them in a named function.

Also: You can refactor the test as well, by creating a function that removes the duplication created when we create the new server. Additionally, you can improve the way in which we're dealing with `Close` can be improved using `defer`.

## v3: Introducing concurrency and synchronization

There are a couple of things that can be improved in the previous version:

1. Why are we measuring the response time in sequence when Go is especially good at concurrency?

2. Why are we measuring the response time at all? We only need to know which responds first, so a signal should suffice no matter whether it takes 20 ms or 5 seconds.

This can be achieved with a construct called `select`.

The `select` statement lets a goroutine wait on multiple communication operations in a clear and succinct way.

A `select` statement blocks until one of its cases run, and then it executes that case. It chooses one at random it multiple are ready:

```go
select {
case val1 <- chanX:
  // ... do something ...
case val2 <- chanY:
  // ... do something else ...
}
```

Aditionally, you can include a `default` case, which will run if no other case is ready.

```go
select {
case i := <-c:  // blocks until something is ready on c
  // ... do something with i ...
default:
  // ... do something if c is blocking ...
}
```

Let's use this to improve the `WebRacer` function.

The implementation is quite minimalist, so attention is needed.

Let's start with a `ping` function whose responsibility will be to signal as soon as the response from the server is obtained:

```go
func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		_, _ = http.Get(url)
		close(ch)
	}()
	return ch
}
```

Note that the function received the URL to ping, and returns a channel. The type is not important.

| NOTE: |
| :---- |
| Why a `chan struct{}` instead of a `chan bool`? A `chan struct{}` is the smallest data type available from a memory perspective because it doesn't allocate anything. |

Then, within the function, we create the channel and probe the URL in the background using a goroutine. In the goroutine we do the HTTP GET and immediately after close the channel. This will indicate that nothing more will be received in that channel, and that can be used as the signal itself. Additionally, that will cause the `select` to take that case.

Then, we change the implementation of the `WebRacer` as follows:

```go
func WebRacer(url1, url2 string) (fasterURL string) {
	select {
	case <-ping(url1):
		return url1
	case <-ping(url2):
		return url2
	}
}
```

Again, very hacky and succinct implementation. We use select to wait on both the channels returned by `ping(url1)` and `ping(url2)` operations. The `close(ch)` operation performed in the function will make select take the corresponding case.

## v4: Implementing the timeout

We also had the requirement to return an error if `Racer` took longer than 10 seconds.

As firm TDD believers, we should start with the test first.

As we Go, error conditions in Go are typically returned as additional field in the returned value. As such, the signature for the `WebRacer` function should be changed to:

```go
url, err := WebRacer(server1URL, server2URL)
```

Then, to implement the timeout you can use another case in the select statement:

```go
select {
case <-ping(url1):
  return url1, nil
case <-ping(url2):
  return url2, nil
case <-time.After(10 * time.Second):
  return "", fmt.Errorf("timed out waiting for response from %q and %q", url1, url2)
}
```

The `time.After` function returns a channel that you can receive from. The channel will receive some information after the specified time, so it is the best possible way to implement the timeout.

## v5: fixing the slow test

If you run the test for v4, you will see that you need to wait a long time to get the results. As stated in previous exercises, having long tests is not a good idea.

We also saw in the previous exercises that this can be resolved using either mocking (creating an implementation of `Sleep` or in this case `time.After` that doesn't require the full ten seconds to resolve), or by creating a *configurable* version of the function.

We've been told that the real function needs to wait for the full 10 seconds (so this value should be hardcoded), but we want to be sympathetic with our test requirements too.

Let's use the configurable approach to solve the problem.

```go
func WebRacer(url1, url2 string) (fasterURL string, err error) {
	return ConfigurableWebRacer(url1, url2, 10 * time.Second)
}

func ConfigurableWebRacer(url1, url2 string, timeout time.Duration) (fasterURL string, err error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for response from %q and %q", url1, url2)
	}
}
```

After that, you can let `WebRacer` be used in your happy path tests, while you use the `ConfigurableWebRacer` only on your sad path tests so that both implementations are tested.