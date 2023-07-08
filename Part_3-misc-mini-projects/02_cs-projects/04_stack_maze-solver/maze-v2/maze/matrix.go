package maze

import "fmt"

// Matrix is the representation of a Maze as a bi-dimensional slice of strings
type Matrix [][]string

// New returns a Matrix representation of a maze from a slice of strings.
func New(maze []string) *Matrix {
	var m Matrix
	for _, s := range maze {
		if len(s) != len(maze[0]) {
			panic(fmt.Sprintf("malformed maze: %d != %d", len(s), len(maze[0])))
		}
		var row []string
		for _, c := range s {
			row = append(row, string(c))
		}
		m = append(m, row)
	}
	return &m
}

func (m Matrix) findInitialState() State {
	for i, row := range m {
		for j := range row {
			if m[i][j] == Start {
				return State{i, j}
			}
		}
	}
	panic("malformed maze: should have initial state")
}

func (m Matrix) testGoal(state State) bool {
	if m[state.Row][state.Col] == End {
		return true
	}
	return false
}

// PrettyPrint prints a maze defined as a 2D matrix of characters
func (m Matrix) PrettyPrint() {
	ppFirstLine(m)
	for row := range m {
		fmt.Print("║")
		for col := range m[row] {
			fmt.Printf(" %s ", m[row][col])
			switch {
			case col == len(m[row])-1:
				fmt.Println("║")
			default:
				fmt.Print("║")
			}
		}
		switch {
		case row == len(m)-1:
			ppLastLine(m)
		default:
			ppSeparatorLine(m)
		}
	}
}

func ppFirstLine(maze [][]string) {
	fmt.Printf("╔")
	for i := 0; i < len(maze[0])-1; i++ {
		fmt.Printf("═══╦")
	}
	fmt.Println("═══╗")
}

func ppSeparatorLine(maze [][]string) {
	fmt.Print("╠")
	for i := 0; i < len(maze[0])-1; i++ {
		fmt.Print("═══╬")
	}
	fmt.Println("═══╣")
}

func ppLastLine(maze [][]string) {
	fmt.Print("╚")
	for i := 0; i < len(maze[0])-1; i++ {
		fmt.Print("═══╩")
	}
	fmt.Println("═══╝")
}