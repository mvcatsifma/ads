package p90

import (
	"reflect"
	"slices"
	"sort"
	"testing"
)

func Test_subsetsWithDup(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{1, 2, 2},
			},
			want: [][]int{{}, {1}, {2}, {1, 2}, {2, 2}, {1, 2, 2}},
		},
		{
			name: "case 2",
			args: args{
				nums: []int{0},
			},
			want: [][]int{{}, {0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subsetsWithDup(tt.args.nums)
			if !reflect.DeepEqual(sortSubsets(got), tt.want) {
				t.Errorf("subsetsWithDup() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Helper function to sort a 2D slice for comparison (ignores order)
func sortSubsets(subsets [][]int) [][]int {
	// Sort each individual subset
	for _, subset := range subsets {
		slices.Sort(subset)
	}

	// Sort the subsets themselves for consistent comparison
	sort.Slice(subsets, func(i, j int) bool {
		// Compare by length first
		if len(subsets[i]) != len(subsets[j]) {
			return len(subsets[i]) < len(subsets[j])
		}
		// Then compare element by element
		for k := 0; k < len(subsets[i]); k++ {
			if subsets[i][k] != subsets[j][k] {
				return subsets[i][k] < subsets[j][k]
			}
		}
		return false
	})

	return subsets
}
