# fetch: prints the body obtained after sending an HTTP GET to the given URLs

## v0: initial implementation

Create a program that issues an HTTP GET request to the list of URLs received as arguments in the CLI.

HINTS:
+ use `http.Get` to send an HTTP GET request to a given URL
+ use `os.Exit(1)` to terminate the program when an error is found
+ use `ioutil.ReadAll` on the response to obtain the response from the response `Body` field.

## v1: a few minor enhancements

The funcion `io.Copy(dst, src)`. Use it to stream the response directly to `os.Stdout` to prevent having to materialize a buffer with the response.

Also, inspect the arguments to add the prefix `https://` if it doesn't include one. HINT: use `strings.HasPrefix`.

Finally, print also the HTTP status code. HINT: use `resp.Status`.