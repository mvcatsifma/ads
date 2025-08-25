package p594

import "math"

// findLHS finds the length of the longest harmonious subsequence in the array.
// A harmonious subsequence is defined as a subsequence where the difference
// between the maximum and minimum values is exactly 1.
//
// Algorithm:
// 1. Count frequency of each number using a hash map
// 2. For each number, check if its consecutive number (num+1) exists
// 3. If both numbers exist, their combined frequency forms a harmonious subsequence
// 4. Return the maximum length found among all valid pairs
//
// A subsequence maintains relative order but elements don't need to be contiguous.
// Only consecutive integer pairs (differing by 1) can form harmonious subsequences.
//
// Time Complexity: O(n) - single pass to build map, single pass to find max
// Space Complexity: O(n) - hash map stores at most n unique elements
//
// Examples:
//   findLHS([1,3,2,2,5,2,3,7]) returns 5 (subsequence: [3,2,2,2,3])
//   findLHS([1,2,3,4]) returns 2 (any consecutive pair: [1,2] or [2,3] or [3,4])
//   findLHS([1,1,1,1]) returns 0 (no consecutive numbers exist)
//
// Parameters:
//   nums: input array of integers
//
// Returns: length of the longest harmonious subsequence, or 0 if none exists
func findLHS(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	var result int
	for k := range m {
		if _, ok := m[k+1]; ok {
			result = int(math.Max(float64(result), float64(m[k]+m[k+1])))
		}
	}

	return result
}
