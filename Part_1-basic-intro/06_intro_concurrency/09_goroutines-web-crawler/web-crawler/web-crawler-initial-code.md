Starting point for the crawler

```go
package main

import "fmt"

type Fetcher interface {
	// Fetch returns the body of URL and a slice of URLs found on that page
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl pages starting with url, to a maximum of depth
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel
	// TODO: Don't fech the same URL twice
	// This implementation should be updated
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth - 1, fetcher)
	}
	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is a mock implementation of Fetcher that returns canned results
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body 	string
	urls	[]string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher with some canned results
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
```

## Step 1: non-concurrent solution

This version only implements the caching part

```go
package main

import "fmt"

type Fetcher interface {
	// Fetch returns the body of URL and a slice of URLs found on that page
	Fetch(url string) (body string, urls []string, err error)
}

var visitedUrls map[string]bool

// Crawl uses fetcher to recursively crawl pages starting with url, to a maximum of depth
func Crawl(url string, depth int, fetcher Fetcher) {
	if _, ok := visitedUrls[url]; ok {
		fmt.Printf("skipping %s as it has been already been crawled\n", url)
		return
	}

	visitedUrls[url] = true

	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth - 1, fetcher)
	}
	return
}

// Commented out to prevent execution
func main() {
	visitedUrls = make(map[string]bool)
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is a mock implementation of Fetcher that returns canned results
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body 	string
	urls	[]string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher with some canned results
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
```