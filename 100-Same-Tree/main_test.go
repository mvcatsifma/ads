package p100

import (
	"math"
	"testing"

	"leetcode/lib"
)

func Test_isSameTree(t *testing.T) {
	type args struct {
		p []int
		q []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				p: []int{1, 2, 3},
				q: []int{1, 2, 3},
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				p: []int{1, 2},
				q: []int{1, math.MaxInt, 2},
			},
			want: false,
		},
		{
			name: "case 3",
			args: args{
				p: []int{1, 2, 1},
				q: []int{1, 1, 2},
			},
			want: false,
		},
		{
			name: "case 4",
			args: args{
				p: []int{1, -1},
				q: []int{1, math.MaxInt, -1},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := buildTree(tt.args.p)
			q := buildTree(tt.args.q)
			if got := isSameTree(p, q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
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
