package main

import (
	"fmt"
	"math/rand"
)

func main() {
	number := rand.Intn(10) + 1
	fmt.Println(fmt.Sprintf("Hello, my favorite number is %v", number))
}