package main

import (
	"fmt"

	"golang.org/x/exp/maps"
)


func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || xv != yv {
			return false
		}
	}
	return true
}

func main() {
	x := map[string]int {
		"A": 0,
		"B": 1,
		"C": 0,
	}

	y := map[string]int {
		"A": 0,
		"B": 1,
	}

	fmt.Println("equal: ", equal(x, y))
	fmt.Println("Equal: ", maps.Equal(x, y))
}