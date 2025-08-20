package p46

import (
	"reflect"
	"testing"
)

func TestLRUCache(t *testing.T) {
	tests := []struct {
		name       string
		capacity   int
		operations []operation
		expected   []int
	}{
		{
			name:     "basic example from problem",
			capacity: 2,
			operations: []operation{
				{op: "put", key: 1, val: 1},
				{op: "put", key: 2, val: 2},
				{op: "get", key: 1},         // returns 1
				{op: "put", key: 3, val: 3}, // evicts key 2
				{op: "get", key: 2},         // returns -1 (not found)
				{op: "put", key: 4, val: 4}, // evicts key 1
				{op: "get", key: 1},         // returns -1 (not found)
				{op: "get", key: 3},         // returns 3
				{op: "get", key: 4},         // returns 4
			},
			expected: []int{1, -1, -1, 3, 4},
		},
		{
			name:     "testcase from problem",
			capacity: 2,
			operations: []operation{
				{op: "put", key: 2, val: 1},
				{op: "put", key: 1, val: 1},
				{op: "put", key: 2, val: 3},
				{op: "put", key: 4, val: 1},
				{op: "get", key: 1},
				{op: "get", key: 2},
			},
			expected: []int{-1, 3},
		},
		{
			name:     "single capacity",
			capacity: 1,
			operations: []operation{
				{op: "put", key: 1, val: 10},
				{op: "get", key: 1},          // returns 10
				{op: "put", key: 2, val: 20}, // evicts key 1
				{op: "get", key: 1},          // returns -1
				{op: "get", key: 2},          // returns 20
			},
			expected: []int{10, -1, 20},
		},
		{
			name:     "update existing key",
			capacity: 2,
			operations: []operation{
				{op: "put", key: 1, val: 1},
				{op: "put", key: 2, val: 2},
				{op: "put", key: 1, val: 10}, // update key 1
				{op: "get", key: 1},          // returns 10
				{op: "get", key: 2},          // returns 2
			},
			expected: []int{10, 2},
		},
		{
			name:     "get non-existent key",
			capacity: 2,
			operations: []operation{
				{op: "get", key: 1}, // returns -1
				{op: "put", key: 1, val: 1},
				{op: "get", key: 2}, // returns -1
			},
			expected: []int{-1, -1},
		},
		{
			name:     "larger capacity with eviction",
			capacity: 3,
			operations: []operation{
				{op: "put", key: 1, val: 1}, // age: 1
				{op: "put", key: 2, val: 2}, // age: 2
				{op: "put", key: 3, val: 3}, // age: 3
				{op: "get", key: 1},         // returns 1, age: 4 (most recent)
				{op: "get", key: 2},         // returns 2, age: 5 (most recent)
				{op: "get", key: 3},         // returns 3, age: 6 (most recent)
				{op: "put", key: 4, val: 4}, // age: 7, evicts key 1 (oldest access at age 4)
				{op: "get", key: 1},         // returns -1 (evicted)
				{op: "get", key: 2},         // returns 2
				{op: "get", key: 3},         // returns 3
				{op: "get", key: 4},         // returns 4
			},
			expected: []int{1, 2, 3, -1, 2, 3, 4},
		},
		{
			name:     "access pattern affects eviction order",
			capacity: 2,
			operations: []operation{
				{op: "put", key: 1, val: 1}, // age: 1
				{op: "put", key: 2, val: 2}, // age: 2
				{op: "get", key: 1},         // returns 1, age: 3 (refreshes key 1)
				{op: "put", key: 3, val: 3}, // age: 4, evicts key 2 (older than key 1)
				{op: "get", key: 1},         // returns 1 (still exists)
				{op: "get", key: 2},         // returns -1 (evicted)
				{op: "get", key: 3},         // returns 3
			},
			expected: []int{1, 1, -1, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := Constructor(tt.capacity)
			var results []int

			for _, op := range tt.operations {
				switch op.op {
				case "put":
					cache.Put(op.key, op.val)
				case "get":
					result := cache.Get(op.key)
					results = append(results, result)
				}
			}

			if !reflect.DeepEqual(results, tt.expected) {
				t.Errorf("got %v, want %v", results, tt.expected)
			}
		})
	}
}

// Helper struct for test operations
type operation struct {
	op  string
	key int
	val int
}
