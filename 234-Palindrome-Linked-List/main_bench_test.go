package main

import (
	"fmt"
	"testing"
)

func BenchmarkIsPalindrome(b *testing.B) {
	testCases := []struct {
		name   string
		values []int
	}{
		{"Small_Palindrome", []int{1, 2, 2, 1}},
		{"Medium_Palindrome", []int{1, 2, 3, 4, 4, 3, 2, 1}},
		{"Large_Palindrome", []int{1, 2, 3, 4, 5, 6, 7, 8, 8, 7, 6, 5, 4, 3, 2, 1}},
		{"Small_NonPalindrome", []int{1, 2, 3, 4}},
		{"Empty", []int{}},
		{"Single", []int{1}},
	}

	for _, tc := range testCases {
		// Create a sub-benchmark for each test case
		b.Run(fmt.Sprintf("Size_%s", tc.name), func(b *testing.B) {
			list := createLinkedList(tc.values)
			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				isPalindrome(list)
			}
		})
	}
}
