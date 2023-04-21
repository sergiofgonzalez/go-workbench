package concurrency

import "time"


type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		// naively using goroutines
		url := url	// idiomatic as per Effective Go
		go func() {
			results[url] = wc(url)
		}()
	}

	// wait for all the functions to complete
	time.Sleep(5 * time.Second)

	return results
}