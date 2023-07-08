package main

import "fmt"

func nonEmpty(strings []string) []string {
	out := strings[:0]	// zero-length string of the original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}


func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("before  : %q\n", data)						// ["one" "" "three"]
	fmt.Printf("nonEmpty: %q\n", nonEmpty(data)) 	// ["one" "three"]
	fmt.Printf("after   : %q\n", data) 						// ["one" "three" "three"]
}