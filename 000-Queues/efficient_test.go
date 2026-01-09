package p000Queues

import "testing"

func TestEfficientIntQueue_Dequeue(t *testing.T) {
	type fields struct {
		items []int
		head  int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
		want1  bool
	}{
		{
			name: "case 1",
			fields: fields{
				items: []int{1, 2, 3},
				head:  0,
			},
			want:  1,
			want1: true,
		},
		{
			name: "case 2",
			fields: fields{
				items: []int{1, 2, 3},
				head:  1,
			},
			want:  2,
			want1: true,
		},
		{
			name: "case 3",
			fields: fields{
				items: []int{1, 2, 3},
				head:  3,
			},
			want:  0,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &EfficientIntQueue{
				items: tt.fields.items,
				head:  tt.fields.head,
			}
			got, got1 := q.Dequeue()
			if got != tt.want {
				t.Errorf("Dequeue() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Dequeue() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
