package p572

import (
	"leetcode/lib"
)

// IsSubtree determines whether subRoot is a subtree of root.
// A subtree is defined as a tree that consists of a node in root and all of its descendants.
// Two trees are considered the same if they are structurally identical and have the same node values.
//
// Time complexity: O(m*n) where m is the number of nodes in root and n is the number of nodes in subRoot.
// Space complexity: O(max(h1, h2)) where h1 and h2 are the heights of root and subRoot respectively (due to recursion stack).
//
// Returns true if subRoot is a subtree of root, false otherwise.
// Returns true if subRoot is nil (empty tree is subtree of any tree).
// Returns false if root is nil but subRoot is not nil.
func IsSubtree(root *lib.TreeNode, subRoot *lib.TreeNode) bool {
	if root == nil {
		return false
	}
	if subRoot == nil {
		return true
	}

	if isSameTree(root, subRoot) {
		return true
	}

	return IsSubtree(root.Left, subRoot) || IsSubtree(root.Right, subRoot)
}

func isSameTree(p *lib.TreeNode, q *lib.TreeNode) bool {
	if p == nil || q == nil {
		return p == q // true only if both nil
	}

	return p.Val == q.Val &&
		isSameTree(p.Left, q.Left) &&
		isSameTree(p.Right, q.Right)
}
