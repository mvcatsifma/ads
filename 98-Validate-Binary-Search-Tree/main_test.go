package p98

import (
	"math"
	"testing"

	"leetcode/lib"
)

func Test_isValidBST(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				root: []int{2, 1, 3},
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				root: []int{5, 1, 4, math.MaxInt, math.MaxInt, 3, 6},
			},
			want: false,
		},
		{
			name: "case 3",
			args: args{
				root: []int{5, 4, 6, math.MaxInt, math.MaxInt, 3, 7},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := lib.BuildTree(tt.args.root)
			if got := isValidBST(root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
