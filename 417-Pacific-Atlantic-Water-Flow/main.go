package p417

// pacificAtlantic finds all cells where water can flow to both Pacific and Atlantic oceans.
// Water flows from higher or equal height cells to lower height cells in 4 directions.
// Pacific ocean borders the top and left edges, Atlantic ocean borders the bottom and right edges.
// Returns coordinates of cells that can reach both oceans.
// Example: A cell at height 5 can flow to a neighboring cell at height 3, but not to height 7.
//
// Approach: Reverse flow analysis using DFS from ocean borders inward.
// Instead of checking if each cell can reach oceans, start from ocean edges and find
// which cells can be reached by flowing upward (reverse direction).
// A cell can reach both oceans if it's reachable from both Pacific and Atlantic borders.
// Time Complexity: O(m × n) where m and n are grid dimensions (DFS visits each cell at most twice).
// Space Complexity: O(m × n) for Pacific and Atlantic reachability maps + recursion stack.
func pacificAtlantic(heights [][]int) [][]int {
	rows, cols := len(heights), len(heights[0])
	pac := make(map[cell]bool) // Cells reachable from Pacific ocean
	atl := make(map[cell]bool) // Cells reachable from Atlantic ocean

	// DFS helper function to explore cells reachable from ocean borders
	var dfs func(row int, col int, visit map[cell]bool, prevHeight int)
	dfs = func(row int, col int, visit map[cell]bool, prevHeight int) {
		c := cell{row, col}

		// Skip if already visited
		if visit[c] {
			return
		}

		// Boundary check or height constraint violation (water can't flow uphill)
		if row < 0 || row == rows || col < 0 || col == cols || heights[row][col] < prevHeight {
			return
		}

		// Mark cell as reachable from this ocean
		visit[c] = true

		// Explore all 4 directions with current height as new minimum
		dfs(row-1, col, visit, heights[row][col]) // Up
		dfs(row+1, col, visit, heights[row][col]) // Down
		dfs(row, col-1, visit, heights[row][col]) // Left
		dfs(row, col+1, visit, heights[row][col]) // Right
	}

	// Start DFS from Pacific borders (top and left edges)
	for c := range cols {
		dfs(0, c, pac, heights[0][c])           // Top edge → Pacific
		dfs(rows-1, c, atl, heights[rows-1][c]) // Bottom edge → Atlantic
	}

	for r := range rows {
		dfs(r, 0, pac, heights[r][0])           // Left edge → Pacific
		dfs(r, cols-1, atl, heights[r][cols-1]) // Right edge → Atlantic
	}

	// Find intersection: cells reachable from both oceans
	result := make([][]int, 0)
	for row := range rows {
		for col := range cols {
			c := cell{row, col}
			// Cell can reach both oceans
			if pac[c] && atl[c] {
				result = append(result, []int{row, col})
			}
		}
	}

	return result
}

// cell represents a coordinate position in the 2D grid.
type cell struct {
	row int // Row index
	col int // Column index
}

// cellQueue implements a FIFO queue for BFS traversal of grid cells.
type cellQueue struct {
	items []cell // Underlying slice storing queue elements
	head  int    // Index of first element (front of queue)
}

// enqueue adds a cell to the back of the queue.
func (q *cellQueue) enqueue(item cell) {
	q.items = append(q.items, item)
}

// dequeue removes and returns the cell from the front of the queue.
func (q *cellQueue) dequeue() cell {
	item := q.items[q.head]
	q.head++
	return item
}

// len returns the number of elements currently in the queue.
func (q *cellQueue) len() int {
	return len(q.items) - q.head
}
