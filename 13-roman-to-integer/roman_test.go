package main

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Roman III",
			args: args{s: "III"},
			want: 3,
		},
		{
			name: "Roman LVIII",
			args: args{s: "LVIII"},
			want: 58,
		},
		{
			name: "Roman MCMXCIV",
			args: args{s: "MCMXCIV"},
			want: 1994,
		},
		{
			name: "Roman XXXIX",
			args: args{s: "XXXIX"},
			want: 39,
		},
		{
			name: "Roman CCXLVI",
			args: args{s: "CCXLVI"},
			want: 246,
		},
		{
			name: "Roman DCCLXXXIX",
			args: args{s: "DCCLXXXIX"},
			want: 789,
		},
		{
			name: "Roman MMCDXXI",
			args: args{s: "MMCDXXI"},
			want: 2421,
		},
		{
			name: "Roman CLX",
			args: args{s: "CLX"},
			want: 160,
		},
		{
			name: "Roman CCVII",
			args: args{s: "CCVII"},
			want: 207,
		},
		{
			name: "Roman MIX",
			args: args{s: "MIX"},
			want: 1009,
		},
		{
			name: "Roman MLXVI",
			args: args{s: "MLXVI"},
			want: 1066,
		},
		{
			name: "Roman MDCCLXXVI",
			args: args{s: "MDCCLXXVI"},
			want: 1776,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
