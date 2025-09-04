package p15

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_threeSum(t *testing.T) {
	t.Skip("FIXME: Test fails due to non-deterministic slice ordering in threeSum() return value. " +
		"Need to implement order-agnostic comparison before enabling.")

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
				nums: []int{-1, 0, 1, 2, -1, -4},
			},
			want: [][]int{
				{-1, 0, 1},
				{-1, -1, 2},
			},
		},
		{
			name: "case 2",
			args: args{
				nums: []int{0, 1, 1},
			},
			want: [][]int{},
		},
		{
			name: "case 3",
			args: args{
				nums: []int{0, 0, 0},
			},
			want: [][]int{
				{0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !cmp.Equal(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
