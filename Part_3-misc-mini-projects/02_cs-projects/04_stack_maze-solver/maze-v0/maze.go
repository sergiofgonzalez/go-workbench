package main

import (
	"fmt"
	"os"

	"example.com/maze/stack"
)

var maze0 = []string{
	"AB",
	"  ",
}

var maze1 = []string{
	"A ",
	"B ",
}

var maze2 = []string{
	"A ",
	" B",
}

var maze3 = []string{
	" B",
	"A ",
}

var maze4 = []string{
	"A ",
	"  ",
}

var maze5 = []string{
	"AX",
	" B",
}

var maze6 = []string{
	"A ",
	"XB",
}

var maze7 = []string{
	"A X",
	"X  ",
	"  B",
}

var maze8 = []string{
	"A   X",
	"   X ",
	" XX X",
	"    B",
}

var maze9 = []string{
	"A   X",
	"   X ",
	" X  X",
	"    B",
}

var maze10 = []string{
	"A   X",
	"  XX ",
	" X  X",
	"  X B",
}

const (
	start = "A"
	end   = "B"
	wall  = "X"
	empty = " "
)

func mazeTo2D(maze []string) [][]string {
	var matrix [][]string
	for _, s := range maze {
		var row []string
		for _, c := range s {
			row = append(row, string(c))
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func findInitialState(maze [][]string) State {
	for row := range maze {
		for col := range maze {
			if maze[row][col] == start {
				return State{row, col}
			}
		}
	}
	panic("Malformed maze")
}

func testGoal(maze [][]string, state State) bool {
	if maze[state.Row][state.Col] == end {
		return true
	}
	return false
}

func expandNode(node Node, maze [][]string, frontier *Stack, exploredSet *ExploredSet) {
	findTransitionNodes := func(maze [][]string, state State) []Node {
		var transitionNodes []Node
		candidateMoves := []Action{
			{RowMove: -1, ColMove: 0},
			{RowMove: 0, ColMove: 1},
			{RowMove: 1, ColMove: 0},
			{RowMove: 0, ColMove: -1},
		}

		for _, m := range candidateMoves {
			nextRow := state.Row + m.RowMove
			nextCol := state.Col + m.ColMove
			if nextRow >= 0 && nextRow < len(maze) &&
				nextCol >= 0 && nextCol < len(maze[0]) &&
				maze[nextRow][nextCol] != wall {
				transitionNodes = append(
					transitionNodes,
					Node{
						state:  State{nextRow, nextCol},
						parent: &node,
						action: &Action{m.RowMove, m.ColMove},
						cost:   node.cost + 1,
					})
			}
		}
		return transitionNodes
	}

	for _, transitionNode := range findTransitionNodes(maze, node.state) {
		if !exploredSet.contains(transitionNode) {
			frontier.Push(transitionNode)
		}
	}
}

// State tracks the position of the maze we're exploring
type State struct {
	Row int
	Col int
}

func (s State) String() string {
	return fmt.Sprintf("state=(row=%d, col=%d)", s.Row, s.Col)
}

// Action models the movement from one state to another
type Action struct {
	RowMove int
	ColMove int
}

func (a Action) String() string {
	return fmt.Sprintf("action=(rowInc=%d, colInc=%d)", a.RowMove, a.ColMove)
}

// Node holds information about the status of the Search
type Node struct {
	state  State
	parent *Node
	action *Action
	cost   int
}

func (n Node) String() string {
	return fmt.Sprintf("node=(%v; parent=%v, action=%v, cost=%d)", n.state, n.parent, n.action, n.cost)
}

// ExploredSet keeps track of the already visited states to prevent loops
type ExploredSet []State

func (e *ExploredSet) add(n Node) {

	if !e.contains(n) {
		*e = append(*e, n.state)
	}
}

func (e ExploredSet) contains(n Node) bool {
	for _, s := range e {
		if s == n.state {
			return true
		}
	}
	return false
}

// MaxNodesInStack sets the fixed capacity of the stack
const MaxNodesInStack = 100

// Stack is is the representation of a LIFO data structure
type Stack struct {
	length int
	elems  [MaxNodesInStack]Node
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

// NewStack returns a pointer to an initialized, empty stack
func NewStack() *Stack {
	return &Stack{}
}

// Empty returns true if the stack is empty, false otherwise
func (s *Stack) Empty() bool {
	return s.length == 0
}

// Full returns true if the stack is full, false otherwise
func (s *Stack) Full() bool {
	return s.length == MaxNodesInStack
}

// Push inserts an element as the top of the stack
func (s *Stack) Push(elem Node) error {
	if s.length == MaxNodesInStack {
		return ErrStackFull
	}
	s.elems[s.length] = elem
	s.length++
	return nil
}

// Peek returns the top of the stack without removing it from the stack
func (s *Stack) Peek() (Node, error) {
	if s.length == 0 {
		return Node{}, ErrStackEmpty
	}
	return s.elems[s.length-1], nil
}

// Pop returns and remove the element from the top of the stack
func (s *Stack) Pop() (Node, error) {
	if s.length == 0 {
		return Node{}, ErrStackEmpty
	}
	top := s.elems[s.length-1]
	s.length--
	return top, nil
}

func (s *Stack) String() string {
	var lines []string
	for !s.Empty() {
		node, err := s.Pop()
		if err != nil {
			panic(fmt.Sprintf("could not Pop(): %v", err))
		}
		line := fmt.Sprintf("from %v move %v to %v (cost=%d)\n", node.parent.state, node.action, node.state, node.cost)
		lines = append(lines, line)
	}

	result := ""
	for i := len(lines) - 1; i >= 0; i-- {
		result += lines[i]
	}
	return result
}

func printSolution(node Node) {
	reversedMoves := stack.New()

	for node.parent != nil {
		reversedMoves.Push(stack.Elem(fmt.Sprintf("from (row=%d, col=%d), move (%d, %d) to (%d, %d) (cost=%d)", node.parent.state.Row, node.parent.state.Col, node.action.RowMove, node.action.ColMove, node.state.Row, node.state.Col, node.cost)))
		node = *node.parent
	}
	reversedMoves.Push("Solution found!!!")

	for !reversedMoves.Empty() {
		move, err := reversedMoves.Pop()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error found while building the solution")
			os.Exit(1)
		}
		fmt.Println(move)
	}
}

func main() {
	matrix := mazeTo2D(maze10)
	PrettyPrint(matrix)
	initialState := findInitialState(matrix)

	startNode := Node{
		state:  initialState,
		parent: nil,
		action: nil,
		cost:   0,
	}

	frontier := NewStack()
	frontier.Push(startNode)
	exploredSet := &ExploredSet{}

	node := startNode
	var err error
	for {
		if frontier.Empty() {
			break
		}
		node, err = frontier.Pop()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unexpected situation while searching: %v", err)
			os.Exit(1)
		}
		if testGoal(matrix, node.state) {
			break
		}
		exploredSet.add(node)
		expandNode(node, matrix, frontier, exploredSet)
	}

	if testGoal(matrix, node.state) {
		printSolution(node)
	} else {
		fmt.Println("No solution was found!")
	}

}
