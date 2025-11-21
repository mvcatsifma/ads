package p235

import (
	"math"
	"testing"

	"leetcode/lib"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root []int
		p    int
		q    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				root: []int{6, 2, 8, 0, 4, 7, 9, math.MaxInt, math.MaxInt, 3, 5},
				p:    2,
				q:    8,
			},
			want: 6,
		},
		{
			name: "case 2",
			args: args{
				root: []int{6, 2, 8, 0, 4, 7, 9, math.MaxInt, math.MaxInt, 3, 5},
				p:    2,
				q:    4,
			},
			want: 2,
		},
		{
			name: "case 3",
			args: args{
				root: []int{2, 1},
				p:    2,
				q:    1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.args.root)
			p := &lib.TreeNode{Val: tt.args.p}
			q := &lib.TreeNode{Val: tt.args.q}
			got := lowestCommonAncestor(root, p, q)
			if got.Val != tt.want {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got.Val, tt.want)
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
