package concurrency

import (
	"testing"
	"time"

	"golang.org/x/exp/maps"
)

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

	if !maps.Equal(want, got) {
		t.Errorf("wanted %v, got %v", want, got)
	}
}

func stubSlowWebsiteChecker(url string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "http://some.url.com"
	}

	b.ResetTimer()	// Signal the real beginning of the test
	for i := 0; i < b.N; i++ {
		_ = CheckWebsites(stubSlowWebsiteChecker, urls)
	}
}
