package main

import (
	"fmt"
	"os"

	"tetris-optimizer/actions"
)

func main() {
	// This CLI expects exactly one argument: the path to the tetromino file.
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run . <filename>")
		os.Exit(1)
	}

	// Read the filename from the command line.
	filename := os.Args[1]

	// Guard against an empty filename so we fail early with a clear message.
	if filename == "" {
		fmt.Fprintln(os.Stderr, "Error: No file name")
		os.Exit(1)
	}

	// Load and validate all tetromino blocks from disk.
	lot, err := actions.LoadTetrominoes(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file:", err)
		os.Exit(1)
	}

	// Replace each '#' with a unique letter (A, B, C, ...) per piece.
	lot, err = actions.LabelTetrominoes(lot)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error labeling tetrominoes:", err)
		os.Exit(1)
	}

	// Run the packing algorithm and print the smallest valid square solution.
	solution := actions.SolveSmallestSquare(lot)
	for _, row := range solution {
		fmt.Println(row)
	}
}
