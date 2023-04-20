package slices_test

import (
	"fmt"
	"testing"

	"example.com/05_tdd-cover/slices"
)

func TestSumAll(t *testing.T) {
	slicesEqual := func(s1, s2 []int) bool {
		if len(s1) != len(s2) {
			return false
		}
		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				return false
			}
		}
		return true
	}

	t.Run("Sum two non-empty slices", func(t *testing.T) {
		want := []int{3, 9}
		got := slices.SumAll([]int{1, 2}, []int{0, 9})
		if !slicesEqual(got, want) {
			t.Errorf("got %v, but wanted %v", got, want)
		}
	})

	t.Run("Sum single non-empty slice", func(t *testing.T) {
		want := []int{3}
		got := slices.SumAll([]int{1, 1, 1})
		if !slicesEqual(got, want) {
			t.Errorf("got %v, but wanted %v", got, want)
		}
	})

	t.Run("Sum empty slices", func(t *testing.T) {
		want := []int{0, 0}
		got := slices.SumAll([]int{}, []int{})
		if !slicesEqual(got, want) {
			t.Errorf("got %v, but wanted %v", got, want)
		}
	})

	t.Run("Sum nil slices", func(t *testing.T) {
		want := []int{0, 0}
		got := slices.SumAll(nil, nil)
		if !slicesEqual(got, want) {
			t.Errorf("got %v, but wanted %v", got, want)
		}
	})
}


func ExampleSumAll() {
	result := slices.SumAll([]int{1, 2}, []int{0, 9})
	fmt.Println(result)
	// Output:
	// [3 9]
}

func Example_SumAll_2() {
	result := slices.SumAll([]int{1, 1, 1})
	fmt.Println(result)
	// Output:
	// [3]
}
