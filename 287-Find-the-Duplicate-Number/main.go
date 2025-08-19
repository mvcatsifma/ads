package p287

func findDuplicate(nums []int) int {
	left, right := 1, len(nums)-1

	for left < right {
		mid := left + (right-left)/2
		count := countLessOrEqual(nums, mid)

		if count > mid {
			right = mid
		} else {
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
