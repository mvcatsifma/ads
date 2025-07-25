package p1863

func subsetXORSum(nums []int) int {
	return -1
}

func dfs(nums []int, subset []int, i int) (result [][]int) {
	if i >= len(nums) {
		result = append(result, subset)
		return
	}

	// decision to include nums[i]
	subset = append(subset, nums[i])
	dfs(nums, subset, i+1)

	// decision not to include nums[i]
	subset = subset[0 : len(subset)-1]
	dfs(nums, subset, i+1)

	return result
}
