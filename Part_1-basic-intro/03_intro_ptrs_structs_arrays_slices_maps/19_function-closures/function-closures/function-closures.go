package main

import "fmt"

func adder() func(int) int {
	total := 0
	return func(n int) int {
		total += n
		return total
	}
}

func main() {
	posAdder, negAdder := adder(), adder()

	for i := 0; i < 10; i++ {
		fmt.Println(posAdder(i), negAdder(-2*i))
	}
}