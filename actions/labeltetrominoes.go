package actions

import "fmt"

func LabelTetrominoes(lot [][]rune) ([][]rune, error) {
	// Assign one letter per tetromino in input order.
	for i := range lot {
		letter := rune('A' + i)

		// We support A..Z only, so reject piece counts above 26.
		if letter > 'Z' {
			return nil, fmt.Errorf("ERROR")
		}

		// Replace each occupied cell marker '#' with the piece letter.
		for j := range lot[i] {
			if lot[i][j] == '#' {
				lot[i][j] = letter
			}
		}
	}
	// Return the now-labeled piece list.
	return lot, nil
}
