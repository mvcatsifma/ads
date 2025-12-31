package p200

// numIslands counts the number of islands in a 2D binary grid using DFS with in-place modification.
// An island is a group of connected '1's (land) surrounded by '0's (water).
// Connection is horizontal or vertical only (not diagonal).
// Example: grid = [["1","1","0"],["1","0","0"],["0","0","1"]] returns 2
//
// Approach: Uses DFS to explore each connected component of '1's.
// Modifies the grid in-place by changing visited '1's to '0's to avoid revisiting.
// For each unvisited '1', performs DFS to mark all connected land cells and counts as one island.
// Time Complexity: O(m × n) where m and n are grid dimensions (visit each cell once).
// Space Complexity: O(min(m, n)) for recursion stack in worst case (long thin island).
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	count := 0

	var dfs func(i, j int)
	dfs = func(i, j int) {
		// Boundary check or already visited/water
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != '1' {
			return
		}

		// Mark as visited by changing '1' to '0'
		grid[i][j] = '0'

		// Explore 4 directions
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfs(i, j)
				count++
			}
		}
	}

	return count
}
