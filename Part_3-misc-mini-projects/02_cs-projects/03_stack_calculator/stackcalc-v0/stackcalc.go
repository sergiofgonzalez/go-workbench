package main

import (
	"fmt"
	"os"
	"strconv"

	"example.com/arraystack/stack"
)

func main() {
	s := stack.New()

	// Get elements from the command line
	for _, elem := range os.Args[1:] {
		_, err := strconv.Atoi(elem)
		if err == nil {
			s.Push(stack.Elem(elem))
			continue
		}

		switch elem {
		case "+":
			n1, n2 := getElemsForBinaryOp(s)
			s.Push(stack.Elem(strconv.Itoa(n1 + n2)))
		case "-":
			n1, n2 := getElemsForBinaryOp(s)
			s.Push(stack.Elem(strconv.Itoa(n1 - n2)))
		case "*":
			n1, n2 := getElemsForBinaryOp(s)
			s.Push(stack.Elem(strconv.Itoa(n1 * n2)))
		case "/":
			n1, n2 := getElemsForBinaryOp(s)
			s.Push(stack.Elem(strconv.Itoa(n1 / n2)))
		case "neg":
			n := getElemAsInt(s)
			s.Push(stack.Elem(strconv.Itoa(n)))
		default:
			doError(fmt.Errorf("argument is neither an int or a valid operation: %q", elem))
		}
	}

	result, err := s.Pop()
	if err != nil {
		doError(fmt.Errorf("couldn't get result: %v", err))
	}
	if !s.Empty() {
		fmt.Printf("WARNING: stack was not empty: %v", *s)
	} else {
		fmt.Printf("Result: %s\n", result)
	}
}


func getElemsForBinaryOp(s *stack.Stack) (int, int) {
	n2 := getElemAsInt(s)
	n1 := getElemAsInt(s)
	return n1, n2
}

func getElemAsInt(s *stack.Stack) int {
	elem, err := s.Pop()
	if err != nil {
		doError(err)
	}

	n, err := strconv.Atoi(string(elem))
	if err != nil {
		doError(err)
	}
	return n
}


func doError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
