package p703

import "testing"

func TestKthLargest_Add(t *testing.T) {
	type args struct {
		k    int
		nums []int
		vals []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				k:    3,
				nums: []int{4, 5, 8, 2},
				vals: []int{3, 5, 10, 9, 4},
			},
			want: []int{4, 5, 5, 8, 8},
		},
		{
			name: "case 2",
			args: args{
				k:    4,
				nums: []int{7, 7, 7, 7, 8, 3},
				vals: []int{2, 10, 9, 9},
			},
			want: []int{7, 7, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := Constructor(tt.args.k, tt.args.nums)
			for i := 0; i < len(tt.args.vals); i++ {
				val := tt.args.vals[i]
				want := tt.want[i]
				if got := k.Add(val); got != want {
					t.Errorf("Add() = %v, want %v", got, want)
				}
			}
		})
	}
}
