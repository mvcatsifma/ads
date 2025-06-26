package main

import (
	"sort"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 { // return immediately
		return 0
	}

	var longestStreak int
	var updateLongestStreak = func(currentStreak int) {
		if currentStreak > longestStreak {
			longestStreak = currentStreak
		}
	}

	// Sort the input array
	sort.Ints(nums)

	var currentStreak int
	var prev int
	for i, num := range nums {
		if i == 0 { // first element in the input array starts the first streak
			currentStreak++
			prev = num
			continue
		}
		diff := num - prev
		switch diff {
		case 0: // duplicate element
			continue
		case 1: // consecutive element
			currentStreak++
		default: // non-consecutive element, start a new streak
			updateLongestStreak(currentStreak)
			currentStreak = 1
		}
		prev = num
	}
	updateLongestStreak(currentStreak) // don't forget to update for the last streak

	return longestStreak
}
