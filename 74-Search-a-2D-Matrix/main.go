package p74

// searchMatrix searches for target in a sorted m×n matrix where each row is
// sorted and the first integer of each row is greater than the last integer
// of the previous row. Uses O(log(m*n)) time and O(1) space.
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	rows, cols := len(matrix), len(matrix[0])
	l, r := 0, rows*cols-1

	for l <= r {
		mid := l + (r-l)/2
		val := matrix[mid/cols][mid%cols]
		if val == target {
			return true
		} else if val > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return false
}
