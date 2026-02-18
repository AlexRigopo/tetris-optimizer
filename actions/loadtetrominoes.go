package actions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LoadTetrominoes(filename string) ([][]rune, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	var (
		all      [][]rune
		block    []string
		blockNum int
	)

	flush := func() error {
		if len(block) == 0 {
			return nil
		}
		blockNum++

		r, err := validateTetrominoes(block, blockNum)
		if err != nil {
			return err
		}

		all = append(all, r)
		block = block[:0]
		return nil
	}

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

	if err := sc.Err(); err != nil {
		return nil, err
	}

	// flush last block if file doesn't end with blank line
	if err := flush(); err != nil {
		return nil, err
	}

	if len(all) == 0 {
		return nil, fmt.Errorf("no tetrominoes found")
	}

	return all, nil
}
