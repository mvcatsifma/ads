package p36

type key struct {
	x int // row / 3
	y int // column /3
}

func isValidSudoku(board [][]byte) bool {
	const N = 9
	rows := make([]map[byte]bool, N)
	columns := make([]map[byte]bool, N)
	boxes := make(map[key]map[byte]bool, N)
	for i := 0; i < N; i++ {
		rows[i] = make(map[byte]bool)
		columns[i] = make(map[byte]bool)
	}

	for ri := 0; ri < N; ri++ {
		for ci := 0; ci < N; ci++ {
			val := board[ri][ci]
			if val == '.' {
				continue
			}
			if rows[ri][val] {
				return false
			}
			rows[ri][val] = true

			if columns[ci][val] {
				return false
			}
			columns[ci][val] = true

			k := key{ri / 3, ci / 3}
			if boxes[k] == nil {
				boxes[k] = make(map[byte]bool)
			}
			if boxes[k][val] {
				return false
			}
			boxes[k][val] = true
		}
	}

	return true
}
