package actions

import "fmt"

func validateAndToRunes(lines []string, blockNum int) ([]rune, error) {
	if len(lines) != 4 {
		return nil, fmt.Errorf("tetromino %d: expected 4 lines, got %d", blockNum, len(lines))
	}

	hashCount := 0
	var out []rune

	for i, line := range lines {
		if len(line) != 4 {
			return nil, fmt.Errorf("tetromino %d: line %d must be 4 chars, got %d", blockNum, i+1, len(line))
		}

		for _, ch := range line {
			if ch != '#' && ch != '.' {
				return nil, fmt.Errorf("tetromino %d: invalid character %q", blockNum, ch)
			}
			if ch == '#' {
				hashCount++
			}
			out = append(out, ch)
		}
		out = append(out, '\n') // keep line breaks inside the rune slice
	}

	if hashCount != 4 {
		return nil, fmt.Errorf("tetromino %d: expected exactly 4 '#', got %d", blockNum, hashCount)
	}

	return out, nil
}
