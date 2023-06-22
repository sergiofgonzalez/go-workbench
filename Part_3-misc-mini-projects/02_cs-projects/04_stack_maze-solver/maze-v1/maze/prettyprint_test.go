package maze_test

import "example.com/maze/maze"

func Example_prettyPrint_empty2x2() {
	m := maze.Matrix([][]string{
		{" ", " "},
		{" ", " "},
	})

	m.PrettyPrint()
	// Output:
	// ╔═══╦═══╗
	// ║   ║   ║
	// ╠═══╬═══╣
	// ║   ║   ║
	// ╚═══╩═══╝
}

func Example_prettyPrint_2x2() {
	m := maze.Matrix{
		{"A", "B"},
		{"X", " "},
	}
	m.PrettyPrint()
	// Output:
	// ╔═══╦═══╗
	// ║ A ║ B ║
	// ╠═══╬═══╣
	// ║ X ║   ║
	// ╚═══╩═══╝
}

func Example_prettyPrint_empty3x3() {
	m := maze.Matrix{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
	m.PrettyPrint()
	// Output:
	// ╔═══╦═══╦═══╗
	// ║   ║   ║   ║
	// ╠═══╬═══╬═══╣
	// ║   ║   ║   ║
	// ╠═══╬═══╬═══╣
	// ║   ║   ║   ║
	// ╚═══╩═══╩═══╝
}

func Example_prettyPrint_3x3() {
	m := maze.Matrix{
		{"A", "B", " "},
		{"X", " ", "X"},
		{"B", "X", "A"},
	}
	m.PrettyPrint()
	// Output:
	// ╔═══╦═══╦═══╗
	// ║ A ║ B ║   ║
	// ╠═══╬═══╬═══╣
	// ║ X ║   ║ X ║
	// ╠═══╬═══╬═══╣
	// ║ B ║ X ║ A ║
	// ╚═══╩═══╩═══╝
}
