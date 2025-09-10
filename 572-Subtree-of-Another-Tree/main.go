package p572

import (
	"leetcode/lib"
)

func isSubtree(root *lib.TreeNode, subRoot *lib.TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}

	var same bool
	var dfs func(root *lib.TreeNode)
	dfs = func(root *lib.TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		dfs(root.Right)

		if isSameTree(root, subRoot) {
			same = true
		}
	}

	dfs(root)

	return same
}

func isSameTree(p *lib.TreeNode, q *lib.TreeNode) bool {
	if p == nil || q == nil {
		return p == q // true only if both nil
	}

	return p.Val == q.Val &&
		isSameTree(p.Left, q.Left) &&
		isSameTree(p.Right, q.Right)
}
