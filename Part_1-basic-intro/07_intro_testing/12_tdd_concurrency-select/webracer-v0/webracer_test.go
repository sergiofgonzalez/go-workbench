package webracer

import "testing"

func TestWebRacer(t *testing.T) {
	slowURL := "http://some.slow.url.com"
	fastURL := "http://some.fast.url.com"

	want := fastURL
	got := WebRacer(slowURL, fastURL)

	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}