package main

import (
	"fmt"
	"strings"
)

// TextScreen is a text screen modeled as a 2D slice
type TextScreen [][]byte


const (
	// Rows is the number of rows in the CRT screen
	Rows = 25
	// Cols is the number of cols in the CRT screen
 	Cols = 80
)


func (scr TextScreen) String() string {
	sb := strings.Builder{}
	sb.WriteString("===\n")
	for i := range scr {
		sb.WriteString(string(scr[i]))
		sb.WriteString("\n")
	}
	sb.WriteString("===\n")
	return sb.String()
}

func main() {

	// Allocation option #1: allocating each row independently
	var scr TextScreen = make([][]byte, Rows)
	for row := range scr {
		scr[row] = make([]byte, Cols)
	}

	lines := []string{"Line1", "Line2", "Line3"}

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			scr[i][j] = lines[i][j]
		}
	}


	fmt.Println(scr)

	// Allocation option #2: using a single allocation, then readjusting
	var scr2 TextScreen = make([][]byte, Rows)

	// Allocate all the screen bytes in one shot
	scrBytes := make([]byte, Rows*Cols)

	// Readjust row by row
	// Note that we're using an idiomatic way:
	//  assigning each row up to cols
	//  diminish the amount of storage by cols on each iteration
	for row := range scr2 {
		scr2[row], scrBytes = scrBytes[:Cols], scrBytes[Cols:]
	}

	lines = []string{"Another first line", "Another second line", "Another third line"}

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			scr2[i][j] = lines[i][j]
		}
	}

	fmt.Println(scr2)
}
