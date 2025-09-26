package p875

import "slices"

// minEatingSpeed finds the minimum eating speed (bananas per hour) required
// to finish all piles within h hours using binary search.
//
// The algorithm searches between 1 (minimum possible speed) and max(piles)
// (maximum pile size) to find the smallest speed where all bananas can be
// consumed within the time limit. For each pile, eating time is calculated
// as ceil(pile_size / speed).
//
// Time: O(n * log(max_pile)) where n is len(piles)
// Space: O(1)
func minEatingSpeed(piles []int, h int) int {
	l, r := 1, slices.Max(piles)
	res := r
	for l <= r {
		k := (l + r) / 2
		hours := eatingSpeed(piles, k)
		if hours <= h {
			if k < res {
				res = k
			}
			r = k - 1
		} else {
			l = k + 1
		}
	}

	return res
}

func eatingSpeed(piles []int, speed int) int {
	var h int
	for _, b := range piles {
		h += ceilDiv(b, speed)
	}
	return h
}

func ceilDiv(a, b int) int {
	if a%b == 0 {
		return a / b
	}
	return a/b + 1
}
