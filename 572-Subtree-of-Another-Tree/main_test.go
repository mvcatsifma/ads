package p572

import (
	"math"
	"testing"

	"leetcode/lib"
)

func Test_isSubtree(t *testing.T) {
	type args struct {
		root    []int
		subRoot []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				root:    []int{3, 4, 5, 1, 2},
				subRoot: []int{4, 1, 2},
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				root:    []int{3, 4, 5, 1, 2, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, 0},
				subRoot: []int{4, 1, 2},
			},
			want: false,
		},
		{
			name: "case 3",
			args: args{
				root:    []int{1},
				subRoot: []int{},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := lib.BuildTree(tt.args.root)
			subRoot := lib.BuildTree(tt.args.subRoot)
			if got := IsSubtree(root, subRoot); got != tt.want {
				t.Errorf("IsSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}
