package main

import (
	"fmt"

	"example.com/tracing/tracing"
)

func a() {
	defer tracing.Un(tracing.Trace("a"))
	fmt.Println("In a()")
}

func b() {
	defer tracing.Un(tracing.Trace("b"))
	fmt.Println("In b()")
	a()
}


func main() {
	b()
}