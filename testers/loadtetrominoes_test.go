package testers

import (
	"testing"

	"tetris-optimizer/actions"
)

func TestLoadTetrominoes(t *testing.T) {
	// Test loading valid tetrominoes
	_, err := actions.LoadTetrominoes("valid_tetrominoes.txt")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Test loading from a non-existent file
	_, err = actions.LoadTetrominoes("non_existent_file.txt")
	if err == nil {
		t.Error("Expected an error for non-existent file, got none")
	}
}
