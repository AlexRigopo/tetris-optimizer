package actions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// LoadTetrominoes reads tetromino blocks from a file and validates each block.
func LoadTetrominoes(filename string) ([][]rune, error) {
	// Open input file for line-by-line parsing.
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Scanner reads one line at a time from the file.
	sc := bufio.NewScanner(f)

	var (
		all      [][]rune
		block    []string
		blockNum int
	)

	// flush finalizes the current 4-line block (if any), validates it,
	// and appends it to the full tetromino list.
	flush := func() error {
		if len(block) == 0 {
			// Multiple blank lines are tolerated; nothing to flush here.
			return nil
		}
		blockNum++

		// Validate shape/format and convert to rune representation.
		r, err := ValidateTetrominoes(block, blockNum)
		if err != nil {
			return err
		}

		// Keep valid piece, then reset the accumulator for the next one.
		all = append(all, r)
		block = block[:0]
		return nil
	}

	// Build tetromino blocks while scanning file lines.
	for sc.Scan() {
		line := strings.TrimRight(sc.Text(), "\r") // handle Windows CRLF too

		// blank line => end of current tetromino block
		if strings.TrimSpace(line) == "" {
			if err := flush(); err != nil {
				return nil, err
			}
			continue
		}

		block = append(block, line)
	}

	// Propagate low-level scanner failures (I/O, tokenization limits, etc.).
	if err := sc.Err(); err != nil {
		return nil, err
	}

	// flush last block if file doesn't end with blank line
	if err := flush(); err != nil {
		return nil, err
	}

	// Reject empty files (or files with only separators).
	if len(all) == 0 {
		return nil, fmt.Errorf("no tetrominoes found")
	}

	return all, nil
}
