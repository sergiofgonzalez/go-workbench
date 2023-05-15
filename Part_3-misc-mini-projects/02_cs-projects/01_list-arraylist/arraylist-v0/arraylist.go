package main

import "fmt"


// MaxElems sets the capacity of the list
const MaxElems = 10

// Elem is the type of each of the elements in the list
type Elem string

// List is a sequential list backed by an array of maxElems elements
type List struct {
	length int
	elems [MaxElems]Elem
}

type listErr string

func (e listErr) Error() string {
	return string(e)
}

const (
	// ErrPosOutOfRange is returned when the position is out of the expected boundaries
	ErrPosOutOfRange = listErr("given position is out of range")

	// ErrListFull is returned when you try to insert an item on a full list
	ErrListFull = listErr("list is full")
)

// NewList initializes a new empty list and returns a pointer to it
func NewList() *List {
	return &List{}
}

// Empty returns true if the list is empty
func (l List) Empty() bool {
	if l.length == 0 {
		return true
	}
	return false
}

// Insert inserts the given element in the given position, shifting the elements
// that come after the one inserted.
// If the given positions is less than zero or greater than the number of elements
// if the list an error is returned
func (l *List) Insert(e Elem, pos int) error {
	if pos < 0 || pos > l.length {
		return ErrPosOutOfRange
	}

	if cap(l.elems) == l.length {
		return ErrListFull
	}

	// shift to the right as required
	for i := l.length; i > pos; i-- {
		l.elems[i] = l.elems[i - 1]
	}

	// place the element in the correct position
	l.elems[pos] = e
	l.length++
	return nil
}

// Remove removes the element at the given position, shifting the elements that
// come after the one removed.
// If the given position is less than zero or greater than the position of the
// last element an error is returned
func (l *List) Remove(pos int) error {
	if pos < 0 || pos >= l.length {
		return ErrPosOutOfRange
	}

	// shift to the left as required
	for i := pos; i < l.length - 1; i++ {
		l.elems[i] = l.elems[i + 1]
	}

	// adjust length
	l.length--
	return nil
}

// Len returns the number of elements in the current list
func (l *List) Len() int {
	return l.length
}


// Traverse invokes the given function for all the elements in the list in order
func (l *List) Traverse(fn func (Elem)) {
	for i := 0; i < l.length; i++ {
		fn(l.elems[i])
	}
}

func main() {
	fmt.Println("Run go test")
}