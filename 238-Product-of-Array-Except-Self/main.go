package main

func productExceptSelf(nums []int) []int {
	var leftProduct, rightProduct []int
	leftProduct = make([]int, len(nums))
	leftProduct[0] = 1
	for i := 1; i < len(nums); i++ {
		leftProduct[i] = leftProduct[i-1] * nums[i-1]
	}

	rightProduct = make([]int, len(nums))
	rightProduct[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		rightProduct[i] = rightProduct[i+1] * nums[i+1]
	}

	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = leftProduct[i] * rightProduct[i]
	}

	return result
}
