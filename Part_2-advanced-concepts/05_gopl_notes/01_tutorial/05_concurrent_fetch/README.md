# fetch_all
> probes many URLs received received as arguments concurrently

## v0: Initial version

Create a program that performs an HTTP GET of all the URLs received from the command line in parallel, reporting their time and sizes.

Implement `main` using the following guidelines
1. Start by taking the start time (HINT: use `time.Now()`)
2. Create a channel to receive all the output from the goroutines
3. Implement a loop to perform the GET operation as a goroutine. The function should be passed the URL to call and the channel.
4. Then, iterate over the arguments received using `<-ch` to receive from the channel. That operation will block until a message is received from the channel.
5. Then, print the elapsed overall time using `time.Since(start).Seconds()`.

Then you need to implement the *goroutine*, that receives the url to probe and an output channel to write to (HINT: define the channel as `ch chan<- string` to indicate is an output channel).
1. Start by taking the time, so that you can report how long it took to fetch the URL and read the response contents.
2. Issue an `http.Get`.
3. Control the error. If an error is found, send it to the channel.
4. If it's not an error, read the response body discarding its contents (HINT: using `io.Copy(ioutil.Discard, resp.Body))`).
5. Close the response body
6. If an error was found while reading the response, send the error to the channel and return.
7. Otherwise, compute the elapsed time doing `time.Since(start).Seconds()` and send the report to the channel.