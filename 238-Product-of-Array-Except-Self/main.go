package main

func productExceptSelf(nums []int) []int {
	var leftProduct []int
	leftProduct = make([]int, len(nums))
	leftProduct[0] = 1
	for i := 1; i < len(nums); i++ {
		leftProduct[i] = leftProduct[i-1] * nums[i-1]
	}

	result := make([]int, len(nums))
	rightProduct := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] = leftProduct[i] * rightProduct
		rightProduct = rightProduct * nums[i]
	}

	return result
}
