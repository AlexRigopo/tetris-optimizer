package testers

import (
	"testing"

	"tetris-optimizer/actions"
)

func TestLabelTetrominoes(t *testing.T) {
	// Two valid tetrominoes in the same format as the program uses:
	// 4 lines of 4 + '\n', containing exactly 4 '#'.
	t1 := []rune("##..\n##..\n....\n....\n") // square
	t2 := []rune("####\n....\n....\n....\n") // line

	lot := [][]rune{t1, t2}
	labeled := actions.LabelTetrominoes(lot)

	// First piece should contain only '.' and 'A' (and '\n')
	for _, ch := range labeled[0] {
		if ch != '.' && ch != 'A' && ch != '\n' {
			t.Fatalf("piece 1: expected 'A', '.' or '\\n', got %q", ch)
		}
	}

	// Second piece should contain only '.' and 'B' (and '\n')
	for _, ch := range labeled[1] {
		if ch != '.' && ch != 'B' && ch != '\n' {
			t.Fatalf("piece 2: expected 'B', '.' or '\\n', got %q", ch)
		}
	}
}
