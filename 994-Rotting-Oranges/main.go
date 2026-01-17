package p994

// orangesRotting calculates the minimum time for all fresh oranges to rot.
// Rotten oranges spread to adjacent fresh oranges every minute (4-directional spread).
// Returns the time in minutes, or -1 if some oranges can never rot.
// Example: grid with fresh oranges at (0,0) and rotten at (0,1) returns 1 minute.
//
// Approach: Multi-source BFS starting from all initially rotten oranges.
// Each BFS level represents one minute of rot spreading.
// Time Complexity: O(rows × cols) - visit each cell at most once.
// Space Complexity: O(rows × cols) for the queue in worst case.
func orangesRotting(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	queue := &cellQueue{}
	fresh := 0 // Count of fresh oranges remaining

	// Phase 1: Find all initially rotten oranges and count fresh ones
	for r := range rows {
		for c := range cols {
			if grid[r][c] == 2 { // Found a rotten orange
				queue.enqueue(cell{r, c})
			} else if grid[r][c] == 1 { // Found a fresh orange
				fresh++
			}
		}
	}

	// Helper function to safely rot adjacent fresh oranges
	enqueueCell := func(item cell) {
		// Boundary check
		if item.row < 0 || item.row >= rows || item.col < 0 || item.col >= cols {
			return
		}
		// Only rot fresh oranges (value 1)
		if grid[item.row][item.col] != 1 {
			return
		}

		// Convert fresh orange to rotten
		grid[item.row][item.col] = 2
		queue.enqueue(item)
		fresh-- // Decrease fresh orange count
	}

	// Phase 2: BFS simulation - spread rot minute by minute
	time := 0
	for queue.len() > 0 && fresh > 0 {
		levelSize := queue.len() // Number of oranges that will rot this minute

		// Process all oranges that rot in the current minute
		for i := 0; i < levelSize; i++ {
			item := queue.dequeue()

			// Spread rot to all 4 adjacent cells
			enqueueCell(cell{row: item.row - 1, col: item.col}) // Up
			enqueueCell(cell{row: item.row + 1, col: item.col}) // Down
			enqueueCell(cell{row: item.row, col: item.col - 1}) // Left
			enqueueCell(cell{row: item.row, col: item.col + 1}) // Right
		}
		time++ // One minute has passed after processing all oranges at this level
	}

	// Return result based on remaining fresh oranges
	if fresh == 0 {
		return time // All oranges rotted
	}
	return -1 // Some oranges can never rot (isolated)
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
