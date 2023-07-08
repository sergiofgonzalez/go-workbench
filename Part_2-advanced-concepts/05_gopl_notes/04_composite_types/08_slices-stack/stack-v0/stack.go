package main

import "fmt"

type stack []string

func (s *stack) push(v string) {
	*s = append(*s, v)
}

func (s *stack) pop() string {
	v := (*s)[len(*s) - 1]
	*s = (*s)[:len(*s) - 1]
	return v
}

func (s *stack) top() string {
	return (*s)[len(*s)-1]
}

func (s *stack) remove(i int) {
	copy((*s)[i:], (*s)[i+1:])
	*s = (*s)[:len(*s)-1]
}


func main() {
	s := stack{}

	s.push("a")
	fmt.Println(s)

	s.push("b")
	fmt.Println(s)

	elem := s.pop()
	fmt.Println(s)
	fmt.Println(elem)

	fmt.Println(s.top())

	s.push("b")
	s.push("x")
	s.push("c")
	fmt.Println(s)
	s.remove(2)
	fmt.Println(s)
}