package main

import "fmt"

func main() {
	// composite literal
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("Before DoubleMutate: %v\n", s)
	Mutate(s)
	fmt.Printf("After DoubleMutate: %v\n", s) // Item modification works fine

	fmt.Printf("\nBefore DoubleMutateReassign: %v\n", s)
	MutateReassign(s)
	fmt.Printf("After DoubleMutateReassign: %v\n", s) // Item modification works fine

	fmt.Printf("\nBefore Reassign: %v\n", s)
	s = Reassign(s)
	fmt.Printf("After Reassign: %v\n", s) // Item modification works fine

}

// Mutate mutates the elements of the slice by doubling them
func Mutate(s []int) {
	for i := 0; i < len(s); i++ {
		s[i] *= 2
	}
}

// MutateReassign creates a new slice, reassigns it in the function and then returns it
func MutateReassign(s []int) {
	sNew := make([]int, len(s))
	for i := 0; i < len(sNew); i++ {
		sNew[i] = s[i] + 1
	}
	s = sNew // Reassignment
	fmt.Printf(" >> sNew: %v\n", sNew)
}

// Reassign creates a new slice, and then returns it
func Reassign(s []int) []int {
	sNew := make([]int, len(s))
	for i := 0; i < len(sNew); i++ {
		sNew[i] = s[i] + 1
	}
	return sNew // Reassignment
}
