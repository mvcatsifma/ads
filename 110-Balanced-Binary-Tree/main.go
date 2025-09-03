package p110

import (
	"math"

	"leetcode/lib"
)

// isBalanced determines if a binary tree is height-balanced. A height-balanced
// binary tree is defined as a tree in which the left and right subtrees of every
// node differ in height by no more than 1.
//
// The algorithm uses a single DFS traversal that simultaneously:
// 1. Calculates the height of each subtree
// 2. Checks the balance condition at each node
// 3. Sets a flag if any imbalance is found
//
// At each node, we compute the absolute difference between left and right subtree
// heights. If this difference exceeds 1 at any node, the entire tree is unbalanced.
//
// Time complexity: O(n) where n is the number of nodes
// Space complexity: O(h) where h is the height of the tree (recursion stack)
//
// Example:
//     3
//    / \
//   9  20
//     /  \
//    15   7
// This tree is balanced (all height differences ≤ 1)
//
// Counter-example:
//     1
//    /
//   2
//  /
// 3
// This tree is unbalanced (height difference = 2 at root)
func isBalanced(root *lib.TreeNode) bool {
	isBalanced := true

	var dfs func(root *lib.TreeNode) int
	dfs = func(root *lib.TreeNode) int {
		if root == nil {
			return 0
		}

		hl := dfs(root.Left)
		hr := dfs(root.Right)

		diff := math.Abs(float64(hl - hr))
		if diff > 1 {
			isBalanced = false
		}

		// Return height of current subtree
		if hl > hr {
			return hl + 1
		}
		return hr + 1
	}

	dfs(root)

	return isBalanced
}
