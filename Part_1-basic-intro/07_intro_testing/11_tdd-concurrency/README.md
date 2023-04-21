# Concurrency in Go using TDD
> this example illustrates concurrency concepts in Go


You are assigned with the maintenance and evolution of a library that checks the status of a list of URLs.

```go
type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
	}

	return results
}
```

The function returns a map of each URL checked to a boolean value: true for a good response, false for a bad response.

The `WebsiteChecker` is a function that takes a single URL and returns a boolean.

When you receive the library, you discover that the team that developed it used *DI* to test the function without making real HTTP calls, making it reliable and fast.

```go
func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurtwe.geds" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurtwe.geds",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurtwe.geds":      false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if maps.Equal(want, got) {
		t.Errorf("wanted %v, got %v", want, got)
	}
}
```

The function is in production, and while working well, you've started to get complains that the function is too slow.

## v0: benchmarking

In order to have a baseline of the performance, it will be necessary to include a benchmark in the test, so that we can compare the changes we introduce.

Write a benchmark to check how long it takes to check 100 websites. To make it more realistic, create a new mocked implementation of the website checker that takes 20 milliseconds to check a website and returns true `stubSlowWebsiteChecker`.

Then in the benchmark, arrange an array of 100 urls and run the tests.

HINT: the time it takes to do the arrangements shouldn't be taken into consideration. Make sure you use `ResetTimer` to signal the start of the real test benchmark.

Once implemented, run `go test -bench="."` and annotate the results:

```bash
$ go test -bench="."
goos: linux
goarch: amd64
pkg: example.com/11_tdd-concurrency
cpu: 11th Gen Intel(R) Core(TM) i7-11850H @ 2.50GHz
BenchmarkCheckWebsites-16              1        2043933241 ns/op
PASS
ok      example.com/11_tdd-concurrency  2.049s
```

## v1: a naive concurrent solution that fails

Naively, in this version you try to make the solution faster by running the website checker in a goroutine. As that will be executed in a non-blocking way, you are expecting to get the results faster.

HINT: use an anonymous function to implement the concurrency.

Confirm that the function is no longer working. Can you provide an explanation?

Try to find a quick a dirty fix for the function.

HINT: introduce a delay to allow the functions to add the result to the map before the program is complete. Also, you might need to prevent the loop variable to be reused.

Run `go test -race` and see that you haven't really fixed anything.

## v2: a safer approach to concurrency using channels

In this version, introduce channels to coordinate the concurrency results in a safe way.

The channel will store the url, along with whether the site contacted was alive or not. Therefore, you should start by defining such structure: `siteResult`.

Then, in the goroutine, instead of updating the map, collect the results in a value and send it to the channel.

Finally, implement a loop to iterate over the urls collecting the results and placing them in the map.

Confirm that now the test runs and there are no race conditions.

## v3: simplyfing the implementation

While v2 implementation is clear, is a bit verbose. Do some refactoring to make it more succinct:

+ do not create names for the `siteResult` data structure.
+ do not create a value separately before feeding it into the channel - do it in one step
+ update the loop to get the values from the channel without using the variable names.

Run the test benchmark and compare the results with the baseline:

```bash
$ go test -bench="."
goos: linux
goarch: amd64
pkg: example.com/11_tdd-concurrency
cpu: 11th Gen Intel(R) Core(TM) i7-11850H @ 2.50GHz
BenchmarkCheckWebsites-16             56          21042252 ns/op
PASS
ok      example.com/11_tdd-concurrency  1.206s
```

| TEST | ns/op | Total time (s) |
| :--- | -----: | ---------: |
| BEFORE | 2043933241 | 2.049 |
| AFTER  | 21042252 | 1.206 |

Clearly, adding concurrency to the function is a good idea.

Note that the solution works OK because what we're doing is pretty lightweight, but we should be aware that we're not putting any limit to the amount of goroutines that are spawned.

If we're expecting to have thousands of website checkers, it might be a better idea to set up a set of workers with a particular size so that you don't saturate the resources of the machine in which the program is running.