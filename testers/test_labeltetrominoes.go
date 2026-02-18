package testers

import (
    "testing"
    "tetris-optimizer/actions"
)

func TestLabelTetrominoes(t *testing.T) {
    lot := [][]rune{
        {'#', '#', '#', '#'},
        {'#', '.', '.', '.'},
    }
    labeled := actions.LabelTetrominoes(lot)

    // Check if the tetromino is labeled correctly
    for _, row := range labeled {
        for _, ch := range row {
            if ch != '.' && ch != 'A' {
                t.Errorf("Expected 'A' or '.', got %c", ch)
            }
        }
    }
}
