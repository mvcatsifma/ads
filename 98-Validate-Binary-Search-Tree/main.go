package p98

import (
	"math"

	"leetcode/lib"
)

// isValidBST validates whether a binary tree is a valid binary search tree.
// A valid BST requires: all left subtree values < node value AND all right subtree values > node value.
// This constraint applies recursively to all descendants, not just immediate children.
// Example: [5,1,4,null,null,3,6] returns false because 3 is in right subtree of 5 but 3 < 5.
//
// Approach: Uses DFS with min/max bounds. For each node, validates that its value is strictly
// within the valid range [left, right] inherited from ancestors. Updates bounds as we traverse:
// left subtree gets upper bound = node.Val, right subtree gets lower bound = node.Val.
// Time Complexity: O(n) where n is the number of nodes (visit each node once).
// Space Complexity: O(h) for recursion stack where h is the height of the tree.
func isValidBST(root *lib.TreeNode) bool {
	var dfs func(*lib.TreeNode, int, int) bool
	dfs = func(node *lib.TreeNode, left int, right int) bool {
		if node == nil {
			return true
		}

		if !(node.Val < right && node.Val > left) {
			return false
		}

		return dfs(node.Left, left, node.Val) && dfs(node.Right, node.Val, right)
	}

	return dfs(root, math.MinInt, math.MaxInt)
}
