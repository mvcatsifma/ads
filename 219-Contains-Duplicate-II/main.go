package p219

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]bool)
	for i, num := range nums {
		if ok := m[num]; ok {
			return true
		}
		m[num] = true
		if len(m) > k {
			delete(m, nums[i-k])
		}
	}
	return false
}
