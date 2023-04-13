# Mini-project: An HTTP server that counts the number of times a page has been visited.
> this example illustrates how to work with interfaces and methods by way of creating an HTTP server with a handler that counts the number of times a page has been visited.

In particular, in the example it is illustrated how to make an HTTP server from a struct, a plain int, a channel, and a regular function (with a given signature).

# v0: setting up the server

This section illustrates how to make a struct implement an HTTP server function.

Almost anything can have methods attached, and therefore, almost anything can satisfy an interface.

One example is the `http` package, which defines the Handler interface. Any *object* implementing the `Handler` interface can serve HTTP requests.

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

`ResponseWriter` is itself an interface that provides access to the methods needed to return the response to the client. Those methods include the standard `Write` method, so an `http.ResponseWriter` can be used wherever an `io.Writer` can be used (e.g., `fmt.Fprintf(w, "Hello, world")` can be used to provide the response in a handler).

`Request` is a struct containing a parsed representation of the request from the client.

Let's assume we only need to deal with HTTP GET request. The following snippet is a trivial implementation of a handler to count the number of times the page is visited.

```go
type Counter struct {
    n int
}

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    ctr.n++
    fmt.Fprintf(w, "counter = %d\n", ctr.n)
}
```

Now, we can set up the server:

```go
func main() {
	const port = ":8080"
	ctr := new(Counter) // same as Counter{} but returns a pointer, which we need for http.Handle

	http.Handle("/counter", ctr)

	log.Println("Starting HTTP server on", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
```

Use the previous information to create a first version of an HTTP server.

## v1: simplifying the counter

This section illustrates how to make a plain int implement an HTTP server function.

Transform the previous version of the program so that it doesn't need a struct (HINT: use a type alias)

## v2: sending a notification

This section illustrates how to make a channel implement an HTTP server function.

Enhance the previous version so that you define a channel that sends out the content of the request to a channel and associate it to the `/notification` path.

1. Define a type `Chan` of type "channel that is used to send `*http.Request` values.

2. Make `Chan` implement the `ServeHTTP` interface. In the method implementation, send `req` in the channel and send an HTTP response to the caller to indicate that the notification has been sent.

3. In order to test it, create a channel in `main` using `make`. Then define a simple go routine that takes the content from the channel and writes the URL and user-agent found in the request. HINT: as you will want to run the *process notification routine* forever, use an infinite loop in the body of the go routine function.

## v3: displaying the contents of the args used to start the server

This section illustrates how to make a regular function implement an HTTP server function.

A function can also be turned into an HTTP handler and associated with a given URL.

You can use `os.Args` to get the arguments used to start a Go binary. In this version, we will enhance the previous project to include a function that returns that value when the `/args` URL is hit.

The `http` package contains the following code that lets you use regular functions as handlers, as longs as they comply with the expected interface:

```go
// HandlerFunc is an alias for a function that mimics ServeHTTP signature
type HandlerFunc func(http.ResponseWriter, *http.Request)

// The method HandlerFunc.ServeHTTP calls the corresponding wrapped function
func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	f(w, req)
}
```

1. Define a function `ArgServer` that receives an `http.ResponseWriter` and a pointer to a `http.Request` that sends `os.Args` in the request.

2. Establish a handler for `/args` that invokes `ArgServer` when called using `http.Handle` (in order not to break the consistency for the other handlers). HINT: you will need to convert `ArgServer` into an `http.HandlerFunc`.

