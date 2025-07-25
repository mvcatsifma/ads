package p15

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	triplets := make(map[string][]int)
	dups := make(map[int]bool)
	seen := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if ok := dups[nums[i]]; !ok {
			dups[nums[i]] = true
			for j := i + 1; j < len(nums); j++ {
				complement := -nums[i] - nums[j]
				if v, ok := seen[complement]; ok && v == i {
					triplet := []int{nums[i], nums[j], complement}
					sort.Ints(triplet)
					key := sliceToString(triplet)
					triplets[key] = triplet

				}
				seen[nums[j]] = i
			}
		}
	}

	res := make([][]int, 0)
	for _, v := range triplets {
		res = append(res, v)
	}
	return res
}

func sliceToString(s []int) string {
	return fmt.Sprint(s)
}
