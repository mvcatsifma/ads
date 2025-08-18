package p118

// GeneratePascalsTriangle generates Pascal's triangle with the specified number of rows.
// Pascal's triangle is a triangular array where each number is the sum of the two
// numbers directly above it. The triangle starts with 1 at the top, and each row
// begins and ends with 1.
//
// Example for numRows = 5:
//     1
//    1 1
//   1 2 1
//  1 3 3 1
// 1 4 6 4 1
//
// Time complexity: O(numRows²)
// Space complexity: O(numRows²)
//
// Returns an empty slice if numRows is 0 or negative.
func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}

	result := make([][]int, numRows)
	result[0] = []int{1}

	for n := 1; n < numRows; n++ {
		result[n] = make([]int, n+1)
		result[n][0] = 1
		result[n][n] = 1

		for k := 1; k < n; k++ {
			result[n][k] = result[n-1][k-1] + result[n-1][k]
		}
	}

	return result
}
