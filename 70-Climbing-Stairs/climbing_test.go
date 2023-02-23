package main

import "testing"

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{n: 0}, want: 0},
		{args: args{n: 2}, want: 2},
		{args: args{n: 3}, want: 3},
		{args: args{n: 4}, want: 5},
		{args: args{n: 5}, want: 8},
		{args: args{n: 45}, want: 1836311903},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
