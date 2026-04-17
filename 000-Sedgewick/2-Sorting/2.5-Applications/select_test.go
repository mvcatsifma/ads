package applications

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectKthSmallest_Minimum(t *testing.T) {
	arr := []int{3, 7, 1, 4, 6, 5, 2}
	result := selectKthSmallest(arr, 0)
	assert.Equal(t, 1, result, "k=0 should return minimum element")
}

func TestSelectKthSmallest_Maximum(t *testing.T) {
	arr := []int{3, 7, 1, 4, 6, 5, 2}
	result := selectKthSmallest(arr, len(arr)-1)
	assert.Equal(t, 7, result, "k=len-1 should return maximum element")
}

func TestSelectKthSmallest_Median(t *testing.T) {
	arr := []int{3, 7, 1, 4, 6, 5, 2}
	result := selectKthSmallest(arr, 3)
	assert.Equal(t, 4, result, "k=3 should return 4th smallest (median)")
}

func TestSelectKthSmallest_AllPositions(t *testing.T) {
	arr := []int{5, 2, 8, 1, 9, 3, 7}
	expected := []int{1, 2, 3, 5, 7, 8, 9} // Sorted order

	for k := 0; k < len(arr); k++ {
		// Create fresh copy for each test since selectKthSmallest modifies array
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)

		result := selectKthSmallest(arrCopy, k)
		assert.Equal(t, expected[k], result, "k=%d should return %d", k, expected[k])
	}
}

func TestSelectKthSmallest_AlreadySorted(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	result := selectKthSmallest(arr, 3)
	assert.Equal(t, 4, result)
}

func TestSelectKthSmallest_ReverseSorted(t *testing.T) {
	arr := []int{7, 6, 5, 4, 3, 2, 1}
	result := selectKthSmallest(arr, 3)
	assert.Equal(t, 4, result)
}

func TestSelectKthSmallest_WithDuplicates(t *testing.T) {
	arr := []int{5, 2, 8, 2, 9, 5, 7}
	// Sorted: [2, 2, 5, 5, 7, 8, 9]

	tests := []struct {
		k        int
		expected int
	}{
		{0, 2},
		{1, 2},
		{2, 5},
		{3, 5},
		{4, 7},
		{5, 8},
		{6, 9},
	}

	for _, tt := range tests {
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		result := selectKthSmallest(arrCopy, tt.k)
		assert.Equal(t, tt.expected, result, "k=%d should return %d", tt.k, tt.expected)
	}
}

func TestSelectKthSmallest_AllSameElements(t *testing.T) {
	arr := []int{5, 5, 5, 5, 5}
	result := selectKthSmallest(arr, 2)
	assert.Equal(t, 5, result)
}

func TestSelectKthSmallest_SingleElement(t *testing.T) {
	arr := []int{42}
	result := selectKthSmallest(arr, 0)
	assert.Equal(t, 42, result)
}

func TestSelectKthSmallest_TwoElements(t *testing.T) {
	assert.Equal(t, 3, selectKthSmallest([]int{7, 3}, 0))
	assert.Equal(t, 7, selectKthSmallest([]int{7, 3}, 1))
}

func TestSelectKthSmallest_NegativeNumbers(t *testing.T) {
	arr := []int{-5, 3, -1, 7, -8, 2}
	// Sorted: [-8, -5, -1, 2, 3, 7]

	tests := []struct {
		k        int
		expected int
	}{
		{0, -8},
		{1, -5},
		{2, -1},
		{3, 2},
		{4, 3},
		{5, 7},
	}

	for _, tt := range tests {
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		result := selectKthSmallest(arrCopy, tt.k)
		assert.Equal(t, tt.expected, result, "k=%d should return %d", tt.k, tt.expected)
	}
}

func TestSelectKthSmallest_LargerArray(t *testing.T) {
	arr := []int{15, 3, 9, 8, 5, 2, 7, 1, 6, 4, 12, 10, 14, 11, 13}
	// Sorted: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15]

	// Test a few positions
	assert.Equal(t, 1, selectKthSmallest(copySlice(arr), 0))
	assert.Equal(t, 5, selectKthSmallest(copySlice(arr), 4))
	assert.Equal(t, 8, selectKthSmallest(copySlice(arr), 7))
	assert.Equal(t, 15, selectKthSmallest(copySlice(arr), 14))
}

func TestSelectKthSmallest_OutOfBounds(t *testing.T) {
	// k < 0
	assert.Panics(t, func() {
		selectKthSmallest([]int{1, 2, 3, 4, 5}, -1)
	}, "Should panic when k < 0")

	// k >= len
	assert.Panics(t, func() {
		selectKthSmallest([]int{1, 2, 3, 4, 5}, 5)
	}, "Should panic when k >= len(arr)")
}

func TestSelectKthSmallest_PreservesPartialOrder(t *testing.T) {
	arr := []int{9, 3, 7, 1, 5, 2, 8, 4, 6}
	k := 4 // Find 5th smallest (should be 5)

	result := selectKthSmallest(arr, k)
	assert.Equal(t, 5, result)

	// After quickselect, elements arr[0..k-1] should all be <= arr[k]
	// and elements arr[k+1..n-1] should all be >= arr[k]
	for i := 0; i < k; i++ {
		assert.LessOrEqual(t, arr[i], arr[k], "Elements before k should be <= arr[k]")
	}
	for i := k + 1; i < len(arr); i++ {
		assert.GreaterOrEqual(t, arr[i], arr[k], "Elements after k should be >= arr[k]")
	}
}

// Helper function to copy a slice
func copySlice(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)
	return result
}
