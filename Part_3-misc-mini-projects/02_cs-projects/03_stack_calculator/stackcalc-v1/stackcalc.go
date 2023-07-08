package main

import (
	"fmt"
	"os"
	"strconv"

	"example.com/arraystack/stack"
)

func main() {
	s := stack.New[int]()

	// Get elements from the command line
	for _, elem := range os.Args[1:] {
		n, err := strconv.Atoi(elem)
		if err == nil {
			s.Push(n)
			continue
		}

		switch elem {
		case "+":
			n1, n2 := getElemsForBinaryOp(s)
			s.Push(n1 + n2)
		case "-":
			n1, n2 := getElemsForBinaryOp(s)
			s.Push(n1 - n2)
		case "*":
			n1, n2 := getElemsForBinaryOp(s)
			s.Push(n1 * n2)
		case "/":
			n1, n2 := getElemsForBinaryOp(s)
			s.Push(n1 / n2)
		case "neg":
			n := getElem(s)
			s.Push(-n)
		default:
			doError(fmt.Errorf("argument is neither an int or a valid operation: %v", elem))
		}
	}

	result, err := s.Pop()
	if err != nil {
		doError(fmt.Errorf("couldn't get result: %v", err))
	}
	if !s.Empty() {
		fmt.Printf("WARNING: stack was not empty: %v", *s)
	} else {
		fmt.Printf("Result: %d\n", result)
	}
}


func getElemsForBinaryOp(s *stack.Stack[int]) (int, int) {
	n2 := getElem(s)
	n1 := getElem(s)
	return n1, n2
}

func getElem(s *stack.Stack[int]) int {
	elem, err := s.Pop()
	if err != nil {
		doError(err)
	}

	return elem
}

func doError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
