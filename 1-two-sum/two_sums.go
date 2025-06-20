package main

func twoSum(nums []int, target int) []int {
	var m = make(map[int]int)
	for i, num := range nums {
		sub := target - num
		if j, ok := m[sub]; ok {
			return []int{j, i}
		}
		m[num] = i
	}
	return []int{}
}
