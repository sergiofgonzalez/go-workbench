package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	i := T{"Hello to Jason Isaacs!!"}
	i.M()
}