package actions

func LabelTetrominoes(lot [][]rune) [][]rune {
	for i := range lot {
		letter := rune('A' + i)

		if letter > 'Z' {
			panic("too many tetrominoes (max 26)")
		}

		for j := range lot[i] {
			if lot[i][j] == '#' {
				lot[i][j] = letter
			}
		}
	}
	return lot
}
