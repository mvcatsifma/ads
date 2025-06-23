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

	var pairs = make([][]int, 0, len(counts))
	for num, freq := range counts {
		pairs = append(pairs, []int{num, freq})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] > pairs[j][1] // sort by count
	})

	var result = make([]int, topK)
	for i := 0; i < topK; i++ {
		result[i] = pairs[i][0]
	}

	return result
}
