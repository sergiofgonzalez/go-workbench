package main

import (
	"fmt"
	"strings"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	sb := strings.Builder{}
	// n % 3 gives us the pos of the leftmost decimal sep
	// when there are a number of characters that are 
	firstDecSep := n % 3
	if firstDecSep != 0 {
		for i := 0; i < firstDecSep; i++ {
			sb.WriteByte(s[i])
		}
		sb.WriteString(",")
	}

	// now we put a decimal separator every three positions
	for i, j := firstDecSep, 1; i < n; i, j = i+1, j+1 {
		sb.WriteByte(s[i])
		if j % 3 == 0 && i != n - 1{
			sb.WriteString(",")
		}
	}

	return sb.String()
}

// len:  4    | 7         | 10
// i:    0123 | 0123456   | 0123456789
// s:    1234 | 1234567   | 1234567890
// res: 1,234 | 1,234,567 | 1,234,567,890 -> 10 / 3 = 3, 10 % 3 = 1
// div: 1     | 2         | 3
// rem: 1     | 1         | 1
func main() {
	fmt.Println("Run 'go test'")
}