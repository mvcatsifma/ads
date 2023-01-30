package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{0, 1, 2, 2, 3, 0, 4, 2}
	k := removeElement(arr, 2)
	fmt.Println(arr)
	fmt.Println(k)
}

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Slice(nums, func(i, j int) bool {
		if nums[j] == val {
			return true
		}
		return false
	})
	for i, num := range nums {
		if num == val {
			return i
		}
	}
	return len(nums)
}
