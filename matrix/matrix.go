package matrix

func IntMatrix(rows, cols int) [][]int {
	cells := make([][]int, rows)
	for i := range cells {
		cells[i] = make([]int, cols)
	}

	return cells
}
