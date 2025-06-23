package main

import (
	"reflect"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "case 1",
			args: args{
				strs: []string{"Hello", "World"},
			},
			want: []string{"Hello", "World"},
		},
		{
			name: "case 2",
			args: args{
				strs: []string{"Hello"},
			},
			want: []string{"Hello"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeDecode(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodeDecode() = %v, want %v", got, tt.want)
			}
		})
	}
}
