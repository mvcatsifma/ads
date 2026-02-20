package directed

import (
	"errors"
)

// DepthFirstOrder performs DFS traversal and returns vertices in three different orderings.
// Returns pre-order (when first visited), post-order (when finished), and reverse post-order.
// Reverse post-order is equivalent to topological ordering for DAGs (Directed Acyclic Graphs).
// Time: O(V + E), Space: O(V) where V = vertices, E = edges
func DepthFirstOrder(g *Digraph) (pre *IntQueue, post *IntQueue, reverse *IntStack) {
	// Initialize data structures for the three orderings
	pre = &IntQueue{}     // Pre-order: vertices in order of first discovery
	post = &IntQueue{}    // Post-order: vertices in order of completion
	reverse = &IntStack{} // Reverse post-order: topological order for DAGs

	marked := make([]bool, g.V) // Track visited vertices

	// DFS traversal with ordering tracking
	var dfs func(v int)
	dfs = func(v int) {
		pre.Enqueue(v)   // Record vertex when first discovered (pre-order)
		marked[v] = true // Mark as visited to prevent revisiting

		// Recursively visit all unvisited adjacent vertices
		for _, w := range g.Adjacent(v) {
			if !marked[w] {
				dfs(w) // Explore unvisited neighbor
			}
		}

		// Record vertex when completely processed (post-order)
		post.Enqueue(v) // Post-order: added after all descendants processed
		reverse.Push(v) // Reverse post-order: stack reverses the post-order
	}

	// Process all vertices to handle disconnected components
	for v := range g.V {
		if !marked[v] {
			dfs(v) // Start DFS from unvisited vertex
		}
	}

	return // Named return values
}

// IntStack represents a LIFO (Last In, First Out) data structure for integers
type IntStack struct {
	items []int
}

// New creates a new stack with pre-allocated capacity
func New(capacity int) *IntStack {
	return &IntStack{
		items: make([]int, 0, capacity),
	}
}

// Push adds an item to the top of the stack
func (s *IntStack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack
func (s *IntStack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]

	return item, nil
}

// Peek returns the top item without removing it
func (s *IntStack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}

	return s.items[len(s.items)-1], nil
}

// IsEmpty returns true if the stack is empty
func (s *IntStack) IsEmpty() bool {
	return len(s.items) == 0
}
