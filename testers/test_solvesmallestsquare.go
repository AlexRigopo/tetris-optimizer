package testers

import (
    "testing"
    "tetris-optimizer/actions"
)

func TestSolveSmallestSquare(t *testing.T) {
    // Create a test case with labeled tetrominoes
    labeled := [][]rune{
        {'A', 'A', 'A', 'A'},
        {'A', '.', '.', '.'},
    }

    solution := actions.SolveSmallestSquare(labeled)

    // Check if the solution is valid (this will depend on your expected output)
    if len(solution) == 0 {
        t.Error("Expected a solution, got none")
    }
}

func TestNewBoard(t *testing.T) {
    board := actions.NewBoard(5) // Ensure NewBoard is exported

    // Check if the board is initialized correctly
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

    // Check if the piece was placed correctly
    if board[0] != "AA..." || board[1] != "AA..." {
        t.Error("Expected piece to be placed incorrectly on the board")
    }
}

func TestParsePiece(t *testing.T) {
    lot := []string{
        "....",
        ".##.",
        ".##.",
        "....",
    }
    piece := actions.ParsePiece(lot)

    // Check if the piece has the correct ID and blocks
    if piece.ID != 'A' {
        t.Errorf("Expected piece ID 'A', got %c", piece.ID)
    }
    if len(piece.Blocks) != 4 {
        t.Errorf("Expected 4 blocks, got %d", len(piece.Blocks))
    }
}