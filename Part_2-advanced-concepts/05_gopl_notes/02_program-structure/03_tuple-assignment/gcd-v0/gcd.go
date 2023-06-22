package main

import "fmt"

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func main() {
	fmt.Println(gcd(2, 3))
	fmt.Println(gcd(4, 6))
	fmt.Println(gcd(6, 36))
}