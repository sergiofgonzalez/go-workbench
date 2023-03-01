package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	// No-guardrails!! it could be nil
	fmt.Println(t.S)
}


func main() {
	var i I				// nil interface
	describe(i)
	// i.M()					// -> Run-time error as there's no concrete type on which to call M()

}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}