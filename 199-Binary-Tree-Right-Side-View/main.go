package p199

import "leetcode/lib"

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
