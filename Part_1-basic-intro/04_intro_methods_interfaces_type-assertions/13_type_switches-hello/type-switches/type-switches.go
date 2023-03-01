package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("The double of %v is %v\n", v, 2 * v)
	case string:
		fmt.Printf("The string %q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I've received a %T and I don't know what to do with it\n", v)
	}
}

type Vertex struct {
	X, Y float64
}

func main() {
	do(5)
	do("Hello to Jason Isaacs!")
	do(true)
	do(Vertex{3, 4})
}