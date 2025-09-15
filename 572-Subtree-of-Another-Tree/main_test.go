package p572

import (
	"math"
	"testing"

	"leetcode/lib"
)

func Test_isSubtree(t *testing.T) {
	type args struct {
		root    []int
		subRoot []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				root:    []int{3, 4, 5, 1, 2},
				subRoot: []int{4, 1, 2},
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				root:    []int{3, 4, 5, 1, 2, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, 0},
				subRoot: []int{4, 1, 2},
			},
			want: false,
		},
		{
			name: "case 3",
			args: args{
				root:    []int{1},
				subRoot: []int{},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.args.root)
			subRoot := buildTree(tt.args.subRoot)
			if got := isSubtree(root, subRoot); got != tt.want {
				t.Errorf("isSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BuildTree constructs a binary tree from a level-order slice representation.
// Uses math.MaxInt as a sentinel value to represent nil nodes in the input slice.
// Non-nil nodes (including negative values) are created normally.
//
// Example:
//   nums := []int{3, 9, 20, math.MaxInt, math.MaxInt, 15, 7}
//   Creates tree:
//       3
//      / \
//     9  20
//       /  \
//      15   7
//
// Time complexity: O(n) where n is len(nums)
// Space complexity: O(n) for the queue and tree nodes
// TODO: move to lib
func buildTree(nums []int) *lib.TreeNode {
	if len(nums) == 0 || nums[0] == -1 {
		return nil
	}

	root := &lib.TreeNode{Val: nums[0]}
	queue := []*lib.TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(nums) {
		node := queue[0]
		queue = queue[1:]

		// Process left child
		if i < len(nums) {
			if nums[i] != -1 {
				node.Left = &lib.TreeNode{Val: nums[i]}
				queue = append(queue, node.Left)
			}
			i++
		}

		// Process right child
		if i < len(nums) {
			if nums[i] != math.MaxInt {
				node.Right = &lib.TreeNode{Val: nums[i]}
				queue = append(queue, node.Right)
			}
			i++
		}
	}

	return root
}
