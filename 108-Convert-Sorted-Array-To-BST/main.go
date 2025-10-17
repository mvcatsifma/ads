package p108

func sortedArrayToBST(nums []int) *TreeNode {
	var dfs func(l int, r int) *TreeNode
	dfs = func(l int, r int) *TreeNode {
		if l > r {
			return nil
		}
		m := (l + r) / 2
		root := &TreeNode{
			Val: nums[m],
		}

		root.Left = dfs(l, m-1)
		root.Right = dfs(m+1, r)
		return root
	}

	return dfs(0, len(nums)-1)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
