package p74

// searchMatrix searches for a target value in a 2D matrix by flattening it into
// a 1D slice and performing binary search. Returns true if target is found.
//
// Note: This implementation has O(m*n) space complexity due to flattening.
// For sorted matrices, consider searching directly on the 2D structure for O(1) space.
func searchMatrix(matrix [][]int, target int) bool {
	var ints []int
	for _, v := range matrix {
		for _, i := range v {
			ints = append(ints, i)
		}
	}

	res := search(ints, target)
	if res != -1 {
		return true
	}

	return false
}

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		m := l + (r - l/2)
		if nums[m] > target {
			r = m - 1
		} else if nums[m] < target {
			l = m + 1
		} else {
			return m
		}
	}

	return -1
}
