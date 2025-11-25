package p102

import (
	"leetcode/lib"
)

// levelOrder performs a level-order (breadth-first) traversal of a binary tree using DFS recursion.
// Returns a 2D array where each inner array contains all node values at that depth level.
// Example: tree [3,9,20,null,null,15,7] returns [[3],[9,20],[15,7]]
//
// Approach: Uses depth-first recursion with a level parameter to track current depth.
// Dynamically creates a new level slice when reaching a new depth for the first time.
// Time Complexity: O(n) where n is the number of nodes (visit each node once).
// Space Complexity: O(n) for result array + O(h) for recursion stack, where h is tree height.
func levelOrder(root *lib.TreeNode) [][]int {
	levels := [][]int{}
	var helper func(*lib.TreeNode, int)
	helper = func(node *lib.TreeNode, level int) {
		if len(levels) == level {
			levels = append(levels, []int{})
		}
		levels[level] = append(levels[level], node.Val)
		if node.Left != nil {
			helper(node.Left, level+1)
		}
		if node.Right != nil {
			helper(node.Right, level+1)
		}
	}

	if root != nil {
		helper(root, 0)
	}

	return levels
}
