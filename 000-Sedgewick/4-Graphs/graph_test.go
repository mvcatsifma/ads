package p4Graphs

import (
	"strings"
	"testing"
)

const input1 = `6
8
0 5
2 4
2 3
1 2
0 1
3 4
3 5
0 2`

func TestNewGraphFromReader(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		wantV   int
		wantE   int
		wantErr bool
	}{
		{
			name: "case 1",
			args: args{
				input: input1,
			},
			wantV:   6,
			wantE:   8,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.args.input)
			got, err := NewGraphFromReader(reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGraphFromReader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.V != tt.wantV {
				t.Errorf("NewGraphFromReader() got vertices = %v, want %v", got.V, tt.wantV)
			}
			if got.E != tt.wantE {
				t.Errorf("NewGraphFromReader() got edges = %v, want %v", got.E, tt.wantE)
			}
		})
	}
}
