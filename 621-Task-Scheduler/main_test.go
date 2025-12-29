package p621

import "testing"

func Test_leastInterval(t *testing.T) {
	type args struct {
		tasks []byte
		n     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'},
				n:     2,
			},
			want: 8,
		},
		{
			name: "case 2",
			args: args{
				tasks: []byte{'A', 'C', 'A', 'B', 'D', 'B'},
				n:     1,
			},
			want: 6,
		},
		{
			name: "case 3",
			args: args{
				tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'},
				n:     3,
			},
			want: 10,
		},
		{
			name: "case 4",
			args: args{
				tasks: []byte{'B', 'C', 'D', 'A', 'A', 'A', 'A', 'G'},
				n:     1,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leastInterval(tt.args.tasks, tt.args.n); got != tt.want {
				t.Errorf("leastInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}
