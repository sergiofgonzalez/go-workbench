package main

import "fmt"

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		// Not enough space, allocate a new array doubling current len
		// if desired capacity is below the double of current len
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)	// built-in function that copies elements from one slice to another
	}
	z[len(x)] = y
	return z
}

func main() {
	// basic
	x := []int{1, 2, 3}
	z := appendInt(x, 4)
	fmt.Printf("%v\n", z)

	// how we enlarge the capacity
	var x1, x2 []int
	for i := 0; i < 10; i++ {
		x2 = appendInt(x1, i)
		fmt.Printf("%d cap=%2d\t%v\n", i, cap(x2), x2)
		x1 = x2
	}

	// what happens
	a := [...]int{0, 1, 2, 3}
	x = a[:3]
	fmt.Printf("%v\tlen=%d\tcap=%d\n", x, len(x), cap(x))
}