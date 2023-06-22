package main

import (
	"fmt"
	"math/rand"
)

// Number is a container type for ints and floats
type Number interface {
	int | int64 | float64
}

// SortedList is a container for numbers that always keep its values sorted in ascending order
type SortedList[T Number] struct {
	values []T
}

func (list *SortedList[T]) Values() []T {
	// return a copy of the values
	return append([]T{}, list.values...)
}

func (list *SortedList[T]) Add(value T) {
	for i, v := range list.values {
		if v > value {
			list.values = append(list.values[:i+1], list.values[i:]...)
			list.values[i] = value
			return
		}
	}
	list.values = append(list.values, value)
}

func main() {
	l := SortedList[float64]{}
	for i := 0; i < 10; i++ {
		l.Add(rand.Float64())
	}
	fmt.Println(l.Values())
}
