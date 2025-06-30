package main

func maxArea(height []int) int {
	var result int

	lo := 0
	hi := len(height) - 1

	for lo < hi {
		area := min(height[lo], height[hi]) * (hi - lo)
		if area > result {
			result = area
		}

		if height[lo] > height[hi] {
			hi--
		} else {
			lo++
		}
	}

	return result
}
