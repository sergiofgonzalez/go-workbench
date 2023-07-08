package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	chars := make(map[rune]int)
	for input.Scan() {
		line := input.Text()
		for _, r := range line {
			chars[r]++
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "charcount: %v", err)
	}

	for r, cnt := range chars {
		fmt.Printf("%c\t%#x\t%d\n", r, r, cnt)
	}
}