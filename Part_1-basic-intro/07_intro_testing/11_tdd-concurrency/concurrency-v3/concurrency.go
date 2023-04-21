package concurrency

// WebsiteChecker is the signature of a function that is used to check if a site
// is alive
type WebsiteChecker func(string) bool

type siteResult struct {
	string
	bool
}

// CheckWebsites uses the given website checker function on each of the urls passed
// and return a map indexed by urls with the results
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	resultsChan := make(chan siteResult) // blocking channel

	for _, url := range urls {

		url := url // idiomatic as per Effective Go
		go func() {
			resultsChan <- siteResult{url, wc(url)}
		}()
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultsChan
		results[result.string] = result.bool
	}

	return results
}
