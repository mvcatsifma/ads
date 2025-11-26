package p199

import "leetcode/lib"

func rightSideView(root *lib.TreeNode) []int {
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

	retval := []int{}
	if root != nil {
		helper(root, 0)
	} else {
		return retval
	}

	for _, level := range levels {
		retval = append(retval, level[len(level)-1])
	}

	return retval
}
