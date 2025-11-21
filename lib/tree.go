package lib

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
func BuildTree(nums []int) *TreeNode {
	if len(nums) == 0 || nums[0] == -1 {
		return nil
	}

	root := &TreeNode{Val: nums[0]}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(nums) {
		node := queue[0]
		queue = queue[1:]

		// Process left child
		if i < len(nums) {
			if nums[i] != math.MaxInt {
				node.Left = &TreeNode{Val: nums[i]}
				queue = append(queue, node.Left)
			}
			i++
		}

		// Process right child
		if i < len(nums) {
			if nums[i] != math.MaxInt {
				node.Right = &TreeNode{Val: nums[i]}
				queue = append(queue, node.Right)
			}
			i++
		}
	}

	return root
}

// IsSameTree determines if two binary trees are structurally identical and have
// the same node values. Two trees are considered the same if they have identical
// structure and all corresponding nodes contain the same values.
//
// The algorithm uses recursive depth-first traversal to compare corresponding nodes
// simultaneously. At each step, it checks:
// 1. Both nodes are nil (base case - subtrees match)
// 2. Exactly one node is nil (structural mismatch)
// 3. Node values differ (value mismatch)
// 4. Recursively compare left and right subtrees
//
// The XOR check efficiently detects structural mismatches - if exactly one of the
// nodes is nil, the trees have different structures at this position.
//
// Time complexity: O(min(m,n)) where m,n are the number of nodes in each tree
// Space complexity: O(min(h1,h2)) where h1,h2 are the heights (recursion stack)
//
// Example:
//   Tree 1:  1      Tree 2:  1
//           / \              / \
//          2   3            2   3
// Returns true (identical structure and values)
//
//   Tree 1:  1      Tree 2:  1
//           /                \
//          2                  2
// Returns false (different structure)
func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q // true only if both nil
	}

	return p.Val == q.Val &&
		IsSameTree(p.Left, q.Left) &&
		IsSameTree(p.Right, q.Right)
}
