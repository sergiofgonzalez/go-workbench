package main

import "fmt"

type Person struct {
	Name 	string
	Age		int
}

func (p Person) String() string {
	fmt.Println("\t=> String() method called!")
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Jason Isaacs", 56}
	fmt.Println(a)
}