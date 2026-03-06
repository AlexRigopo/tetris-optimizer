package actions

import (
	"strings"
	"testing"
)

func TestLoadTetrominoes_FromAllMarkdownCases(t *testing.T) {
	files := listMarkdownCaseFiles(t)

	for _, name := range files {
		name := name
		t.Run(name, func(t *testing.T) {
			markdown := readMarkdownCase(t, name)
			input, hasInput := extractInputBlock(markdown)

			if name == "edge_case_empty_file.md" {
				input = ""
				hasInput = true
			}
			if !hasInput {
				t.Fatalf("markdown case %s does not contain a parseable input block", name)
			}

			if name == "edge_case_windows_crlf.md" {
				input = strings.ReplaceAll(input, "\n", "\r\n")
			}

			inputPath := writeTempInput(t, input)
			lot, err := LoadTetrominoes(inputPath)

			expectError := strings.HasPrefix(name, "invalid_") || name == "edge_case_empty_file.md"
			if name == "invalid_too_many_tetrominoes.md" {
				// The markdown case expects parser success and label rejection later.
				expectError = false
			}

			if expectError {
				if err == nil {
					t.Fatalf("expected LoadTetrominoes error for %s, got nil", name)
				}
				return
			}

			if err != nil {
				t.Fatalf("expected LoadTetrominoes success for %s, got: %v", name, err)
			}
			if len(lot) == 0 {
				t.Fatalf("expected at least one tetromino for %s", name)
			}

			if name == "edge_case_multiple_blank_lines_between_tetrominoes.md" && len(lot) != 2 {
				t.Fatalf("expected 2 tetrominoes for %s, got %d", name, len(lot))
			}
		})
	}
}
