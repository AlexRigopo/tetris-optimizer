package actions

import "fmt"

func ValidateTetrominoes(lines []string, blockNum int) ([]rune, error) {
	// Each tetromino must be exactly 4 lines tall.
	if len(lines) != 4 {
		return nil, fmt.Errorf("ERROR")
	}

	hashCount := 0
	var out []rune

	// Validate each row and translate it to our rune format.
	for _, line := range lines {
		// Each row must be exactly 4 columns wide.
		if len(line) != 4 {
			return nil, fmt.Errorf("ERROR")
		}

		// Only '.' and '#' are legal grid characters.
		for _, ch := range line {
			if ch != '#' && ch != '.' {
				return nil, fmt.Errorf("ERROR")
			}
			// Count occupied cells to enforce tetromino size = 4 blocks.
			if ch == '#' {
				hashCount++
			}
			out = append(out, ch)
		}
		// Preserve row boundaries for later parsing stages.
		out = append(out, '\n') // keep line breaks inside the rune slice
	}

	// A valid tetromino always contains exactly 4 occupied cells.
	if hashCount != 4 {
		return nil, fmt.Errorf("ERROR")
	}

	// Ensure all '#' cells belong to one connected component.
	if !isConnected(lines) {
		return nil, fmt.Errorf("ERROR")
	}

	return out, nil
}

func isConnected(lines []string) bool {
	type point struct{ x, y int }

	// Collect coordinates of every occupied cell.
	var blocks []point
	for y, line := range lines {
		for x, ch := range line {
			if ch == '#' {
				blocks = append(blocks, point{x: x, y: y})
			}
		}
	}

	if len(blocks) == 0 {
		return false
	}

	// Build a quick lookup set for adjacency checks.
	occupied := make(map[point]bool, len(blocks))
	for _, b := range blocks {
		occupied[b] = true
	}

	// Breadth-first search from the first occupied cell.
	visited := make(map[point]bool, len(blocks))
	queue := []point{blocks[0]}
	visited[blocks[0]] = true

	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		// Visit all 4-neighbor adjacent occupied cells.
		for _, d := range dirs {
			next := point{x: cur.x + d[0], y: cur.y + d[1]}
			if occupied[next] && !visited[next] {
				visited[next] = true
				queue = append(queue, next)
			}
		}
	}

	// Connected means BFS reached all occupied cells.
	return len(visited) == len(blocks)
}
