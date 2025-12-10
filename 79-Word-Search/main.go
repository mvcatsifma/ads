package p79

// exist determines if a given word can be found in a 2D character board.
// The word must be constructed from letters of sequentially adjacent cells,
// where adjacent cells are horizontally or vertically neighboring (not diagonal).
// Each cell can only be used once per word search.
// Example: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED" returns true
//
// Approach: Uses backtracking to explore all possible paths from each starting position.
// For each cell, tries to build the word by exploring all 4 directions recursively.
// Uses a visited array to prevent reusing cells within the same path.
// Time Complexity: O(m × n × 4^L) where m,n are board dimensions and L is word length.
// Space Complexity: O(m × n) for visited array + O(L) for recursion stack.
func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 || len(word) == 0 {
		return false
	}

	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	var backtrack func(row, col, idx int) bool
	backtrack = func(row, col, idx int) bool {
		// Found complete word
		if idx == len(word) {
			return true
		}

		// Boundary checks and character match
		if row < 0 || col < 0 || row >= m || col >= n ||
			visited[row][col] || board[row][col] != word[idx] {
			return false
		}

		// Mark as visited
		visited[row][col] = true

		// Explore all 4 directions
		found := backtrack(row+1, col, idx+1) ||
			backtrack(row-1, col, idx+1) ||
			backtrack(row, col+1, idx+1) ||
			backtrack(row, col-1, idx+1)

		// Backtrack
		visited[row][col] = false

		return found
	}

	// Try starting from each cell
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if backtrack(i, j, 0) {
				return true // Early termination
			}
		}
	}

	return false
}
