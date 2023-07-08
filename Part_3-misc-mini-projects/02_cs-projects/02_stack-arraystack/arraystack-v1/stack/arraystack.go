package stack

// MaxElems sets the fixed capacity of the stack
const MaxElems = 100

// Stack is is the representation of a LIFO data structure
type Stack[T any] struct {
	length int
	elems  [MaxElems]T
}

type stackErr string

func (e stackErr) Error() string {
	return string(e)
}

const (
	// ErrStackFull is returned when you try to push an item on stack that is already at its capacity limit
	ErrStackFull = stackErr("stack is full")

	// ErrStackEmpty is returned when you try to pop/peek from an empty stack
	ErrStackEmpty = stackErr("stack is empty")
)

// New returns a pointer to an initialized, empty stack
func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Empty returns true if the stack is empty, false otherwise
func (s *Stack[T]) Empty() bool {
	return s.length == 0
}

// Full returns true if the stack is full, false otherwise
func (s *Stack[T]) Full() bool {
	return s.length == MaxElems
}

// Push inserts an element as the top of the stack
func (s *Stack[T]) Push(elem T) error {
	if s.length == MaxElems {
		return ErrStackFull
	}
	s.elems[s.length] = elem
	s.length++
	return nil
}

// Peek returns the top of the stack without removing it from the stack
func (s *Stack[T]) Peek() (t T, err error) {
	if s.length == 0 {
		return t, ErrStackEmpty // t used for its zero-value
	}
	return s.elems[s.length-1], nil
}

// Pop returns and remove the element from the top of the stack
func (s *Stack[T]) Pop() (t T, err error) {
	if s.length == 0 {
		return t, ErrStackEmpty // t used for its zero-value
	}
	top := s.elems[s.length-1]
	s.length--
	return top, nil
}
