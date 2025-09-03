package p543

import (
	"leetcode/lib"
)

// diameterOfBinaryTree calculates the diameter of a binary tree, which is defined
// as the length of the longest path between any two nodes in the tree. This path
// may or may not pass through the root.
//
// The algorithm uses a single DFS traversal that simultaneously:
// 1. Calculates the height of each subtree
// 2. Updates the maximum diameter found so far
//
// At each node, the diameter passing through that node equals the sum of the
// heights of its left and right subtrees. The overall diameter is the maximum
// of all such node-specific diameters.
//
// Time complexity: O(n) where n is the number of nodes
// Space complexity: O(h) where h is the height of the tree (recursion stack)
//
// Example:
//     1
//    / \
//   2   3
//  / \
// 4   5
// Diameter = 3 (path: 4 -> 2 -> 1 -> 3 or 5 -> 2 -> 1 -> 3)
func diameterOfBinaryTree(root *lib.TreeNode) int {
	maxDiameter := 0

	var dfs func(root *lib.TreeNode) int
	dfs = func(root *lib.TreeNode) int {
		if root == nil {
			return 0
		}

		hl := dfs(root.Left)
		hr := dfs(root.Right)

		currentDiameter := hl + hr
		if currentDiameter > maxDiameter {
			maxDiameter = currentDiameter
		}

		// Return height of current subtree
		if hl > hr {
			return hl + 1
		}
		return hr + 1
	}

	dfs(root)

	return maxDiameter
}
