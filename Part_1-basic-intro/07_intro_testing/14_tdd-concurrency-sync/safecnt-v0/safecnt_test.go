package safecnt

import (
	"testing"
)

func TestCounter(t *testing.T) {
	cnt := Counter{}
	cnt.Inc()
	cnt.Inc()
	cnt.Inc()

	want := 3
	got := cnt.Value()

	if got != want {
		t.Errorf("got %d, but want %d", got, want)
	}
}