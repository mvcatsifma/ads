package p973

import (
	"reflect"
	"sort"
	"testing"
)

func Test_kClosest(t *testing.T) {
	type args struct {
		points [][]int
		k      int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case 1",
			args: args{
				points: [][]int{
					{1, 3},
					{-2, 2},
				},
				k: 1,
			},
			want: [][]int{{-2, 2}},
		},
		{
			name: "case 2",
			args: args{
				points: [][]int{
					{3, 3},
					{5, -1},
					{-2, 4},
				},
				k: 2,
			},
			want: [][]int{{-2, 4}, {3, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := kClosest(tt.args.points, tt.args.k)
			sortByFirstElement(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}

// sortByFirstElement sorts a 2D slice by the first element of each sub-slice in ascending order.
// Empty sub-slices are moved to the end of the slice.
func sortByFirstElement(data [][]int) {
	sort.Slice(data, func(i, j int) bool {
		// Handle edge case: empty sub-slices should be sorted to the end
		if len(data[i]) == 0 && len(data[j]) == 0 {
			return false
		}
		if len(data[i]) == 0 {
			return false
		}
		if len(data[j]) == 0 {
			return true
		}
		return data[i][0] < data[j][0]
	})
}
