package p128

func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}
	longestStreak := 0
	for num := range numSet {
		if _, exists := numSet[num-1]; !exists {
			currentNum := num
			currentStreak := 1
			for {
				if !numSet[currentNum+1] {
					break
				}
				currentNum++
				currentStreak++
			}
			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}
		}
	}
	return longestStreak
}
