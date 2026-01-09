package p286

// wallsAndGates fills empty rooms with their distance to the nearest gate using multi-source BFS.
// Uses 0 for gates, -1 for walls, and INF (2147483647) for empty rooms to be filled.
// After execution, each empty room contains the minimum distance to any gate.
// Example: A room with value 3 means the nearest gate is 3 steps away.
//
// Approach: Multi-source BFS starting from all gates simultaneously.
// All gates are added to the queue initially, then BFS explores outward level by level.
// Each level represents rooms at distance d, d+1, d+2, etc. from the nearest gate.
// Time Complexity: O(m × n) where m and n are grid dimensions (visit each cell once).
// Space Complexity: O(m × n) for visited array and queue storage.
func wallsAndGates(rooms [][]int) {
	// Handle empty grid
	if len(rooms) == 0 || len(rooms[0]) == 0 {
		return
	}

	rows, cols := len(rooms), len(rooms[0])

	// Track visited cells to avoid revisiting during BFS
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	q := newQueue()

	// Helper function to safely add valid neighboring cells to queue
	var addRoom func(r, c int)
	addRoom = func(r, c int) {
		// Skip if out of bounds, already visited, or is a wall
		if r < 0 || r == rows || c < 0 || c == cols || visited[r][c] || rooms[r][c] == -1 {
			return
		}
		visited[r][c] = true
		q.Enqueue(cell{r, c})
	}

	// Phase 1: Add all gates (value 0) to queue as starting points for multi-source BFS
	for r := range rows {
		for c := range cols {
			if rooms[r][c] == 0 {
				q.Enqueue(cell{r, c})
				visited[r][c] = true
			}
		}
	}

	// Phase 2: Level-order BFS to fill distances
	dist := 0
	for q.Len() > 0 {
		// Process all cells at current distance level
		for levelSize := q.Len(); levelSize > 0; levelSize-- {
			item := q.Dequeue()
			r, c := item.r, item.c

			// Set distance for current cell
			rooms[r][c] = dist

			// Add all valid neighboring cells for next distance level
			addRoom(r+1, c) // Down
			addRoom(r-1, c) // Up
			addRoom(r, c+1) // Right
			addRoom(r, c-1) // Left
		}
		dist += 1 // Increment distance for next level
	}
}

// cell represents a coordinate position in the 2D grid.
type cell struct {
	r, c int // Row and column indices
}

// queue implements an efficient FIFO queue for BFS traversal using head pointer optimization.
// Provides O(1) enqueue and dequeue operations.
type queue struct {
	items []cell // Underlying slice storing queue elements
	head  int    // Index of first element (front of queue)
}

// newQueue creates and returns a new empty queue ready for use.
func newQueue() *queue {
	return &queue{
		items: make([]cell, 0),
		head:  0,
	}
}

// Enqueue adds a cell to the back of the queue.
// Time Complexity: O(1) amortized.
func (q *queue) Enqueue(cell cell) {
	q.items = append(q.items, cell)
}

// Dequeue removes and returns the cell from the front of the queue.
// Returns zero-value cell if queue is empty.
// Time Complexity: O(1).
func (q *queue) Dequeue() cell {
	if q.Len() == 0 || q.head >= len(q.items) {
		return cell{} // Return zero-value for empty queue
	}
	item := q.items[q.head]
	q.head++
	return item
}

// Len returns the number of elements currently in the queue.
// Time Complexity: O(1).
func (q *queue) Len() int {
	return len(q.items) - q.head
}
