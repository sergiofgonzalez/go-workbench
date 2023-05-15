package main

import (
	"fmt"
	"os"

	"example.com/dup"
)

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		counts := dup.GetCounts()
		dup.WriteDups(counts)
	} else {
		counts := make(map[string]int)
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			defer f.Close()
			dup.ConfigurableGetCounts(f, counts)
		}
		dup.WriteDups(counts)
	}
}
