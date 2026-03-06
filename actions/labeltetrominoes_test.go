package actions

import (
	"strings"
	"testing"
)

func TestLabelTetrominoes_MarkdownExpectedOutputBlocks(t *testing.T) {
	cases := []string{
		"valid_single_tetromino.md",
		"valid_multiple_tetrominoes.md",
		"edge_case_trailing_newline_at_end.md",
	}

	for _, name := range cases {
		name := name
		t.Run(name, func(t *testing.T) {
			markdown := readMarkdownCase(t, name)
			input, hasInput := extractInputBlock(markdown)
			if !hasInput {
				t.Fatalf("no input block in %s", name)
			}

			expected, hasExpected := extractExpectedOutputBlock(markdown)
			if !hasExpected {
				t.Fatalf("no expected output block in %s", name)
			}

			lot, err := LoadTetrominoes(writeTempInput(t, input))
			if err != nil {
				t.Fatalf("load failed for %s: %v", name, err)
			}

			labeled, err := LabelTetrominoes(lot)
			if err != nil {
				t.Fatalf("label failed for %s: %v", name, err)
			}
			got := lotToBlockString(labeled)
			if strings.ReplaceAll(got, "\r\n", "\n") != strings.ReplaceAll(expected, "\r\n", "\n") {
				t.Fatalf("unexpected labeled blocks for %s\nexpected:\n%s\n\ngot:\n%s", name, expected, got)
			}
		})
	}
}

func TestLabelTetrominoes_MaxBoundaryAtoZ(t *testing.T) {
	markdown := readMarkdownCase(t, "edge_case_max_tetrominoes.md")
	input, ok := extractInputBlock(markdown)
	if !ok {
		t.Fatal("edge_case_max_tetrominoes.md is missing an input block")
	}

	lot, err := LoadTetrominoes(writeTempInput(t, input))
	if err != nil {
		t.Fatalf("expected parser success for max boundary case: %v", err)
	}
	if len(lot) != 26 {
		t.Fatalf("expected 26 tetrominoes, got %d", len(lot))
	}

	labeled, err := LabelTetrominoes(lot)
	if err != nil {
		t.Fatalf("expected label success for 26 tetrominoes, got %v", err)
	}
	for i := 0; i < 26; i++ {
		letter := rune('A' + i)
		found := false
		for _, piece := range labeled {
			for _, ch := range piece {
				if ch == letter {
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			t.Fatalf("expected to find label %c in labeled output", letter)
		}
	}
}

func TestLabelTetrominoes_TooManyPiecesShouldNotPanic(t *testing.T) {
	markdown := readMarkdownCase(t, "invalid_too_many_tetrominoes.md")
	input, ok := extractInputBlock(markdown)
	if !ok {
		t.Fatal("invalid_too_many_tetrominoes.md is missing an input block")
	}

	lot, err := LoadTetrominoes(writeTempInput(t, input))
	if err != nil {
		t.Fatalf("expected parser success before label limit check: %v", err)
	}
	if len(lot) <= 26 {
		t.Fatalf("invalid_too_many_tetrominoes.md should contain >26 pieces, got %d", len(lot))
	}

	_, err = LabelTetrominoes(lot)
	if err == nil {
		t.Fatal("expected graceful error for >26 pieces, got nil")
	}
}
