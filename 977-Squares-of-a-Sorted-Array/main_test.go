package p977

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedSquares(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "example 1: mixed negative and positive",
			nums:     []int{-4, -1, 0, 3, 10},
			expected: []int{0, 1, 9, 16, 100},
		},
		{
			name:     "example 2: mixed negative and positive",
			nums:     []int{-7, -3, 2, 3, 11},
			expected: []int{4, 9, 9, 49, 121},
		},
		{
			name:     "all negative numbers",
			nums:     []int{-5, -4, -3, -2, -1},
			expected: []int{1, 4, 9, 16, 25},
		},
		{
			name:     "all positive numbers",
			nums:     []int{1, 2, 3, 4, 5},
			expected: []int{1, 4, 9, 16, 25},
		},
		{
			name:     "single element",
			nums:     []int{-3},
			expected: []int{9},
		},
		{
			name:     "contains zero",
			nums:     []int{-2, 0, 2},
			expected: []int{0, 4, 4},
		},
		{
			name:     "all zeros",
			nums:     []int{0, 0, 0},
			expected: []int{0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sortedSquares(tt.nums)
			assert.Equal(t, tt.expected, result)
		})
	}
}
