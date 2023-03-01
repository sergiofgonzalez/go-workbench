package main

import "fmt"

func split(n int) (x, y int) {
	x = n * 4 / 9
	y = n - x
	return
}

func main() {
	fmt.Println(split(17))
}