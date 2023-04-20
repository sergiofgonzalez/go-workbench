# Dependency Injection using TDD
> this example illustrates how to apply dependency injection to your Go projects

## v0: getting started

Write a function `Greet` that receives an `io.Writer` and a string and writes in the underlying Writer a personalized greeting message.

HINT: As you will need to find an appropriate implementation for the `Writer`, consider the `bytes.Buffer` struct that implements the Writer interface.

## v1: using the solution on an HTTP server

This is quite a leap!

With the `Greet` function ready and tested, implement a simple HTTP server that uses the `Greet` function to return a hardcoded "Hello to Jason Isaacs!" message to any HTTP request it receives.

1. Start by defining a constant string that will identify the port in which the HTTP server will be listening to requests.

2. Then in `main`, add a couple of log statements indicating the user where the HTTP server is going to be started and how to exit the server.

3. Then, include a call to `log.Fatal` that includes as an argument the statement `http.ListenAndServe` which starts the server at the given port and with the behavior identified by the corresponding `http.Handler`.

    `http.Handler` is an interface for a value that includes a method `ServeHTTP`.

    ```go
    type Handler interface {
	    ServeHTTP(ResponseWriter, *Request)
    }
    ```

    You can create one using `http.HandlerFunc` type, which is defined as:

    ```go
    type HandlerFunc(ResponseWriter, *Request)
    ```

    That is, you can use the syntax:

    ```go
    http.HandlerFunc(myFunc)
    ```

    to provide a `HandlerFunc` to `http.ListenAndServe`.

4. Define a function that satisfies the `HandlerFunc` type. Within it, invoke the `Greet` function developed in the v0 version.

5. Start the server using `go run .` and check that it works as expected.


### Note

When doing:

```go
log.Fatal(http.ListenAndServe(port, handler))
```

We're doing the same as in:

```go
err := http.ListenAndServe(port, handler)
if err != nil {
  log.Fatal(err)
}
```

As `http.ListenAndServe` is a blocking statement, nothing will be printed while the server is running. However, if the statement fails, then error will be immediately logged.