package iteration_test

import (
	"fmt"
	"testing"

	"example.com/04_tdd-benchmarks/iteration"
)

func TestRepeat(t *testing.T) {
	got := iteration.Repeat("a")
	want := "aaaaa"

	if got != want {
		t.Errorf("got %q, but expected %q", got, want)
	}
}

func ExampleRepeat() {
	result := iteration.Repeat("abc")
	fmt.Println(result)
	// Output:
	// abcabcabcabcabc

}


func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iteration.Repeat("a")
	}
}

func BenchmarkRepeatWithSb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iteration.RepeatWithSb("a")
	}
}

func BenchmarkRepeatNative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iteration.RepeatNative("a")
	}
}