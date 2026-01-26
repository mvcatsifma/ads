package p210

// findOrder returns a valid course ordering using topological sort with cycle detection.
// Returns empty slice if prerequisites contain a cycle (impossible to complete all courses).
// Time: O(V + E), Space: O(V) where V = numCourses, E = len(prerequisites)
func findOrder(numCourses int, prerequisites [][]int) []int {
	// Build adjacency list: prereq -> [dependent courses]
	adj := make([][]int, numCourses)
	for _, prerequisite := range prerequisites {
		course := prerequisite[0] // Course that depends on prereq
		prereq := prerequisite[1] // Prerequisite course
		adj[prereq] = append(adj[prereq], course)
	}

	stack := NewStack()

	// Track visited nodes and current recursion stack for cycle detection
	visited := make([]bool, numCourses)  // Tracks all visited nodes (white -> gray/black)
	recStack := make([]bool, numCourses) // Tracks current DFS path (gray nodes)

	// DFS helper: returns true if cycle detected in current path
	var dfs func(v int) bool
	dfs = func(v int) bool {
		visited[v] = true  // Mark as visited (gray)
		recStack[v] = true // Add to current recursion path

		// Explore all courses that depend on current course
		for _, neighbor := range adj[v] {
			if !visited[neighbor] {
				// Unvisited neighbor - recurse
				if dfs(neighbor) {
					return true // Cycle found in subtree
				}
			} else if recStack[neighbor] {
				// Visited neighbor in current path = back edge = cycle
				return true
			}
			// else: visited neighbor not in current path (cross/forward edge) - safe
		}

		stack.Push(v)       // Add to result in reverse post-order
		recStack[v] = false // Remove from current path (backtrack to white)
		return false        // No cycle found
	}

	// Process all disconnected components
	for course := 0; course < numCourses; course++ {
		if !visited[course] {
			if dfs(course) {
				return []int{} // Cycle detected - no valid ordering exists
			}
		}
	}

	// Pop stack to get topological ordering
	var retVal []int
	for !stack.IsEmpty() {
		item, _ := stack.Pop() // Safe to ignore bool since we know stack isn't empty
		retVal = append(retVal, item)
	}

	return retVal
}

// Stack implements a simple integer stack using a slice
type Stack struct {
	items []int
}

// NewStack creates a new empty stack with pre-allocated capacity
func NewStack() *Stack {
	return &Stack{
		items: make([]int, 0, 8), // Start with capacity 8 to reduce allocations
	}
}

// Push adds an item to the top of the stack
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack
// Returns (0, false) if stack is empty
func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false // Empty stack
	}

	idx := len(s.items) - 1
	item := s.items[idx]
	s.items = s.items[:idx] // Shrink slice
	return item, true
}

// IsEmpty returns true if the stack contains no items
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}
