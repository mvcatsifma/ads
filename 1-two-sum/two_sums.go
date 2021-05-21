package main

func twoSum(nums []int, target int) []int {
	var numToIdx = make(map[int]int)
	for idx2, num := range nums {
		sub := target - num
		if idx1, ok := numToIdx[sub]; ok {
			return []int{idx1, idx2}
		}
		numToIdx[num] = idx2
	}
	return []int{}
}
