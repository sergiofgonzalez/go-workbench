package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := "X"
	c2 := "x"

	h1 := sha256.Sum256([]byte(c1))
	h2 := sha256.Sum256([]byte(c2))

	fmt.Printf("h1: %x\nh2: %x\nh1 == h2: %t\nSHA256 type: %T\n", h1, h2, h1 == h2, h1)
}