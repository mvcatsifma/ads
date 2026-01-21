package p207

// canFinish determines if all courses can be completed given prerequisites.
// Uses DFS cycle detection on the prerequisite dependency graph.
// Time: O(V + E), Space: O(V + E) where V = courses, E = prerequisites
func canFinish(numCourses int, prerequisites [][]int) bool {
	// Build adjacency list: prereq -> [dependent courses]
	adj := make([][]int, numCourses)
	for _, prerequisite := range prerequisites {
		course := prerequisite[0]
		prereq := prerequisite[1]
		adj[prereq] = append(adj[prereq], course)
	}

	// Track visited nodes and current recursion stack
	visited := make([]bool, numCourses)
	recStack := make([]bool, numCourses) // Detects back edges (cycles)

	// DFS helper: returns true if cycle detected
	var dfs func(v int) bool
	dfs = func(v int) bool {
		visited[v] = true
		recStack[v] = true // Mark as part of current path

		// Explore all dependent courses
		for _, neighbor := range adj[v] {
			if !visited[neighbor] {
				if dfs(neighbor) {
					return true // Cycle found in subtree
				}
			} else if recStack[neighbor] {
				return true // Back edge detected = cycle
			}
		}

		recStack[v] = false // Remove from current path (backtrack)
		return false
	}

	// Check all disconnected components for cycles
	for course := 0; course < numCourses; course++ {
		if !visited[course] {
			if dfs(course) {
				return false // Cycle found - impossible to complete
			}
		}
	}

	return true // No cycles - all courses can be completed
}
