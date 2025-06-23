package main

import (
	"fmt"
	"sort"
)

func main() {
	nested := [][]int{
		{3, 2},
		{1, 4},
		{5, 0},
	}

	// Sort by first element
	sort.Slice(nested, func(i, j int) bool {
		return nested[i][0] < nested[j][0]
	})
	fmt.Printf("Sorted by first element: %v\n", nested)

	// Example 2: Sort by second element
	sort.Slice(nested, func(i, j int) bool {
		return nested[i][1] < nested[j][1]
	})
	fmt.Printf("Sorted by second element: %v\n", nested)

	// Example 3: Sort by sum of elements
	sort.Slice(nested, func(i, j int) bool {
		sumI := nested[i][0] + nested[i][1]
		sumJ := nested[j][0] + nested[j][1]
		return sumI < sumJ
	})
	fmt.Printf("Sorted by sum: %v\n", nested)
}

func topKFrequent(nums []int, topK int) []int {
	var counts = make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}

	var buckets = make([][]int, len(nums)+1)
	for num, freq := range counts {
		if buckets[freq] == nil {
			buckets[freq] = make([]int, 0)
		}
		buckets[freq] = append(buckets[freq], num)
	}

	var result []int
	for i := len(buckets) - 1; i > 0; i-- {
		for _, num := range buckets[i] {
			result = append(result, num)
		}
	}

	return result[:topK]
}
