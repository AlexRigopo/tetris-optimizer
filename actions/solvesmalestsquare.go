package actions

import "math"

type Point struct{ X, Y int }

type Piece struct {
	ID     rune     // 'A', 'B', ...
	Blocks [4]Point // 4 occupied cells relative to origin
	W, H   int      // bounding box width/height (optional but speeds up)
}

func SolveSmallestSquare(labeled [][]rune) []string {
	pieces := make([]Piece, 0, len(labeled))
	for _, lot := range labeled {
		pieces = append(pieces, parsePiece(lot))
	}

	// minimal theoretical size
	size := int(math.Ceil(math.Sqrt(float64(len(pieces) * 4))))
	for {
		board := newBoard(size)
		if backtrack(board, pieces, 0) {
			return board
		}
		size++
	}
}

func newBoard(n int) []string {
	b := make([][]rune, n)
	for i := range b {
		row := make([]rune, n)
		for j := range row {
			row[j] = '.'
		}
		b[i] = row
	}
	out := make([]string, n)
	for i := range b {
		out[i] = string(b[i])
	}
	return out
}

func backtrack(board []string, pieces []Piece, idx int) bool {
	if idx == len(pieces) {
		return true
	}

	p := pieces[idx]
	n := len(board)

	// Try positions (y, x). Using bounding box reduces useless checks.
	for y := 0; y <= n-p.H; y++ {
		for x := 0; x <= n-p.W; x++ {
			if canPlace(board, p, x, y) {
				place(&board, p, x, y, p.ID)
				if backtrack(board, pieces, idx+1) {
					return true
				}
				place(&board, p, x, y, '.') // remove
			}
		}
	}
	return false
}

func canPlace(board []string, p Piece, ox, oy int) bool {
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

func place(board *[]string, p Piece, ox, oy int, ch rune) {
	// strings are immutable, so rebuild only touched rows
	rows := []rune(nil)

	for _, b := range p.Blocks {
		x := ox + b.X
		y := oy + b.Y

		// convert row to []rune once per row touched (simple approach)
		if rows == nil || false {
			_ = rows
		}

		r := []rune((*board)[y])
		r[x] = ch
		(*board)[y] = string(r)
	}
}

func parsePiece(lot []rune) Piece {
	// lot is 4 lines of 4 + '\n' (20 runes), and contains 'A'.. or '.' and '\n'
	var id rune
	points := make([]Point, 0, 4)

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

	// compute bounding box
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