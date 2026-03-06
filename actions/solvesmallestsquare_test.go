package actions

import (
	"os"
	"strings"
	"testing"
)

func TestSolveSmallestSquare_ValidMarkdownCases(t *testing.T) {
	cases := []string{
		"valid_single_tetromino.md",
		"valid_multiple_tetrominoes.md",
		"edge_case_trailing_newline_at_end.md",
		"edge_case_multiple_blank_lines_between_tetrominoes.md",
		"edge_case_windows_crlf.md",
	}

	for _, name := range cases {
		name := name
		t.Run(name, func(t *testing.T) {
			markdown := readMarkdownCase(t, name)
			input, hasInput := extractInputBlock(markdown)
			if !hasInput {
				t.Fatalf("no input block in %s", name)
			}
			if name == "edge_case_windows_crlf.md" {
				input = strings.ReplaceAll(input, "\n", "\r\n")
			}

			lot, err := LoadTetrominoes(writeTempInput(t, input))
			if err != nil {
				t.Fatalf("load failed for %s: %v", name, err)
			}

			labeled, err := LabelTetrominoes(lot)
			if err != nil {
				t.Fatalf("label failed for %s: %v", name, err)
			}
			board := SolveSmallestSquare(labeled)

			n := len(board)
			if n == 0 {
				t.Fatalf("solver returned empty board for %s", name)
			}
			for _, row := range board {
				if len(row) != n {
					t.Fatalf("board is not square for %s: row length %d, board %d", name, len(row), n)
				}
			}

			for i := range labeled {
				letter := rune('A' + i)
				if countRuneInBoard(board, letter) != 4 {
					t.Fatalf("expected exactly 4 cells for %c in %s", letter, name)
				}
			}

			if name == "valid_single_tetromino.md" {
				if n != 3 {
					t.Fatalf("expected minimal 3x3 board for single L-piece, got %dx%d", n, n)
				}
			}
		})
	}
}

func TestSolveSmallestSquare_MaxCase_Optional(t *testing.T) {
	if os.Getenv("RUN_SLOW_TESTS") != "1" {
		t.Skip("set RUN_SLOW_TESTS=1 to run the 26-piece solver stress test")
	}

	markdown := readMarkdownCase(t, "edge_case_max_tetrominoes.md")
	input, hasInput := extractInputBlock(markdown)
	if !hasInput {
		t.Fatal("no input block in edge_case_max_tetrominoes.md")
	}

	lot, err := LoadTetrominoes(writeTempInput(t, input))
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	if len(lot) != 26 {
		t.Fatalf("expected 26 pieces, got %d", len(lot))
	}

	labeled, err := LabelTetrominoes(lot)
	if err != nil {
		t.Fatalf("label failed: %v", err)
	}

	board := SolveSmallestSquare(labeled)
	if len(board) == 0 {
		t.Fatal("solver returned empty board")
	}
}
