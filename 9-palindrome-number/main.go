package p9

import "fmt"

// TODO: Follow up: Could you solve it without converting the integer to a string?
func isPalindrome(x int) bool {

	// Special cases:
	// As discussed above, when x < 0, x is not a palindrome.
	// Also if the last digit of the number is 0, in order to be a palindrome,
	// the first digit of the number also needs to be 0.
	// Only 0 satisfy this property.
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

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

func isPalindrome2(x int) bool {
	// Special cases:
	// As discussed above, when x < 0, x is not a palindrome.
	// Also if the last digit of the number is 0, in order to be a palindrome,
	// the first digit of the number also needs to be 0.
	// Only 0 satisfy this property.
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	// When the length is an odd number, we can get rid of the middle digit by revertedNumber/10
	// For example when the input is 12321, at the end of the while loop we get x = 12, revertedNumber = 123,
	// since the middle digit doesn't matter in palidrome(it will always equal to itself), we can simply get rid of it.
	return x == revertedNumber || x == revertedNumber/10
}
