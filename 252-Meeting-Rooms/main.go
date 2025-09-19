package p252

import "sort"

// canAttendMeetings determines if a person can attend all meetings without conflicts.
// Each interval represents a meeting with [start_time, end_time].
// Returns true if no meetings overlap, false otherwise.
//
// Algorithm:
// 1. Sort meetings by start time
// 2. Check if any meeting starts before the previous one ends
//
// Time complexity: O(n log n) due to sorting
// Space complexity: O(1) excluding input
//
// Example:
//   intervals = [[0,30], [5,10], [15,20]] -> false (overlap between [0,30] and [5,10])
//   intervals = [[7,10], [2,4]]           -> true (no overlap after sorting)
func canAttendMeetings(intervals [][]int) bool {
	sortByFirstElement(intervals)
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i+1][0] < intervals[i][1] {
			return false
		}
	}
	return true
}

// sortByFirstElement sorts a 2D slice by the first element of each sub-slice in ascending order.
// Empty sub-slices are moved to the end of the slice.
func sortByFirstElement(data [][]int) {
	sort.Slice(data, func(i, j int) bool {
		// Handle edge case: empty sub-slices should be sorted to the end
		if len(data[i]) == 0 && len(data[j]) == 0 {
			return false
		}
		if len(data[i]) == 0 {
			return false
		}
		if len(data[j]) == 0 {
			return true
		}
		return data[i][0] < data[j][0]
	})
}
