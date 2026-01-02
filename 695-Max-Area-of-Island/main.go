package p695

// maxAreaOfIsland finds the maximum area of an island in a 2D binary grid using DFS.
// An island is a group of connected 1's (land) surrounded by 0's (water).
// Connection is horizontal or vertical only (not diagonal).
// Returns the area (number of cells) of the largest island, or 0 if no islands exist.
// Example: grid = [[1,1,0],[1,0,0],[0,0,1]] returns 3 (largest island has 3 cells)
//
// Approach: Uses DFS to explore each connected component and calculate its area.
// Modifies the grid in-place by changing visited 1's to 0's to avoid revisiting.
// For each unvisited land cell, performs DFS to count all connected cells and tracks maximum.
// Time Complexity: O(m × n) where m and n are grid dimensions (visit each cell once).
// Space Complexity: O(min(m, n)) for recursion stack in worst case (long thin island).
func maxAreaOfIsland(grid [][]int) int {
	// Edge case: empty grid
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	maxArea := 0

	// DFS helper function to explore and count connected land cells
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		// Base cases: out of bounds, water, or already visited
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != 1 {
			return 0
		}

		// Mark current land cell as visited by converting to water
		grid[i][j] = 0

		// Count current cell (1) plus area from all 4 adjacent directions
		return 1 + dfs(i-1, j) + dfs(i+1, j) + dfs(i, j-1) + dfs(i, j+1)
	}

	// Iterate through each cell in the grid
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// Found an unvisited land cell: calculate island area
			if grid[i][j] == 1 {
				area := dfs(i, j) // Get total area of this connected island
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}
