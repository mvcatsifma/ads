package p22

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func Test_generateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "case 1",
			args: args{
				n: 3,
			},
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name: "case 2",
			args: args{
				n: 1,
			},
			want: []string{"()"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateParenthesis(tt.args.n)
			if !cmp.Equal(got, tt.want, cmpopts.SortSlices(func(a, b string) bool {
				return a < b
			})) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}
