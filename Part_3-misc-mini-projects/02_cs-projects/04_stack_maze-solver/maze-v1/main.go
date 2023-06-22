package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"example.com/maze/maze"
)

var mazes = map[string][]string{
	"maze0": {
		"AB",
		"  ",
	},
	"maze1": {
		"A ",
		"B ",
	},
	"maze2": {
		"A ",
		" B",
	},
	"maze3": {
		" B",
		"A ",
	},
	"maze4": {
		"A ",
		"  ",
	},
	"maze5": {
		"AX",
		" B",
	},
	"maze6": {
		"A ",
		"XB",
	},
	"maze7": {
		"A X",
		"X  ",
		"  B",
	},
	"maze8": {
		"A   X",
		"   X ",
		" XX X",
		"    B",
	},
	"maze9": {
		"A   X",
		"   X ",
		" X  X",
		"    B",
	},
	"maze10": {
		"A   X",
		"  XX ",
		" X  X",
		"  X B",
	},
	"maze11": {
		"A  B",
	},
	"maze12": {
		"A  X B",
		"X     ",
	},
}

// Command-line flags
var (
	filePath = flag.String("file", "", "file representing a maze")
	mazeNum = flag.Int("num", 9, "built-in maze number (0-12)")
)

func main() {

	// Parse the command-line args as flags
	flag.Parse()

	var mLines []string

	mLines = readMazeFromFile("./samples/sample.txt")

	switch {
	case *filePath != "":
		mLines = readMazeFromFile(*filePath)
	default:
		fmt.Printf("Using built-in maze #%d\n", *mazeNum)
		mLines = mazes[fmt.Sprintf("maze%d", *mazeNum)]
		if mLines == nil {
			fmt.Fprintf(os.Stderr, "No such built-in maze index (0-12): %d", *mazeNum)
			os.Exit(1)
		}
	}

	matrix := maze.New(mLines)
	matrix.PrettyPrint()

	solutionMoves := matrix.Solve()
	printSolution(solutionMoves)
}

func readMazeFromFile(path string) []string {
	bytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open %q for reading: %v", path, err)
		os.Exit(1)
	}

	if len(bytes) == 0 {
		fmt.Fprintf(os.Stderr, "file %q is empty", path)
		os.Exit(1)
	}

	// splits a string by the blanks
	lines := strings.Split(string(bytes), "\n")

	// is the last line empty?
	if len(lines[len(lines) - 1]) == 0 {
		return lines[:len(lines) - 1]
	}
	return lines
}


func printSolution(moves []string) {
	if moves != nil {
		for _, m := range moves {
			fmt.Println(m)
		}
	} else {
		fmt.Println("No solution was found!")
	}
}