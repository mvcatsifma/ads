package p4Graphs

import (
	"strings"
	"testing"
)

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
				input: graphConnected,
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
