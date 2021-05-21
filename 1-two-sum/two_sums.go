package main

func twoSum(nums []int, target int) []int {
	// create hashmap num -> idx
	// iterate over  nums
	// drop nums > target
	// save num + idx
	// for each key, subtract from target, lookup result
	// if found return key idx, result idx
	// if not found, continue

	numToIndex := make(map[int]int)

	for idx, num := range nums {
		if num > target {
			continue
		}

		numToIndex[num] = idx
	}

	for num1, idx1 := range numToIndex {
		diff := target - num1
		if idx2, ok := numToIndex[diff]; ok {
			return []int{idx1, idx2}
		}
	}

	return []int{}
}
