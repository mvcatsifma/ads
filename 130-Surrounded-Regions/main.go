// Package p130 implements algorithms for the "Surrounded Regions" problem.
package p130

// solve captures all 'O' regions that are completely surrounded by 'X' on a 2D board.
// A region is surrounded if it cannot reach any border cell through connected 'O' cells.
// Connected means horizontally or vertically adjacent (not diagonal).
// Modifies the board in-place by changing surrounded 'O' cells to 'X'.
// Example: An 'O' region in the middle of the board gets captured, but 'O' regions
//          touching the border remain unchanged.
//
// Approach: Find all connected 'O' components using DFS, then check if each component
// touches the board boundary. Components that don't touch borders are surrounded and get captured.
// Time Complexity: O(m × n) where m and n are board dimensions (visit each cell once).
// Space Complexity: O(m × n) for visited map, components storage, and recursion stack.
func solve(board [][]byte) {
	rows, cols := len(board), len(board[0])
	visited := make(map[cell]bool)  // Track visited cells during DFS
	components := make([][]cell, 0) // Store all connected 'O' components
	count := 0                      // Current component index

	// DFS helper function to explore connected 'O' cells and build components
	var dfs func(c cell)
	dfs = func(c cell) {
		// Boundary check, non-'O' cell, or already visited
		if c.row < 0 || c.row == rows || c.col < 0 || c.col == cols ||
			board[c.row][c.col] != 'O' || visited[c] {
			return
		}

		visited[c] = true

		// Ensure components slice is large enough for current component index
		for len(components) <= count {
			components = append(components, []cell{})
		}

		// Add current cell to current component
		components[count] = append(components[count], c)

		// Explore all 4 adjacent directions
		dfs(cell{c.row - 1, c.col}) // Up
		dfs(cell{c.row + 1, c.col}) // Down
		dfs(cell{c.row, c.col - 1}) // Left
		dfs(cell{c.row, c.col + 1}) // Right
	}

	// Phase 1: Find all connected 'O' components using DFS
	for r := range rows {
		for c := range cols {
			if board[r][c] == 'O' && !visited[cell{r, c}] {
				dfs(cell{r, c}) // Explore entire component
				count++         // Move to next component index
			}
		}
	}

	// Phase 2: Identify surrounded components (don't touch any border)
	surroundedComponents := make([][]cell, 0)
	for _, component := range components {
		surrounded := true

		// Check if any cell in component touches the border
		for _, c := range component {
			if c.row == 0 || c.row == rows-1 || c.col == 0 || c.col == cols-1 {
				surrounded = false
				break // Component touches border, not surrounded
			}
		}

		// Component is completely surrounded
		if surrounded {
			surroundedComponents = append(surroundedComponents, component)
		}
	}

	// Phase 3: Capture all surrounded components by changing 'O' to 'X'
	for _, component := range surroundedComponents {
		for _, c := range component {
			board[c.row][c.col] = 'X'
		}
	}
}

// cell represents a coordinate position in the 2D board.
type cell struct {
	row int // Row index
	col int // Column index
}

// solveEfficient captures all 'O' regions that are completely surrounded by 'X' on a 2D board.
// A region is surrounded if it cannot reach any border cell through connected 'O' cells.
// Connected means horizontally or vertically adjacent (not diagonal).
// Modifies the board in-place by changing surrounded 'O' cells to 'X'.
//
// Approach: Reverse logic - find all 'O' cells that are NOT surrounded (can escape to border).
// Uses DFS from all border cells to mark "safe" 'O' cells, then captures remaining 'O' cells.
// This is more efficient than finding all components and checking each for border contact.
// Time Complexity: O(m × n) where m and n are board dimensions (visit each cell at most once).
// Space Complexity: O(m × n) for recursion stack in worst case (entire board is one region).
func solveEfficient(board [][]byte) {
	rows, cols := len(board), len(board[0])

	// DFS helper function to mark all border-connected 'O' cells as safe
	var dfs func(r, c int)
	dfs = func(r, c int) {
		// Base cases: out of bounds, not an 'O' cell, or already processed
		if r < 0 || r >= rows || c < 0 || c >= cols || board[r][c] != 'O' {
			return
		}

		// Mark as "safe" using temporary marker (won't be captured)
		board[r][c] = 'S'

		// Recursively mark all connected 'O' cells as safe
		dfs(r-1, c) // Up
		dfs(r+1, c) // Down
		dfs(r, c-1) // Left
		dfs(r, c+1) // Right
	}

	// Phase 1: Mark all border-connected 'O' cells as safe
	// Start DFS from all border cells to find escape routes
	for r := 0; r < rows; r++ {
		dfs(r, 0)      // Left border
		dfs(r, cols-1) // Right border
	}
	for c := 0; c < cols; c++ {
		dfs(0, c)      // Top border
		dfs(rows-1, c) // Bottom border
	}

	// Phase 2: Capture surrounded regions and restore safe regions
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if board[r][c] == 'O' {
				board[r][c] = 'X' // Surrounded 'O' → captured
			} else if board[r][c] == 'S' {
				board[r][c] = 'O' // Safe 'S' → restore to 'O'
			}
		}
	}
}
