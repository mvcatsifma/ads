package p199

import "leetcode/lib"

// rightSideView returns the right side view of a binary tree.
// From the perspective of looking at the tree from the right, returns the rightmost node at each level.
// Example: tree [1,2,3,null,5,null,4] returns [1,3,4] (rightmost nodes at each level)
//
// Approach: Uses DFS traversal with level tracking. At each level, overwrites the value with the latest
// (rightmost) node encountered. This works because DFS processes nodes left-to-right, so the last node
// processed at each level is the rightmost visible node.
// Time Complexity: O(n) where n is the number of nodes (visit each node once).
// Space Complexity: O(n) for result array + O(h) for recursion stack, where h is tree height.
func rightSideView(root *lib.TreeNode) []int {
	retval := []int{}
	var helper func(*lib.TreeNode, int)
	helper = func(node *lib.TreeNode, level int) {
		if len(retval) == level {
			retval = append(retval, 0)
		}
		retval[level] = node.Val
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

	return retval
}
