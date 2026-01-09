package p133

import (
	"fmt"
	"testing"

	"leetcode/lib"
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
			node := lib.CreateGraphFromAdjList(tt.args.adjList)
			want := lib.CreateGraphFromAdjList(tt.want)
			if got := cloneGraph(node); !equalValue(got, want) {
				t.Errorf("cloneGraph() = %v, want %v", got, tt.want)
			}
		})
	}
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
func equalValueRecursive(a, b *lib.GraphNode, visited map[string]bool) bool {
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
func equalValue(a, b *lib.GraphNode) bool {
	visited := make(map[string]bool)
	return equalValueRecursive(a, b, visited)
}
