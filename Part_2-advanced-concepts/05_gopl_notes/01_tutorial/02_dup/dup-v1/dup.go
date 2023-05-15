package dup

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// GetCounts returns a frequency map of the lines found in Stdin.
func GetCounts() map[string]int {
	counts := make(map[string]int)
	ConfigurableGetCounts(os.Stdin, counts)
	return counts
}

// WriteDups prints in the standard output the keys in the given frequency map whose
// associated count is greater than one
func WriteDups(counts map[string]int) {
	ConfigurableWriteDups(os.Stdout, counts)
}

// ConfigurableGetCounts returns a frequency map for the line found in in the given readder.
func ConfigurableGetCounts(r io.Reader, counts map[string]int) {
	input := bufio.NewScanner(r)
	for input.Scan() {
		line := input.Text()
		counts[line]++
	}
}

// ConfigurableWriteDups writes in the given writer the keys of the received frequency map
// whose associated count is greater than 1.
func ConfigurableWriteDups(w io.Writer, counts map[string]int) {
	for line, cnt := range counts {
		if cnt > 1 {
			fmt.Fprintf(w, "%d\t%s\n", cnt, line)
		}
	}
}
