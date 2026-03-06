package actions

import (
	"strings"
	"testing"
)

func TestValidateTetrominoes_FromMarkdownInputs(t *testing.T) {
	files := listMarkdownCaseFiles(t)

	mustFail := map[string]bool{
		"invalid_disconnected_blocks.md":                 true,
		"invalid_invalid_character.md":                   true,
		"invalid_missing_newline_between_tetrominoes.md": true,
		"invalid_only_3_lines.md":                        true,
		"invalid_space_character.md":                     true,
		"invalid_wrong_hash_count.md":                    true,
		"invalid_wrong_line_length.md":                   true,
	}

	for _, name := range files {
		name := name
		t.Run(name, func(t *testing.T) {
			if name == "edge_case_empty_file.md" {
				// Empty file behavior is validated by LoadTetrominoes tests.
				return
			}

			markdown := readMarkdownCase(t, name)
			input, hasInput := extractInputBlock(markdown)
			if !hasInput {
				t.Fatalf("no parseable input block in %s", name)
			}

			if name == "edge_case_windows_crlf.md" {
				input = strings.ReplaceAll(input, "\n", "\r\n")
			}

			blocks := splitTetrominoBlocks(input)
			if len(blocks) == 0 {
				t.Fatalf("expected at least one block in %s", name)
			}

			sawErr := false
			for i, block := range blocks {
				_, err := ValidateTetrominoes(block, i+1)
				if err != nil {
					sawErr = true
					break
				}
			}

			expectFail := mustFail[name]
			if expectFail && !sawErr {
				t.Fatalf("expected ValidateTetrominoes failure for %s", name)
			}
			if !expectFail && sawErr {
				t.Fatalf("expected ValidateTetrominoes success for %s", name)
			}
		})
	}
}
