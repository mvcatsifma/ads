package main

import "testing"

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isValid("()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}()[]{}")
	}
}

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{s: "()"},
			want: true,
		},
		{
			args: args{s: "()[]{}"},
			want: true,
		},
		{
			args: args{s: "(]"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
