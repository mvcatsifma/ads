package main

func twoSum(nums []int, target int) []int {
	numToIndex := make(map[int][]int)

	for idx, num := range nums {
		if idxs, ok := numToIndex[num]; !ok {
			numToIndex[num] = []int{idx}
		} else {
			numToIndex[num] = append(idxs, idx)
		}
	}

	for num, i0 := range numToIndex {
		if num*2 == target && len(i0) > 1 {
			return []int{i0[0], i0[1]}
		}

		sub := target - num

		if i1, ok := numToIndex[sub]; ok {
			if i0[0] != i1[0] {
				return []int{i0[0], i1[0]}
			}
		}
	}

	return []int{}
}
