package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

func main() {
	var i interface{}
	describe(i)

	i = "Hello!"
	describe(i)

	i = T{"Hello to Jason Isaacs!"}
	describe(i)
}