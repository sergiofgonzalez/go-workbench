package main

import (
	"fmt"
	"sort"
)

type sequence []int


// sequence implement the sort.Interface
func (seq sequence) Len() int {
	return len(seq)
}

func (seq sequence) Swap(i, j int) {
	seq[i], seq[j] = seq[j], seq[i]
}

func (seq sequence) Less(i, j int) bool {
	return seq[i] < seq[j]
}

// Copy returns a copy of the sequence
func (seq sequence) Copy() sequence {
	// Contrived solution: use a loop
	// res := make(sequence, len(seq), cap(seq))
	// for i := 0; i < len(seq); i++ {
	// 	res[i] = seq[i]
	// }
	// return res

	// Succinct (and more idiomatic) solution: use append
	res := make(sequence, 0, cap(seq))
	return append(res, seq...)
}

// sequence also implements the stringer interface
func (seq sequence) String() string {
	aux := seq.Copy()
	sort.Sort(aux)

	// Easy solution: use conversion
	// return fmt.Sprint([]int(aux))

	// Contrived solution: use a loop
	str := "["
	for i, elem := range aux {
		if i > 0 {
			str += " "
		}
		str += fmt.Sprint(elem)
	}
	return str + "]"
}


func main() {
	seq := sequence{4, 1, 2, 5, 3}
	fmt.Println(seq)
}