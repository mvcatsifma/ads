package p230

import (
	"math"
	"testing"

	"leetcode/lib"
)

func Test_kthSmallest(t *testing.T) {
	type args struct {
		root []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				root: []int{3, 1, 4, math.MaxInt, 2},
				k:    1,
			},
			want: 1,
		},
		{
			name: "case 2",
			args: args{
				root: []int{5, 3, 6, 2, 4, math.MaxInt, math.MaxInt, 1},
				k:    3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := lib.BuildTree(tt.args.root)
			if got := kthSmallest(root, tt.args.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
