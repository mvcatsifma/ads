package p108

import (
	"math"
	"testing"

	"leetcode/lib"
)

func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{-10, -3, 0, 5, 9},
			},
			want: []int{0, -10, 5, math.MaxInt, -3, math.MaxInt, 9},
		},
		{
			name: "case 2",
			args: args{
				nums: []int{1, 3},
			},
			want: []int{1, math.MaxInt, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := lib.BuildTree(tt.want)
			got := sortedArrayToBST(tt.args.nums)
			if !lib.IsSameTree(got, want) {
				t.Errorf("sortedArrayToBST() = %v, want %v", got, want)
			}
		})
	}
}
