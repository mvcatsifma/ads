package copyListWithRandomPointer

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_copyRandomList(t *testing.T) {
	type args struct {
		head [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case 1",
			args: args{
				head: [][]int{
					{7, -1}, {13, 0}, {11, 4}, {10, 2}, {1, 0},
				},
			},
			want: [][]int{
				{7, -1}, {13, 0}, {11, 4}, {10, 2}, {1, 0},
			},
		},
		{
			name: "case 2",
			args: args{
				head: [][]int{
					{1, 1}, {2, 1},
				},
			},
			want: [][]int{
				{1, 1}, {2, 1},
			},
		},
		{
			name: "case 2",
			args: args{
				head: [][]int{
					{3, -1}, {3, 0}, {3, -1},
				},
			},
			want: [][]int{
				{3, -1}, {3, 0}, {3, -1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createLinkedListWithRandom(tt.args.head)
			got := convertToArray(copyRandomList(head))
			if !cmp.Equal(got, tt.want) {
				t.Errorf("copyRandomList() = %v, want %v", got, tt.want)
			}
		})
	}
}

// createLinkedListWithRandom creates a linked list with random pointers
// from a nested array where each element is [val, random_index]
func createLinkedListWithRandom(arr [][]int) *Node {
	if len(arr) == 0 {
		return nil
	}

	// First pass: create nodes and store in slice for random access
	nodes := make([]*Node, len(arr))
	for i, pair := range arr {
		nodes[i] = &Node{Val: pair[0]}
	}

	// Second pass: connect Next and Random pointers
	for i, pair := range arr {
		// Connect Next pointer
		if i < len(arr)-1 {
			nodes[i].Next = nodes[i+1]
		}

		// Connect Random pointer if random_index is not null (-1)
		if len(pair) > 1 && pair[1] != -1 { // Using -1 to represent null
			nodes[i].Random = nodes[pair[1]]
		}
	}

	return nodes[0]
}

// Helper function to convert linked list back to array format for testing
func convertToArray(head *Node) [][]int {
	if head == nil {
		return nil
	}

	// First pass: create node to index mapping
	nodeToIndex := make(map[*Node]int)
	curr := head
	index := 0
	for curr != nil {
		nodeToIndex[curr] = index
		index++
		curr = curr.Next
	}

	// Second pass: create result array
	result := make([][]int, len(nodeToIndex))
	curr = head
	index = 0
	for curr != nil {
		randomIndex := -1 // represent null
		if curr.Random != nil {
			randomIndex = nodeToIndex[curr.Random]
		}
		result[index] = []int{curr.Val, randomIndex}
		index++
		curr = curr.Next
	}

	return result
}
