package p286

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_wallsAndGates(t *testing.T) {
	type args struct {
		rooms [][]int
	}
	tests := []struct {
		name string
		args args
		out  [][]int
	}{
		{
			name: "case 1",
			args: args{
				rooms: [][]int{
					{2147483647, -1, 0, 2147483647},
					{2147483647, 2147483647, 2147483647, -1},
					{2147483647, -1, 2147483647, -1},
					{0, -1, 2147483647, 2147483647},
				},
			},
			out: [][]int{
				{3, -1, 0, 1},
				{2, 2, 1, -1},
				{1, -1, 2, -1},
				{0, -1, 3, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := tt.args.rooms
			wallsAndGates(in)
			if !cmp.Equal(in, tt.out) {
				t.Errorf("wallsAndGates() = %v, want %v", in, tt.out)
			}
		})
	}
}
