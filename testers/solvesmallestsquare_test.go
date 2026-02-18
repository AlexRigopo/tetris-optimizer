package testers

import (
	"testing"

	"tetris-optimizer/actions"
)

func TestSolveSmallestSquare(t *testing.T) {
	// One simple square tetromino -> smallest solution should be 2x2 of 'A'
	lot := [][]rune{
		[]rune("##..\n##..\n....\n....\n"),
	}

	labeled := actions.LabelTetrominoes(lot)
	solution := actions.SolveSmallestSquare(labeled)

	if len(solution) != 2 {
		t.Fatalf("expected board size 2, got %d", len(solution))
	}
	if solution[0] != "AA" || solution[1] != "AA" {
		t.Fatalf("unexpected solution: %#v", solution)
	}
}

func TestNewBoard(t *testing.T) {
	board := actions.NewBoard(5)

	if len(board) != 5 {
		t.Errorf("Expected board size 5, got %d", len(board))
	}
	for _, row := range board {
		if len(row) != 5 {
			t.Errorf("Expected row size 5, got %d", len(row))
		}
	}
}

func TestCanPlace(t *testing.T) {
	board := []string{
		".....",
		".....",
		".....",
		".....",
		".....",
	}
	p := actions.Piece{ID: 'A', Blocks: [4]actions.Point{{0, 0}, {1, 0}, {0, 1}, {1, 1}}, W: 2, H: 2}

	if !actions.CanPlace(board, p, 0, 0) {
		t.Error("Expected to be able to place piece at (0,0), but couldn't")
	}

	if actions.CanPlace(board, p, 4, 4) {
		t.Error("Expected not to be able to place piece at (4,4), but could")
	}
}

func TestPlace(t *testing.T) {
	board := []string{
		".....",
		".....",
		".....",
		".....",
		".....",
	}
	p := actions.Piece{ID: 'A', Blocks: [4]actions.Point{{0, 0}, {1, 0}, {0, 1}, {1, 1}}, W: 2, H: 2}

	actions.Place(&board, p, 0, 0, 'A')

	if board[0] != "AA..." || board[1] != "AA..." {
		t.Error("Expected piece to be placed correctly on the board")
	}
}

func TestParsePiece(t *testing.T) {
	// ParsePiece expects a single []rune that includes '\n' separators,
	// and it assumes there are exactly 4 non-'.' cells.
	lot := []rune("AA..\nAA..\n....\n....\n")

	piece := actions.ParsePiece(lot)

	if piece.ID != 'A' {
		t.Errorf("Expected piece ID 'A', got %c", piece.ID)
	}

	// Expected normalized blocks: (0,0),(1,0),(0,1),(1,1)
	want := map[actions.Point]bool{
		{0, 0}: true,
		{1, 0}: true,
		{0, 1}: true,
		{1, 1}: true,
	}

	for _, b := range piece.Blocks {
		if !want[b] {
			t.Fatalf("unexpected block coordinate: %+v", b)
		}
		delete(want, b)
	}
	if len(want) != 0 {
		t.Fatalf("missing expected blocks: %+v", want)
	}

	if piece.W != 2 || piece.H != 2 {
		t.Fatalf("expected bounding box 2x2, got %dx%d", piece.W, piece.H)
	}
}
