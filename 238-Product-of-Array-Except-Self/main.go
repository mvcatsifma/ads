package main

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	for i := 0; i < len(result); i++ {
		product := 1
		for j := 0; j < len(nums); j++ {
			if j == i {
				continue
			}
			product = product * nums[j]
		}
		result[i] = product
	}
	return result
}
