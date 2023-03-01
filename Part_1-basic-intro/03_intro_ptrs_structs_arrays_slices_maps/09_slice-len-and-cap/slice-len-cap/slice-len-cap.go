package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)		// len=6, cap=6 [2 3 5 7 11 13]

	s = s[:0]
	printSlice(s)		// len=0, cap=6 []

	// This is unexpected: orig items from the capacity are resuscitated
	s = s[:4]
	printSlice(s)		// len=4, cap=6 [2 3 5 7]

	s = s[2:]
	printSlice(s)		// len=2, cap=4 [5 7]

	// Note that it is a bit trickier than it seems
	// Note that when the slice was empty, any subsequent re-slicing
	// brought back some of the items

	items := []string{"hello", "to", "jason", "isaacs", "!!!"}
	fmt.Printf("len=%d, cap=%d %v\n", len(items), cap(items), items)

	items = items[2:4] 	// len=2, cap=3 [jason isaacs]
	fmt.Printf("len=%d, cap=%d %v\n", len(items), cap(items), items)

	// Now I empty it
	items = items[:0]		// len=0, cap=3 []
	fmt.Printf("len=%d, cap=%d %v\n", len(items), cap(items), items)

	// Any subsequent re-slicing refers to the previous backing slice
	// thus:
	items = items[0:3]  // len=3, cap=3 [jason isaacs !!!]
	fmt.Printf("len=%d, cap=%d %v\n", len(items), cap(items), items)
}


func printSlice(s []int) {
	// %d is base 10 integer
	fmt.Printf("len=%d, cap=%d %v\n", len(s), cap(s), s)
}