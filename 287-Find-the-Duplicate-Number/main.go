package p287

func findDuplicate(nums []int) int {
	left, right := 1, len(nums)-1

	for left < right {
		mid := left + (right-left)/2
		count := countLessOrEqual(nums, mid)

		if count > mid {
			// Duplicate is in range [left, mid]
			right = mid
		} else {
			// Duplicate is in range [mid+1, right]
			left = mid + 1
		}
	}

	return left
}

func countLessOrEqual(nums []int, target int) int {
	count := 0
	for _, num := range nums {
		if num <= target {
			count++
		}
	}
	return count
}
