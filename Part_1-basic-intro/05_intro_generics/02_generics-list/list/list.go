package main

import (
	"fmt"
	"strings"
)

type List[T any] struct {
	next *List[T]
	val 	T
}

func (l *List[T]) add(elem T) {
	n := List[T]{nil, elem}
	var p *List[T]
	for p = l; p.next != nil; p = p.next {
	}
	p.next = &n
}

func (l List[T]) String() string {
	var s []string
	for p := &l; ; p = p.next {
		s = append(s, fmt.Sprint(p.val))
		if p.next == nil {
			break
		}
	}
	return fmt.Sprintf("[%v]", strings.Join(s, ", "))
}

func main() {
	l := List[int]{nil, 0}
	fmt.Println(l)

	l.add(1)
	l.add(2)
	l.add(3)
	fmt.Println(l)
}