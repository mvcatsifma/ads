package main

type key struct {
	x int // row / 3
	y int // column /3
}

func isValidSudoku(board [][]byte) bool {
	var columns = make([][]byte, 9)
	boxes := make(map[key][]byte)

	// check if each row is valid
	for ri := 0; ri < 9; ri++ {
		row := board[ri]
		if !isValid(row) {
			return false
		}
		for ci := 0; ci < 9; ci++ {
			v := row[ci]
			columns[ci] = append(columns[ci], v)
			k := key{x: ci / 3, y: ri / 3} // keys are 0...8, unique for each box
			boxes[k] = append(boxes[k], v)
		}
	}

	// check if each column is valid
	for _, column := range columns {
		if !isValid(column) {
			return false
		}
	}

	// check if each box is valid
	for _, box := range boxes {
		if !isValid(box) {
			return false
		}
	}

	return true
}

func isValid(cells []byte) bool {
	m := make(map[byte]bool)
	for _, value := range cells {
		if _, ok := m[value]; ok && value != '.' { // . is allowed multiple times
			return false
		}
		m[value] = true
	}
	return true
}
