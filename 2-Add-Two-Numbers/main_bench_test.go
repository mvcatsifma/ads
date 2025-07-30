package p2

import (
	"testing"

	"leetcode/lib"
)

// Comprehensive benchmark covering all test scenarios
func BenchmarkAddTwoNumbers(b *testing.B) {
	// Define test cases with different scenarios
	testCases := []struct {
		name    string
		digits1 []int
		digits2 []int
	}{
		// Small numbers (1-3 digits)
		{"Small_3digits", []int{2, 4, 3}, []int{5, 6, 4}},
		{"Small_1digit", []int{5}, []int{5}},

		// Medium numbers (5-7 digits)
		{"Medium_5digits", []int{9, 9, 9, 9, 9}, []int{9, 9, 9, 9}},
		{"Medium_7digits", []int{1, 2, 3, 4, 5, 6, 7}, []int{7, 6, 5, 4, 3, 2, 1}},

		// Large numbers (10+ digits)
		{"Large_10digits",
			[]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
			[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}},
		{"Large_15digits",
			make([]int, 15), // Will be filled with 9s
			make([]int, 15)}, // Will be filled with 9s

		// Unequal length lists
		{"Unequal_1vs10", []int{1}, []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}},
		{"Unequal_3vs8", []int{1, 2, 3}, []int{9, 8, 7, 6, 5, 4, 3, 2}},

		// Maximum carry propagation (all 9s)
		{"MaxCarry_5digits", []int{9, 9, 9, 9, 9}, []int{9, 9, 9, 9, 9}},
		{"MaxCarry_8digits", []int{9, 9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9, 9, 9, 9, 9}},

		// With zeros
		{"WithZeros_leading", []int{0, 0, 1}, []int{0, 0, 2}},
		{"WithZeros_mixed", []int{1, 0, 2, 0, 3}, []int{4, 0, 5, 0, 6}},

		// Edge cases
		{"EdgeCase_zero", []int{0}, []int{0}},
		{"EdgeCase_single_carry", []int{5}, []int{5}}, // Results in 10 (carry)

		// Various sizes for scaling analysis
		{"Size_20", make([]int, 20), make([]int, 20)},
		{"Size_50", make([]int, 50), make([]int, 50)},
		{"Size_100", make([]int, 100), make([]int, 100)},
	}

	// Initialize slices that were created with make()
	for i := range testCases {
		tc := &testCases[i]

		// Fill large digit slices with 9s for maximum carry
		if tc.name == "Large_15digits" {
			for j := range tc.digits1 {
				tc.digits1[j] = 9
				tc.digits2[j] = 9
			}
		}

		// Fill size test cases with alternating patterns
		if tc.name == "Size_20" || tc.name == "Size_50" || tc.name == "Size_100" {
			for j := range tc.digits1 {
				tc.digits1[j] = j % 10
				tc.digits2[j] = (j + 1) % 10
			}
		}
	}

	// Run benchmark for each test case
	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			b.ReportAllocs() // Report memory allocations
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				// Recreate lists for each iteration to avoid mutation
				list1 := lib.CreateLinkedList(tc.digits1)
				list2 := lib.CreateLinkedList(tc.digits2)
				_ = addTwoNumbers(list1, list2)
			}
		})
	}
}
