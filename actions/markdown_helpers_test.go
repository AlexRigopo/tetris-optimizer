package actions

import (
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"testing"
)

func listMarkdownCaseFiles(t *testing.T) []string {
	t.Helper()

	entries, err := os.ReadDir(filepath.Join("..", "test_cases"))
	if err != nil {
		t.Fatalf("failed reading test_cases directory: %v", err)
	}

	files := make([]string, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if filepath.Ext(entry.Name()) == ".md" {
			files = append(files, entry.Name())
		}
	}

	sort.Strings(files)
	if len(files) == 0 {
		t.Fatal("no markdown test cases found")
	}
	return files
}

func readMarkdownCase(t *testing.T, name string) string {
	t.Helper()

	fullPath := filepath.Join("..", "test_cases", name)
	raw, err := os.ReadFile(fullPath)
	if err != nil {
		t.Fatalf("failed to read %s: %v", fullPath, err)
	}
	return strings.ReplaceAll(string(raw), "\r\n", "\n")
}

func extractInputBlock(markdown string) (string, bool) {
	re := regexp.MustCompile("(?s)## Input[^\\n]*\\n.*?```(?:txt)?\\n(.*?)\\n```")
	m := re.FindStringSubmatch(markdown)
	if len(m) < 2 {
		return "", false
	}
	return m[1], true
}

func extractExpectedOutputBlock(markdown string) (string, bool) {
	re := regexp.MustCompile("(?s)## Expected Output[^\\n]*\\n```(?:txt)?\\n(.*?)\\n```")
	m := re.FindStringSubmatch(markdown)
	if len(m) < 2 {
		return "", false
	}
	return m[1], true
}

func writeTempInput(t *testing.T, input string) string {
	t.Helper()

	path := filepath.Join(t.TempDir(), "input.txt")
	if err := os.WriteFile(path, []byte(input), 0o600); err != nil {
		t.Fatalf("failed writing temp input: %v", err)
	}
	return path
}

func splitTetrominoBlocks(input string) [][]string {
	normalized := strings.ReplaceAll(input, "\r\n", "\n")
	lines := strings.Split(normalized, "\n")

	var blocks [][]string
	cur := make([]string, 0, 4)
	flush := func() {
		if len(cur) == 0 {
			return
		}
		cp := make([]string, len(cur))
		copy(cp, cur)
		blocks = append(blocks, cp)
		cur = cur[:0]
	}

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			flush()
			continue
		}
		cur = append(cur, line)
	}
	flush()
	return blocks
}

func lotToBlockString(lot [][]rune) string {
	blocks := make([]string, 0, len(lot))
	for _, piece := range lot {
		blocks = append(blocks, strings.TrimSuffix(string(piece), "\n"))
	}
	return strings.Join(blocks, "\n\n")
}

func countRuneInBoard(board []string, target rune) int {
	count := 0
	for _, row := range board {
		for _, ch := range row {
			if ch == target {
				count++
			}
		}
	}
	return count
}
