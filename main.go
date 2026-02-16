package main

import (
	"fmt"
	"os"

	"tetris-optimizer/actions"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run . <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]

	// Empty name for file → Print error
	if filename == "" {
		fmt.Fprintln(os.Stderr, "Error: No file name")
		os.Exit(1)
	}

	// lot = list of tetraminoes
	lot, err := actions.LoadTetrominoes(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file:", err)
		os.Exit(1)
	}
}
