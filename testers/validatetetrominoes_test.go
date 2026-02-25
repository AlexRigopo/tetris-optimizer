package testers

import (
	"testing"

	"tetris-optimizer/actions"
)

func TestValidateTetrominoes(t *testing.T) {
	validLines := []string{
		"....",
		".##.",
		".##.",
		"....",
	}
	_, err := actions.ValidateTetrominoes(validLines, 1)
	if err != nil {
		t.Errorf("Expected no error for valid tetromino, got %v", err)
	}

	invalidLines := []string{
		"....",
		".#..",
		".##.",
		"....",
	}
	_, err = actions.ValidateTetrominoes(invalidLines, 2)
	if err == nil {
		t.Error("Expected an error for invalid tetromino, got none")
	}
}
