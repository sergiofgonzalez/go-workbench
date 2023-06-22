package main

import "fmt"

func main() {
	o := 0666 // octal, for POSIX file-permissions
	fmt.Printf("%d %[1]o %#[1]o\n", o)

	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
}