package main

import (
	"example.com/dup"
)

func main() {
	counts := dup.GetCounts()
	dup.WriteDups(counts)
}
