package main

import "fmt"

// TODO: Follow up: Could you solve it without converting the integer to a string?
func isPalindrome(x int) bool {
	xStr := fmt.Sprintf("%d", x)
	var l, r = 0, len(xStr) - 1
	for l < r {
		if xStr[l] != xStr[r] {
			return false
		}
		r--
		l++
	}

	return true
}
