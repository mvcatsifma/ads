package p100

import (
	"math"
	"testing"

	"leetcode/lib"
)

func Test_isSameTree(t *testing.T) {
	type args struct {
		p []int
		q []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				p: []int{1, 2, 3},
				q: []int{1, 2, 3},
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				p: []int{1, 2},
				q: []int{1, math.MaxInt, 2},
			},
			want: false,
		},
		{
			name: "case 3",
			args: args{
				p: []int{1, 2, 1},
				q: []int{1, 1, 2},
			},
			want: false,
		},
		{
			name: "case 4",
			args: args{
				p: []int{1, -1},
				q: []int{1, math.MaxInt, -1},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := lib.BuildTree(tt.args.p)
			q := lib.BuildTree(tt.args.q)
			if got := isSameTree(p, q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
