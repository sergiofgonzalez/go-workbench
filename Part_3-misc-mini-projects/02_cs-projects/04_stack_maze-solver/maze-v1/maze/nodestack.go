package maze

import "fmt"

// MaxNodesInStack sets the fixed capacity of the stack
const MaxNodesInStack = 100

// NodeStack is is the representation of a LIFO data structure
type NodeStack struct {
	length int
	elems  [MaxNodesInStack]Node
}

type nodeStackErr string

func (e nodeStackErr) Error() string {
	return string(e)
}

const (
	// ErrNodeStackFull is returned when you try to push an item on stack that is already at its capacity limit
	ErrNodeStackFull = nodeStackErr("stack is full")

	// ErrNodeStackEmpty is returned when you try to pop/peek from an empty stack
	ErrNodeStackEmpty = nodeStackErr("stack is empty")
)

// NewNodeStack returns a pointer to an initialized, empty stack
func NewNodeStack() *NodeStack {
	return &NodeStack{}
}

// Empty returns true if the stack is empty, false otherwise
func (s *NodeStack) Empty() bool {
	return s.length == 0
}

// Full returns true if the stack is full, false otherwise
func (s *NodeStack) Full() bool {
	return s.length == MaxNodesInStack
}

// Push inserts an element as the top of the stack
func (s *NodeStack) Push(elem Node) error {
	if s.length == MaxNodesInStack {
		return ErrNodeStackFull
	}
	s.elems[s.length] = elem
	s.length++
	return nil
}

// Peek returns the top of the stack without removing it from the stack
func (s *NodeStack) Peek() (Node, error) {
	if s.length == 0 {
		return Node{}, ErrNodeStackEmpty
	}
	return s.elems[s.length-1], nil
}

// Pop returns and remove the element from the top of the stack
func (s *NodeStack) Pop() (Node, error) {
	if s.length == 0 {
		return Node{}, ErrNodeStackEmpty
	}
	top := s.elems[s.length-1]
	s.length--
	return top, nil
}

func (s *NodeStack) String() string {
	var lines []string
	for !s.Empty() {
		node, err := s.Pop()
		if err != nil {
			panic(fmt.Sprintf("could not Pop(): %v", err))
		}
		line := fmt.Sprintf("from %v move %v to %v (cost=%d)\n", node.Parent.State, node.Action, node.State, node.Cost)
		lines = append(lines, line)
	}

	result := ""
	for i := len(lines) - 1; i >= 0; i-- {
		result += lines[i]
	}
	return result
}
