package main

import "fmt"

func main() {

	a := []int{2, 3, 5, 7, 11, 13}

	b := []bool{true, false, true, false, true, false}

	c := []struct{
		num int
		b		bool
	}{
		{ 2, true },
		{ 3, false },
		{ 5, true },
		{ 7, false },
		{ 11, true },
		{ 13, false },
	}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}