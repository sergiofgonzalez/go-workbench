package main

import (
	"fmt"
	"strings"
)

func main() {
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	fmt.Printf("len(board)=%v\n", len(board))
	fmt.Printf("len(board[0])=%v\n", len(board[0]))

	// Game starts
	printBoard(board)

	// X player's first move
	board[0][0] = "X"
	printBoard(board)

	// O player's first move
	board[2][2] = "O"
	printBoard(board)

	// X player's second move
	board[1][2] = "X"
	printBoard(board)

	// O player's second move
	board[1][0] = "O"
	printBoard(board)

	// X player's third move
	board[0][2] = "X"
	printBoard(board)
}

func printBoard(b [][]string) {
	for i := 0; i < len(b); i++ {
		fmt.Println(strings.Join(b[i], " "))
	}
	fmt.Println("=========\n")
}