package p133

import (
	"fmt"
	"testing"
)

func Test_cloneGraph(t *testing.T) {
	type args struct {
		adjList [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case 1",
			args: args{
				adjList: [][]int{
					{2, 4}, {1, 3}, {2, 4}, {1, 3},
				},
			},
			want: [][]int{
				{2, 4}, {1, 3}, {2, 4}, {1, 3},
			},
		},
		{
			name: "case 2",
			args: args{
				adjList: [][]int{{}},
			},
			want: [][]int{{}},
		},
		{
			name: "case 3",
			args: args{
				adjList: [][]int{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := createListFromAdjList(tt.args.adjList)
			want := createListFromAdjList(tt.want)
			if got := cloneGraph(node); !equalValue(got, want) {
				t.Errorf("cloneGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}

// createListFromAdjList creates a graph from a 1-indexed adjacency list representation.
// The adjacency list uses 1-based indexing where list[i] contains the neighbors of node (i+1).
// Returns a pointer to node 1, or nil if the graph is empty.
//
// Parameters:
//   - list: adjacency list where list[i] represents neighbors of node (i+1)
//           Example: [[2,4],[1,3],[2,4],[1,3]] represents a 4-node cycle
//
// Returns:
//   - *Node: pointer to the node with value 1, or nil if input is empty
//
// Assumptions:
//   - All neighbor references are valid (within range 1 to len(list))
//   - Graph contains at least node 1 if non-empty
//
// Time Complexity: O(V + E) where V is number of vertices, E is number of edges
// Space Complexity: O(V) for the nodes map
func createListFromAdjList(list [][]int) *Node {
	// Handle empty graph
	if len(list) == 0 {
		return nil
	}

	// Phase 1: Create all nodes (1-indexed: nodes 1, 2, 3, ..., len(list))
	nodes := make(map[int]*Node, len(list))
	for i := range list {
		nodeVal := i + 1 // Convert 0-indexed array to 1-indexed node values
		nodes[nodeVal] = &Node{Val: nodeVal}
	}

	// Phase 2: Build neighbor relationships from adjacency list
	for i, neighbors := range list {
		nodeVal := i + 1
		node := nodes[nodeVal] // Safe access - node guaranteed to exist from phase 1

		// Pre-allocate neighbors slice with known capacity for efficiency
		node.Neighbors = make([]*Node, 0, len(neighbors))

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

// equalValueRecursive compares two Node objects recursively and returns true if they are
// different objects but have the same value and equivalent neighbor structure.
// Returns false if they are the same object, have different values, or different neighbor patterns.
//
// Parameters:
//   - a, b: Node pointers to compare
//   - visited: map to track visited node pairs to avoid infinite recursion in cycles
//
// Returns:
//   - true: if a and b are different objects with equal values and equivalent neighbor structures
//   - false: if they are the same object, have different values, or different neighbor structures
func equalValueRecursive(a, b *Node, visited map[string]bool) bool {
	// Handle nil cases
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false // One nil, one not
	}

	// Same object reference - not different objects
	if a == b {
		return false
	}

	// Different values
	if a.Val != b.Val {
		return false
	}

	// Create unique key for this comparison pair to detect cycles
	key := fmt.Sprintf("%p-%p", a, b)
	if visited[key] {
		return true // Already compared this pair, assume equal to avoid infinite recursion
	}
	visited[key] = true

	// Different number of neighbors
	if len(a.Neighbors) != len(b.Neighbors) {
		return false
	}

	// Recursively compare all neighbors
	// Note: This assumes neighbors are in the same order
	for i := 0; i < len(a.Neighbors); i++ {
		if !equalValueRecursive(a.Neighbors[i], b.Neighbors[i], visited) {
			return false
		}
	}

	return true
}

// Wrapper function for easier usage
func equalValue(a, b *Node) bool {
	visited := make(map[string]bool)
	return equalValueRecursive(a, b, visited)
}
