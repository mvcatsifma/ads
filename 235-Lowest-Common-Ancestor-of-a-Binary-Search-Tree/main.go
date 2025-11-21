package p235

import (
	"leetcode/lib"
)

// lowestCommonAncestor finds the lowest common ancestor of two nodes in a binary search tree.
// The LCA is the deepest node that has both p and q as descendants (a node can be its own ancestor).
// Leverages BST property: if both nodes are greater than root, search right; if both are less, search left;
// otherwise, the current node is the LCA.
//
// Time Complexity: O(h) where h is the height of the tree.
//   - Balanced tree: O(log n)
//   - Skewed tree: O(n)
// Space Complexity: O(h) for recursion stack.
func lowestCommonAncestor(root, p, q *lib.TreeNode) *lib.TreeNode {
	parentVal := root.Val
	pVal := p.Val
	qVal := q.Val

	if pVal > parentVal && qVal > parentVal {
		return lowestCommonAncestor(root.Right, p, q)
	} else if pVal < parentVal && qVal < parentVal {
		return lowestCommonAncestor(root.Left, p, q)
	} else {
		return root // the LCA node
	}
}
