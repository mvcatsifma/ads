package main

import "slices"

func threeSum(nums []int) [][]int {
	var result [][]int
	slices.Sort(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > 0 {
			break
		}
		if i == 0 || nums[i] != nums[i-1] {
			result = append(result, twoSum(nums, i))
		}
	}

	return result
}

func twoSum(nums []int, i int) []int {
	lo := i + 1
	hi := len(nums) - 1

	for lo < hi {
		sum := nums[i] + nums[lo] + nums[hi]
		if sum == 0 {
			return []int{nums[i], nums[lo], nums[hi]}
		}
		if sum < 0 {
			lo++
		}
		if sum > 0 {
			hi--
		}
	}

	return []int{}
}
