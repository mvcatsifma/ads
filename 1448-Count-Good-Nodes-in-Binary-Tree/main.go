package p1448

import "leetcode/lib"

// goodNodes counts the number of "good" nodes in a binary tree.
// A node is "good" if its value is greater than or equal to all node values on the path from root to that node.
// Example: tree [3,1,4,3,null,1,5] returns 4 (nodes with values 3, 4, 3, 5 are good)
//
// Approach: Uses DFS traversal with a max parameter tracking the maximum value seen so far on the current path.
// When a node's value >= max, it's a good node and max is updated for its children.
// Time Complexity: O(n) where n is the number of nodes (visit each node once).
// Space Complexity: O(h) for recursion stack where h is the height of the tree.
func goodNodes(root *lib.TreeNode) int {
	var retval int
	var helper func(node *lib.TreeNode, max int)
	helper = func(node *lib.TreeNode, max int) {
		if node == nil {
			return
		}
		if node.Val >= max {
			retval++
			max = node.Val
		}

		helper(node.Left, max)
		helper(node.Right, max)
	}
	helper(root, root.Val)

	return retval
}
