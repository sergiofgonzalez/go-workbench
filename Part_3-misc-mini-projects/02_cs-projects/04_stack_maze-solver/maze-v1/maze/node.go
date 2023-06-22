package maze

import "fmt"

// Node holds information about the status of the Search
type Node struct {
	State  State
	Parent *Node
	Action *Action
	Cost   int
}

func (n Node) String() string {
	return fmt.Sprintf("node=(%v; parent=%v, action=%v, cost=%d)", n.State, n.Parent, n.Action, n.Cost)
}