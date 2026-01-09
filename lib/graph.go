package lib

// GraphNode Node represents a vertex in an undirected graph.
type GraphNode struct {
	Val       int
	Neighbors []*GraphNode
}

// CreateGraphFromAdjList creates a graph from a 1-indexed adjacency list representation.
// The adjacency list uses 1-based indexing where list[i] contains the neighbors of node (i+1).
// Returns a pointer to node 1, or nil if the graph is empty.
//
// Parameters:
//   - list: adjacency list where list[i] represents neighbors of node (i+1)
//           Example: [[2,4],[1,3],[2,4],[1,3]] represents a 4-node cycle
//
// Returns:
//   - *lib.GraphNode: pointer to the node with value 1, or nil if input is empty
//
// Assumptions:
//   - All neighbor references are valid (within range 1 to len(list))
//   - Graph contains at least node 1 if non-empty
//
// Time Complexity: O(V + E) where V is number of vertices, E is number of edges
// Space Complexity: O(V) for the nodes map
func CreateGraphFromAdjList(list [][]int) *GraphNode {
	// Handle empty graph
	if len(list) == 0 {
		return nil
	}

	// Phase 1: Create all nodes (1-indexed: nodes 1, 2, 3, ..., len(list))
	nodes := make(map[int]*GraphNode, len(list))
	for i := range list {
		nodeVal := i + 1 // Convert 0-indexed array to 1-indexed node values
		nodes[nodeVal] = &GraphNode{Val: nodeVal}
	}

	// Phase 2: Build neighbor relationships from adjacency list
	for i, neighbors := range list {
		nodeVal := i + 1
		node := nodes[nodeVal] // Safe access - node guaranteed to exist from phase 1

		// Pre-allocate neighbors slice with known capacity for efficiency
		node.Neighbors = make([]*GraphNode, 0, len(neighbors))

		// Add each neighbor reference
		for _, neighborVal := range neighbors {
			if neighbor, exists := nodes[neighborVal]; exists {
				node.Neighbors = append(node.Neighbors, neighbor)
			}
			// Note: Silently skips invalid neighbor references
			// Could add error handling here for invalid references
		}
	}

	// Return pointer to node 1 (assumes node 1 exists in well-formed input)
	return nodes[1]
}
