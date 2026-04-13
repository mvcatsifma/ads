package p977

import (
	"slices"
)

func sortedSquares(nums []int) []int {
	var result []int
	for _, num := range nums {
		v := num * num
		result = append(result, v)
	}
	slices.Sort(result)
	return result
}
