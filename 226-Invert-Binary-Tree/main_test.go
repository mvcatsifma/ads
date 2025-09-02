package p226

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_invertTree(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				root: []int{4, 2, 7, 1, 3, 6, 9},
			},
			want: []int{4, 7, 2, 9, 6, 3, 1},
		},
		{
			name: "case 2",
			args: args{
				root: []int{},
			},
			want: []int{},
		},
		{
			name: "case 3",
			args: args{
				root: []int{2, 1, 3},
			},
			want: []int{2, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.args.root)
			want := buildTree(tt.want)
			if got := invertTree(root); !cmp.Equal(got, want) {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

// buildTree constructs a binary tree from a level-order (breadth-first) slice representation.
// The input slice represents nodes level by level, left to right.
// Returns nil if the input slice is empty.
//
// Example:
//   nums := []int{3, 9, 20, 15, 7}
//   Creates tree:
//       3
//      / \
//     9  20
//       /  \
//      15   7
//
// Time complexity: O(n) where n is len(nums)
// Space complexity: O(n) for the queue and tree nodes
func buildTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	root := &TreeNode{Val: nums[0]}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if i < len(nums) {
			node.Left = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Left)
			i++
		}

		if i < len(nums) {
			node.Right = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Right)
			i++
		}
	}

	return root
}
