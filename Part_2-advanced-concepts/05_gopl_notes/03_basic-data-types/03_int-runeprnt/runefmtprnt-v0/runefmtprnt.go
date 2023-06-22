package main

import "fmt"

func main() {
	ascii := 'a'
	unicode := '𢉩'
	newline := '\n'

	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]q\n", newline) // no %c to prevent newline
}
