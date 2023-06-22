package main

import "fmt"

type Flags uint

const (
	One Flags = 1 << iota	// 1
	Two
	Four
	Eight
)

func main() {
	fmt.Printf("%b \t %[1]d \n", One)
	fmt.Printf("%b \t %[1]d \n", Two)
	fmt.Printf("%b \t %[1]d \n", Four)
	fmt.Printf("%b \t %[1]d \n", Eight)
}

