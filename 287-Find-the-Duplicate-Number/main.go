package p287

func findDuplicate(nums []int) int {
	m := make(map[int]bool, len(nums))
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return num
		} else {
			m[num] = true
		}
	}
	return -1
}
