package p119

// getRow returns the rowIndex-th row of Pascal's triangle using array-based index access.
// Each element is the sum of the two elements above it in the previous row.
// The row is 0-indexed: row 0 = [1], row 1 = [1,1], row 2 = [1,2,1], etc.
// Each row starts and ends with 1, and interior elements are computed from the previous row.
// Uses a 2D array to store all rows for efficient index-based access to previous row values.
//
// Time Complexity: O(rowIndex^2) to compute all rows up to rowIndex.
// Space Complexity: O(rowIndex^2) for storing all rows in the 2D array.
func getRow(rowIndex int) []int {
	numRows := rowIndex + 1
	rows := make([][]int, numRows)
	rows[0] = []int{1}

	if rowIndex == 0 {
		return rows[0]
	}

	for n := 1; n < numRows; n++ {
		rows[n] = make([]int, n+1)
		rows[n][0] = 1
		rows[n][n] = 1

		for k := 1; k < n; k++ {
			rows[n][k] = rows[n-1][k-1] + rows[n-1][k]
		}
	}

	return rows[rowIndex]
}
