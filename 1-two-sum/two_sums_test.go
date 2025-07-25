package p1

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	var sumTests = []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{"Example 1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"Example 2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"Example 3", []int{3, 3}, 6, []int{0, 1}},
		{"Testcase 1", []int{-3, 4, 3, 90}, 0, []int{0, 2}},
	}

	for _, tt := range sumTests {
		t.Run(tt.name, func(t *testing.T) {
			actual := twoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("expected %#v, got %#v", tt.expected, actual)
			}
		})
	}
}
