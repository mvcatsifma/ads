package p130

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_solve(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{
			name: "case 1",
			args: args{
				board: [][]byte{
					{'X', 'X', 'X', 'X'},
					{'X', 'O', 'O', 'X'},
					{'X', 'X', 'O', 'X'},
					{'X', 'O', 'X', 'X'},
				},
			},
			want: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
		},
		{
			name: "case 2",
			args: args{
				board: [][]byte{
					{'X'},
				},
			},
			want: [][]byte{
				{'X'},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := tt.args.board
			solve(in)
			if !cmp.Equal(in, tt.want) {
				t.Errorf("solve() = %v, want %v", in, tt.want)
			}
		})
	}
}
