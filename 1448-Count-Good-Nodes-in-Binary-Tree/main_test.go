package p1448

import (
	"math"
	"testing"

	"leetcode/lib"
)

func Test_goodNodes(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				root: []int{3, 1, 4, 3, math.MaxInt, 1, 5},
			},
			want: 4,
		},
		{
			name: "case 2",
			args: args{
				root: []int{3, 3, math.MaxInt, 4, 2},
			},
			want: 3,
		},
		{
			name: "case 3",
			args: args{
				root: []int{1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := lib.BuildTree(tt.args.root)
			if got := goodNodes(root); got != tt.want {
				t.Errorf("goodNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
