package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values from the tree to the channel ch
func Walk(t *tree.Tree, ch chan int) {
	walkRecursive(t, ch)
	close(ch)
}

func walkRecursive(t *tree.Tree, ch chan int) {
	if (t == nil) {
		return
	}
	walkRecursive(t.Left, ch)
	ch <- t.Value
	walkRecursive(t.Right, ch)
}

// PrintTreeContents print the elements of a tree using the in-order approach
func PrintTreeContents(t *tree.Tree) {
	fmt.Println("==== printing tree")
	ch := make(chan int)
	go Walk(t, ch)
	for v := range ch {
		fmt.Println(v)
	}
}

// Same determines whether t1 and t2 trees contain the same values
func Same(t1, t2 *tree.Tree) bool {
	var v1, v2 int
	var ok1, ok2 bool
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		v1, ok1 = <- ch1
		v2, ok2 = <- ch2
		if (ok1 == false && ok2 == false) {
			return true
		} else if (ok1 == true && ok2 == true && v1 == v2) {
			continue
		} else {
			return false
		}
	}
}

func main() {
	// Testing the Walk() function
	t1 := tree.New(1)
	PrintTreeContents(t1)

	t2 := tree.New(1)
	PrintTreeContents(t2)

	t3 := tree.New(2)
	PrintTreeContents(t3)

	// Testing equality of trees
	fmt.Printf("is t1 equivalent to t2?: %v\n", Same(t1, t2))
	fmt.Printf("is t2 equivalent to t1?: %v\n", Same(t2, t1))
	fmt.Printf("is t1 equivalent to t3?: %v\n", Same(t1, t3))
	fmt.Printf("is t2 equivalent to t3?: %v\n", Same(t2, t3))
}