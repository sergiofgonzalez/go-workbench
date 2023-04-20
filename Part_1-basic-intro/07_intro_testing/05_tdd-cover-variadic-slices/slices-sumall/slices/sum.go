package slices

// Sum adds all the numbers in the given slice and return the result
func Sum(s []int) (result int) {
	for _, n := range s {
		result += n
	}
	return
}


// SumAll sums the numbers of all the slices received and return a slice
func SumAll(slices ...[]int) []int {
	sums := make([]int, len(slices))
	for i, s := range slices {
		sums[i] = Sum(s)
	}
	return sums
}