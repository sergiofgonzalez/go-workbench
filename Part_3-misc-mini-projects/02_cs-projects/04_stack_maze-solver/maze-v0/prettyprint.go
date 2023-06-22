package main

import "fmt"

// PrettyPrint prints a maze defined as a 2D matrix of characters
func PrettyPrint(maze [][]string) {
	ppFirstLine(maze)
	for row := range maze {
		fmt.Print("║")
		for col := range maze[row] {
			fmt.Printf(" %s ", maze[row][col])
			switch {
			case col == len(maze[row]) - 1:
				fmt.Println("║")
			default:
				fmt.Print("║")
			}
		}
		switch {
		case row == len(maze) - 1:
			ppLastLine(maze)
		default:
			ppSeparatorLine(maze)
		}
	}
}

func ppFirstLine(maze [][]string) {
	fmt.Printf("╔")
	for i := 0; i < len(maze[0]) - 1; i++ {
		fmt.Printf("═══╦")
	}
	fmt.Println("═══╗")
}

func ppSeparatorLine(maze [][]string) {
	fmt.Print("╠")
	for i := 0; i < len(maze[0]) - 1; i++ {
		fmt.Print("═══╬")
	}
	fmt.Println("═══╣")
}

func ppLastLine(maze [][]string) {
	fmt.Print("╚")
	for i := 0; i < len(maze[0]) - 1; i++ {
		fmt.Print("═══╩")
	}
	fmt.Println("═══╝")
}