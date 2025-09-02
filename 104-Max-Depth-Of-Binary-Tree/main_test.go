package p104

import (
	"testing"

	"leetcode/lib"
)

func Test_maxDepth(t *testing.T) {
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
				root: []int{3, 9, 20, -1, -1, 15, 7},
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				root: []int{1, -1, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := lib.BuildTree(tt.args.root)
			if got := maxDepth(root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
