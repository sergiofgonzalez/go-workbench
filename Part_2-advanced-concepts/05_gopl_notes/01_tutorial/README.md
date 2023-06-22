# 01 &mdash; Tutorial

## Packages

Go code is organized into packages, which are similar to libraries or modules in other languages.

A package consists of one or more `.go` source files in a single directory that define what the package does.

Each source file begins with a package declaration, as in `package main`.

The Go standard library has over 100 packages for common tasks such as input and output management, sorting, text manipulation, etc.

Package `main` is special. It defines a standalone executable program, not a library. Within package `main` the function `main` is where the execution of the program begins.

## Command-line arguments with `os.Args`

In Go, the `os` package provides functions and other values for dealing with the OS in a platform-independent fashion.

In particular, command-line arguments are available to a program in a variable names `Args` that is part of the `os` package. As a result, you can access the arguments using `os.Args`.

The variable `os.Args` is a slice of strings, and slices can be thought as dynamically sized sequence of array elements where individual elements can be accessed as `s[i]` and contiguous subsequence as `s[m:n]`, where the number of elements is `len(s)`.

Go uses the *half-open* intervals that include the first and exclude the last so `s[m:n]`, where $ 0 \le m \le n \le len(s)$ contains $ n - m $ elements.

The first element of `os.Args` is `os.Args[0]` and contain the name of the command itself. Therefore, the slice expression `os.Args[1:len(os.Args)]`or `os.Args[1:]` will give you all the arguments.

### Lab

Create an implementation of the Unix command `echo`, which prints its command-line arguments on a single line.

HINT: use `os.Args` as described above.

| EXAMPLE: |
| :------- |
| See [01_echo](01_echo/README.md) for a runnable example. |

## Finding duplicate lines

Programs for file copying, printing, searching, sorting, counting, etc. have a similar structure: a loop over the input, some computation on each element, and generation of output on the fly or at the end.

The `bufio` package helps make input and output efficient and convenient. The type `Scanner` reads input and breaks it into lines or words, and it's often the easiest way to process input that comes naturally in lines.

To make a scanner read from the standard input you can do:

```go
input := bufio.NewScanner(os.Stdin)
```

Each call to `input.Scan()` reads the next line and removes the newline character from the end. It returns `true` if there is a line and `false` if there's no more input. You can access the results using `input.Text()`.

As a result, you can iterate over the lines you get in `os.Stdin` doing:

```go
input := bufio.NewScanner(os.Stdin)
for input.Scan() {
  line := input.Text()
  // ... do something with line ...
}
```

An alternative approach will make use of `ioutil.ReadFile` which returns the whole contents of a file as a `[]byte`. Then you can use `strings.Split` (which is the opposite of `strings.Join`) to slice the contents by newline and then count them.

### Lab

Create a program named `dup` which looks for duplicate lines.

The program should print the text of each line that appears more than once preceded by its count.

| EXAMPLE: |
| :------- |
| See [02_dup](02_dup/README.md) for a runnable example. |

## Animated GIFs

After importing a package with multiple components, like `image/color`, we refer to the package with a names that comes from the last component:

```go
import (
  "image"
  "image/color"
  ...
)

...

// Composite literals
var palette = []color.Color{color.White, color.Black}
```

Composite literals is a compact notation for instantiating any of Go's composite types (such as slices, arrays, maps or structs) from a sequence of element values.

### Lab

See [03_math_gifs](03_math_gifs/README.md) for a small program that create an animated GIF consisting of Lissajous figures.

## Fetching a URL

The `net/http` and `io/ioutil` packages are useful when dealing with HTTP requests and responses handling.

```go
resp, err := http.Get(url)
if err != nil {
  // ... handle error ...
}
b, err := ioutil.ReadAll(resp.Body)
resp.Body.Close()
```

The `http.Get` function makes an HTTP request, and if there is no error, returns the result in the response struct ``resp`.

The `Body` field of `resp` contains the server response as a readable stream. The function `ioutil.ReadAll` reads the entire response into a slice of bytes.

Then, the `Body` stream is closed to avoid leaking resources.

The field `resp.Status` holds the HTTP status code of the request.

### Lab

See [04_fetch_url](04_fetch_url/README.md) for an example that probes a URL received as an argument.

## Concurrency: Fetching URLs concurrently

One of the most interesting and novel aspects of Go is its support for concurrent programming using *goroutines* and *channels*.

A goroutine is a concurrent function execution. A channel is a communication mechanism that allows one goroutine to pass values of a specified type to another goroutine.

A channel is created using:

```go
ch := make(chan string)
```

You can start the execution of a new *goroutine* doing:

```go
go fetch(url, ch)

func fetch(url string, ch chan<- string) {

}
```

When one goroutine attempts a send or receive on a channel, it blocks until another goroutine attempts the corresponding receive or send operation, at which point the value is transferred and both goroutines proceed.

```go
ch := make(chan string)
go fetch(url, ch)
fmt.Printf(<-ch)  // receive from channel ch


func fetch(url string ch chan<- string) {
  // .. probe url and read response .../

  // Send to channel ch
  ch <- fmt.Sprintf("time elapsed: %.2fs", elapsedSeconds)
}
```

### Lab

See [05_concurrent_fetch](05_concurrent_fetch/README.md) for an example illustrating the *goroutines* and *channels* concepts.

## A Web Server

Go's libraries make it very easy to write a web server that responds to client requests like those made by fetch.

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", rootPathHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func rootPathHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("request received from %q\n", r.URL.Path)
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
```

The line:

```go
http.HandleFunc("/", rootPathHandler)
```

connects a handler function to incoming URLs that begin with `/` (which is all URLs).

A request is represented as a struct of type `http.Request`.

The previous example can be modified to add new endpoints:

```go
http.HandleFunc("/", rootPathHandler)
http.HandleFunc("/count", countHandler)
```

In the example above:
+ A request for `/count` will go to the `countHandler`
+ Any other request (including `/count/hello`, for example) will end up being handled by `rootPathHandler`.

`http.ResponseWriter` is an `io.Writer` interface, which allows you to use the interface value anywhere where where an `io.Writer` is expected. For example, we can use the following to stream the Lissajous gif to the browser:

```go
http.HandleFunc("/lissajous", func(w http.ResponseWrite, r *httpRequest) {
  lissajous(w)
})
```

### Lab

See [06_web-server](06_web-server/README.md) for an example illustrating the concepts of this section.

## Misc loose ends

### Pointers

Go provides pointers, that is, values that contain the address of a variable. Pointers in Go sit in the middle between the lack of control you have in C, and the complete abstraction you see in other programming languages like in Java and JavaScript.

### Methods and interfaces

A method is a function associated with a named type. In Go you can attach a method to almost any named type.

Interfaces are abstract types that let you treat different concrete types in the same way based on what methods they have, not how they are represented or implemented.






