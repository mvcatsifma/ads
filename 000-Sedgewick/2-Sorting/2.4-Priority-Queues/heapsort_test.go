package priorityqueus

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeapsort_EmptyArray(t *testing.T) {
	// arr[0] is unused, no actual elements
	arr := []int{0}
	sort(arr)
	expected := []int{0}
	assert.Equal(t, expected, arr)
}

func TestHeapsort_SingleElement(t *testing.T) {
	// arr[0] unused, arr[1] is the only element
	arr := []int{0, 42}
	sort(arr)
	expected := []int{0, 42}
	assert.Equal(t, expected, arr)
}

func TestHeapsort_TwoElements(t *testing.T) {
	arr := []int{0, 5, 3}
	sort(arr)
	expected := []int{0, 3, 5}
	assert.Equal(t, expected, arr)
}

func TestHeapsort_AlreadySorted(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5}
	sort(arr)
	expected := []int{0, 1, 2, 3, 4, 5}
	assert.Equal(t, expected, arr)
}

func TestHeapsort_ReverseSorted(t *testing.T) {
	arr := []int{0, 9, 7, 5, 3, 1}
	sort(arr)
	expected := []int{0, 1, 3, 5, 7, 9}
	assert.Equal(t, expected, arr)
}

func TestHeapsort_WithDuplicates(t *testing.T) {
	arr := []int{0, 5, 2, 8, 2, 9, 1, 5, 5}
	sort(arr)
	expected := []int{0, 1, 2, 2, 5, 5, 5, 8, 9}
	assert.Equal(t, expected, arr)
}

func TestHeapsort_AllSameElements(t *testing.T) {
	arr := []int{0, 7, 7, 7, 7, 7}
	sort(arr)
	expected := []int{0, 7, 7, 7, 7, 7}
	assert.Equal(t, expected, arr)
}

func TestHeapsort_RandomArray(t *testing.T) {
	arr := []int{0, 3, 7, 1, 4, 6, 5, 2, 8}
	sort(arr)
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	assert.Equal(t, expected, arr)
}

func TestHeapsort_NegativeNumbers(t *testing.T) {
	arr := []int{0, -5, 3, -1, 7, -8, 2}
	sort(arr)
	expected := []int{0, -8, -5, -1, 2, 3, 7}
	assert.Equal(t, expected, arr)
}

func TestHeapsort_MixedPositiveNegative(t *testing.T) {
	arr := []int{0, 4, -2, 0, 9, -7, 3, -1, 5}
	sort(arr)
	expected := []int{0, -7, -2, -1, 0, 3, 4, 5, 9}
	assert.Equal(t, expected, arr)
}

func TestHeapsort_LargeArray(t *testing.T) {
	arr := []int{0, 15, 3, 9, 8, 5, 2, 7, 1, 6, 4, 12, 10, 14, 11, 13}
	sort(arr)
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	assert.Equal(t, expected, arr)
}

func TestHeapsort_VerifySortedProperty(t *testing.T) {
	arr := []int{0, 8, 3, 5, 1, 9, 6, 2, 7, 4}
	sort(arr)

	// Verify array is sorted (skip arr[0])
	for i := 1; i < len(arr)-1; i++ {
		assert.LessOrEqual(t, arr[i], arr[i+1], "Array should be sorted: arr[%d]=%d should be <= arr[%d]=%d", i, arr[i], i+1, arr[i+1])
	}
}
