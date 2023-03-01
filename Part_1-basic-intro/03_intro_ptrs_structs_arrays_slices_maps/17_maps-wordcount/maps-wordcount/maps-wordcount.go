package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount returns a map with the counts of each word found in s
func WordCount(s string) map[string]int {
	wc := make(map[string]int)

	for _, w := range strings.Fields(s) {
		cnt, ok := wc[w]
		if ok {
			wc[w] = cnt + 1
		} else {
			wc[w] = 1
		}
	}

	return wc
}

func main() {
	fmt.Println(WordCount("the quick brown fox jumps over the lazy dog"))
	wc.Test(WordCount)
}