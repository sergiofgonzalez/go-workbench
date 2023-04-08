package main

import "fmt"


func simple() {
	const (
		first = iota
		second
		third
		fourth
	)
	fmt.Println("=== simple")
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
	fmt.Println(fourth)
	fmt.Printf("==========\n\n")
}


func simpleTyped() {
	const (
		first int = iota
		second
		third
		fourth
	)
	fmt.Println("=== simple (typed)")
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
	fmt.Println(fourth)
	fmt.Printf("==========\n\n")
}


func doubling() {
	const (
		first int = iota
		second = iota * 2
		third
		fourth
	)
	fmt.Println("=== Doubling")
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
	fmt.Println(fourth)
	fmt.Printf("==========\n\n")
}

func power2Contrived() {
	const (
		first float64 = iota	// 2^0
		second = 1 << iota		// 2^1
		third									// 2^2
		fourth								// 2^3
	)

	fmt.Println("=== Power of 2 (contrived)")
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
	fmt.Println(fourth)
	fmt.Printf("==========\n\n")
}

func storage() {
	const (
		_ float64 = iota
		KB = 1 << (10 * iota) // KB == (1 << 10) == 2^10
		MB										// MB == (1 << 20) == 2^20
		GB										// GB == (1 << 30) == 2^30
		TB										// TB == (1 << 40) == 2^40
	)

	fmt.Println("=== Power of 2 (contrived)")
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Printf("==========\n\n")

}


func main() {

	simple()
	simpleTyped()
	doubling()
	power2Contrived()
	storage()
}