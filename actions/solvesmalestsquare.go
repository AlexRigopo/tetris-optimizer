package actions

import "math"

type Point struct{ X, Y int }

type Piece struct {
	ID     rune     // 'A', 'B', ...
	Blocks [4]Point // 4 occupied cells relative to origin
	W, H   int      // bounding box width/height (optional but speeds up)
}

func SolveSmallestSquare(labeled [][]rune) []string {
	// Convert labeled rune grids into compact piece structures for fast placement checks.
	pieces := make([]Piece, 0, len(labeled))
	for _, lot := range labeled {
		pieces = append(pieces, ParsePiece(lot))
	}

	// Start from the smallest possible square that can hold all cells by area.
	size := int(math.Ceil(math.Sqrt(float64(len(pieces) * 4))))
	for {
		// Try to solve at current size; increase size only if impossible.
		board := NewBoard(size)
		if Backtrack(board, pieces, 0) {
			return board
		}
		size++
	}
}

func NewBoard(n int) []string {
	// Build an n x n board initialized with '.'.
	b := make([][]rune, n)
	for i := range b {
		row := make([]rune, n)
		for j := range row {
			row[j] = '.'
		}
		b[i] = row
	}
	// Convert rune rows into immutable strings used by the rest of the solver.
	out := make([]string, n)
	for i := range b {
		out[i] = string(b[i])
	}
	return out
}

func Backtrack(board []string, pieces []Piece, idx int) bool {
	// Base case: every piece has been placed successfully.
	if idx == len(pieces) {
		return true
	}

	p := pieces[idx]
	n := len(board)

	// Try positions (y, x). Using bounding box reduces useless checks.
	for y := 0; y <= n-p.H; y++ {
		for x := 0; x <= n-p.W; x++ {
			// If piece fits, place it and recurse for the next piece.
			if CanPlace(board, p, x, y) {
				Place(&board, p, x, y, p.ID)
				if Backtrack(board, pieces, idx+1) {
					return true
				}
				// Dead end: undo placement and continue searching.
				Place(&board, p, x, y, '.') // remove
			}
		}
	}
	// No valid placement found for this piece at this recursion state.
	return false
}

func CanPlace(board []string, p Piece, ox, oy int) bool {
	// Check every occupied cell of piece p against board bounds and collisions.
	n := len(board)
	for _, b := range p.Blocks {
		x := ox + b.X
		y := oy + b.Y
		if x < 0 || y < 0 || x >= n || y >= n {
			return false
		}
		if rune(board[y][x]) != '.' {
			return false
		}
	}
	return true
}

func Place(board *[]string, p Piece, ox, oy int, ch rune) {
	// Write piece cells (or '.') into the board at the given offset.
	for _, b := range p.Blocks {
		x := ox + b.X
		y := oy + b.Y

		r := []rune((*board)[y])
		r[x] = ch
		(*board)[y] = string(r)
	}
}

func ParsePiece(lot []rune) Piece {
	// lot is 4 lines of 4 + '\n' (20 runes), and contains 'A'.. or '.' and '\n'
	var id rune
	points := make([]Point, 0, 4)

	// Read occupied cells and capture their coordinates.
	x, y := 0, 0
	for _, r := range lot {
		if r == '\n' {
			y++
			x = 0
			continue
		}
		if r != '.' {
			if id == 0 {
				id = r
			}
			points = append(points, Point{X: x, Y: y})
		}
		x++
	}

	// normalize top-left to (0,0)
	minX, minY := points[0].X, points[0].Y
	for _, pt := range points {
		if pt.X < minX {
			minX = pt.X
		}
		if pt.Y < minY {
			minY = pt.Y
		}
	}
	for i := range points {
		points[i].X -= minX
		points[i].Y -= minY
	}

	// Compute bounding box and copy into fixed-size [4]Point storage.
	maxX, maxY := 0, 0
	var blocks [4]Point
	for i := 0; i < 4; i++ {
		blocks[i] = points[i]
		if points[i].X > maxX {
			maxX = points[i].X
		}
		if points[i].Y > maxY {
			maxY = points[i].Y
		}
	}

	return Piece{
		ID:     id,
		Blocks: blocks,
		W:      maxX + 1,
		H:      maxY + 1,
	}
}
