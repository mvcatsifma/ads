package main

func maxArea(height []int) int {
	var result int
	for i := 0; i <= len(height)-1; i++ {
		for j := i + 1; j <= len(height)-1; j++ {
			area := min(height[i], height[j]) * (j - i)
			if area > result {
				result = area
			}
		}
	}
	return result
}
