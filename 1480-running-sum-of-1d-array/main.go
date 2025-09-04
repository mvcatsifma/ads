package p1480

func runningSum(nums []int) []int {
	var result []int
	result = append(result, nums[0])
	for i := 1; i < len(nums); i++ {
		result = append(result, nums[i]+result[i-1])
	}
	return result
}
