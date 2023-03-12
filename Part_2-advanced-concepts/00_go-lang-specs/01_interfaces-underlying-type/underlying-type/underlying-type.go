package main

import "fmt"

type MyInt int

type Integer interface {
	~int
}


func Triple[T Integer](n T) MyInt {
	return MyInt(n) * 3
}


type Float interface {
	~float32 | ~float64
}


func Double[T Float](n T) float64 {
	return float64(n) * 2
}

// Cannot use Float outside of a type constraint
// func MyFunc(n Float) Float {
// 	return n * 2
// }

func main() {
	var n MyInt = 5
	fmt.Println(Triple(n))

	fmt.Println(Double(2.5))
}