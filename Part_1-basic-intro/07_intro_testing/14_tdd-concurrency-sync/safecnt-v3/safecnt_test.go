package safecnt

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("single thread", func(t *testing.T) {
		cnt := NewCounter()
		cnt.Inc()
		cnt.Inc()
		cnt.Inc()

		want := 3
		assertCounterValue(t, cnt, want)
	})

	t.Run("concurrent", func(t *testing.T) {
		cnt := NewCounter()
		want := 1000

		var wg sync.WaitGroup
		wg.Add(want)

		for i := 0; i < 1000; i++ {
			go func () {
				cnt.Inc()
				wg.Done()
			}()
		}

		wg.Wait()
		assertCounterValue(t, cnt, want)
	})
}

func assertCounterValue(t testing.TB, c *Counter, want int) {
	t.Helper()
	if c.Value() != want {
		t.Errorf("unexpected value: got %d, but wanted %d", c.Value(), want)
	}
}