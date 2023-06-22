package maze

import "fmt"

// Action models the movement from one state to another
type Action struct {
	RowMove int
	ColMove int
}

func (a Action) String() string {
	return fmt.Sprintf("action=(rowInc=%d, colInc=%d)", a.RowMove, a.ColMove)
}