package maze

// MaxElems sets the fixed capacity of the stack
const MaxElems = 100

// Elem sets the fixed type of the stack elements
type Elem string

// StringStack is is the representation of a LIFO data structure
type StringStack struct {
	length int
	elems  [MaxElems]Elem
}

type stringStackErr string

func (e stringStackErr) Error() string {
	return string(e)
}

const (
	// ErrStackFull is returned when you try to push an item on stack that is already at its capacity limit
	ErrStackFull = nodeStackErr("stack is full")

	// ErrStackEmpty is returned when you try to pop/peek from an empty stack
	ErrStackEmpty = nodeStackErr("stack is empty")
)

// NewStringStack returns a pointer to an initialized, empty stack
func NewStringStack() *StringStack {
	return &StringStack{}
}

// Empty returns true if the stack is empty, false otherwise
func (s *StringStack) Empty() bool {
	return s.length == 0
}

// Full returns true if the stack is full, false otherwise
func (s *StringStack) Full() bool {
	return s.length == MaxElems
}

// Push inserts an element as the top of the stack
func (s *StringStack) Push(elem Elem) error {
	if s.length == MaxElems {
		return ErrNodeStackFull
	}
	s.elems[s.length] = elem
	s.length++
	return nil
}

// Peek returns the top of the stack without removing it from the stack
func (s *StringStack) Peek() (Elem, error) {
	if s.length == 0 {
		return "", ErrNodeStackEmpty
	}
	return s.elems[s.length-1], nil
}

// Pop returns and remove the element from the top of the stack
func (s *StringStack) Pop() (Elem, error) {
	if s.length == 0 {
		return "", ErrNodeStackEmpty
	}
	top := s.elems[s.length-1]
	s.length--
	return top, nil
}
