package maze

import "fmt"

// State tracks the position of the maze we're exploring
type State struct {
	Row int
	Col int
}

func (s State) String() string {
	return fmt.Sprintf("state=(row=%d, col=%d)", s.Row, s.Col)
}
