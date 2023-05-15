# Managing long-running processes with `context` in Go using TDD
> using a webserver to learn about the `context` package

In the example, we're going to start with a webserver that might potentially hit off a long running process to fetch some data and return it.

We want to implement a scenario in which the user cancels a long running request while it is in progress, and we want to do it cleanly, telling the process to give up.

## v0: testing the happy path with a stubbed implementation

In v0 we have the fundamental pieces of the solution.

We start by defining the `Store` interface which represents the behavior available in the store, that is, its service-based API.

Initially, the store supports a `Fetch` method that returns a string response:

```go
type Store interface {
  Fetch() string
}
```

Then, we need to expose the functionality in an HTTP server. An idiomatic way to do so is to define a function that receives an interface and returns the corresponding handler, so that it can be ultimately wired.

```go
func Server(store Store) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, store.Fetch())
  }
}
```

`HandlerFunc` is one of the key types representing the function signature of an HTTP handler in Go.

For example, you could potentially associate a given path to the handler doing:

```go
http.HandleFunc("/items", Server(myStore))
```

In our tests, we need to be able to test the `Server` function. This means, we need to ensure that we can use the value returned from the `Server` function to invoke `ServeHTTP` and that the `response.Body()` is whatever the `Fetch` method provided.

In order to do so, we start by defining a `StubStore` that will be used to return *stubbed* responses. We will want this struct to implement the `Store` interface:

```go
type StubStore struct {
  response string
}

func (s *StubStore) Fetch() string {
  return s.response
}
```

Then in the test, we use both capabilities from `http` and `httptest` packages:

```go
func TestServer(t *testing.T) {
  data := "Hello to Jason Isaacs"
  server := Server{&StubStore{data}}

  request := httptest.NewRequest(http.MethodGet, "/", nil)
  response := httptest.NewRecorder()

  server.ServeHTTP(response, request)
  got := response.Body.String()

  if got != data {
    // ... error ...
  }
}
```

## v1: introducing Cancel

As `Fetch` can potentially be a long-running request, we want to enhance our application with a `Cancel` method:

```go
type SpyStore struct {
  response  string
  cancelled bool
}

func (s *SpyStore) Fetch() string {
  time.Sleep(100 * time.Millisecond)
  return s.response
}

func (s *SpyStore) Cancel() {
  s.cancelled = true
}
```

Then, we prepare a subtest in which we cancel the request before the 100 milliseconds to see if it gets cancelled.


```go
t.Run("fetch gets cancelled", func(t *testing.T) {
  data := "Hello to Jasaon Isaacs"
  store := &SpyStore{response: data}
  server := Server(store)

  request := httptest.NewRequest(http.MethodGet, "/", nil)

  cancellingCtx, cancel := context.WithCancel(request.Context())
  time.AfterFunc(5 * time.Millisecond, cancel)
  request = request.WithContext(cancellingCtx)

  response := httptest.NewRecorder()

  server.ServeHTTP(response, request)

  if !store.cancelled {
    t.Errorf("store was not told to cancel when it should")
  }
})
```

What we do is derive a new `cancellingCtx` from our `request` which returns a `cancel` function. We schedule that function to be called after 5 milliseconds using `time.AfterFunc`. Finally, we use this new context in our requiest calling `request.WithContext`.

When you run the test it will fail. So we will need to provide a new implementation for the `Server` function that cancels.

```go
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		data := make(chan string, 1)
		go func() {
			data <- store.Fetch()
		}()

		select {
		case d := <-data:
			fmt.Fprintf(w, d) // d is the response of store.Fetch
		case <-ctx.Done():
			store.Cancel()
		}
	}
}
```

`context` has a method `Done` which returns a channel that gets sent a signal when the context is "done" or "cancelled".

We use a `select` to listen to that signal, while also waiting for `Fetch` to complete. If `Fetch` gets resolved before the cancelling signal we want to show the results.

## v2: a bit of refactoring

Now we can accommodate a bit of refactoring by including the assertions in the spy.

## v3: following context best practices

To do once I understand better http servers and context.


## ToDo

Review https://go.dev/blog/context to understand what's a context. https://pkg.go.dev/context