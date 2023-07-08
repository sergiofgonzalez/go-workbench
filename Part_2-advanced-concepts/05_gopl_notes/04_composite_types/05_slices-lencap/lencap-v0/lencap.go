package main

import "fmt"

func main() {
	// months is an array
	months := [...]string{1: "Jan", 2: "Feb", 3: "Mar", 4: "Apr", 5: "May", 6: "Jun", 7: "Jul", 8: "Aug", 9: "Sep", 10: "Oct", 11: "Nov", 12: "Dec"}
	fmt.Printf("%v, len(months)=%d, cap(months)=%d\n", months, len(months), cap(months))

	// create a slice for q2
	q2 := months[4:7]
	fmt.Printf("%v, len(q2)=%d, cap(q2)=%d\n", q2, len(q2), cap(q2))

	// create a slice for summer
	summer := months[6:10]
	fmt.Printf("%v, len(summer)=%d, cap(summer)=%d\n", summer, len(summer), cap(summer))

	// a slice can be extended beyond its length, but within its capacity
	extendedSummer := summer[:5]
	fmt.Printf("%v, len(extendedSummer)=%d, cap(extendedSummer)=%d\n", extendedSummer, len(extendedSummer), cap(extendedSummer))

	// you can create a slice with make
	s1 := make([]int, 10)
	fmt.Printf("%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))

	s2 := make([]int, 10, 20)
	fmt.Printf("%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))

	// make(T, len, cap) is the same as make(T, cap)[:len]
	s3 := make([]int, 20)[:10]
	fmt.Printf("%v, len(s3)=%d, cap(s3)=%d\n", s3, len(s3), cap(s3))
}