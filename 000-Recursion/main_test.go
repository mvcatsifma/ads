package recursion

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_fibonacciTopdown(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "case 1",
			n:    5,
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fibonacciDynamicTopdown(tt.n, make([]int, tt.n+1))
			if !cmp.Equal(got, tt.want) {
				t.Errorf("fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciDynamicTopdown(20, make([]int, 21)) // or any n you want to test
	}
}

func Test_fibonacciDynamicBottomUp(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "case 1",
			n:    5,
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fibonacciDynamicBottomUp(tt.n)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
