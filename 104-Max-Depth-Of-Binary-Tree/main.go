package p104

import (
	"math"

	"leetcode/lib"
)

// maxDepth calculates the maximum depth (height) of a binary tree.
// The depth is defined as the number of nodes along the longest path
// from the root node down to the farthest leaf node.
// Returns 0 for an empty tree (nil root).
//
// Example:
//       3
//      / \
//     9  20
//       /  \
//      15   7
//   maxDepth = 3 (path: 3 -> 20 -> 15 or 3 -> 20 -> 7)
//
// Time complexity: O(n) where n is the number of nodes
// Space complexity: O(h) where h is the height of the tree (recursion stack)
func maxDepth(root *lib.TreeNode) int {
	if root == nil {
		return 0
	}

	dl := maxDepth(root.Left)
	dr := maxDepth(root.Right)

	return 1 + int(math.Max(float64(dl), float64(dr)))
}
