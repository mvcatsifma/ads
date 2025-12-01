package p208

import "testing"

type Operation struct {
	Method string
	Arg    string
}

type TestCase struct {
	Name       string
	Operations []Operation
	Expected   []interface{}
}

func TestTrieOperations(t *testing.T) {
	tests := []TestCase{
		{
			Operations: []Operation{
				{Method: "insert", Arg: "apple"},
				{Method: "search", Arg: "apple"},
				{Method: "search", Arg: "app"},
				{Method: "startsWith", Arg: "app"},
				{Method: "insert", Arg: "app"},
				{Method: "search", Arg: "app"},
			},
			Expected: []interface{}{nil, true, false, true, nil, true},
		},
		{
			Operations: []Operation{
				{Method: "insert", Arg: "hello"},
				{Method: "search", Arg: "hello"},
				{Method: "search", Arg: "hell"},
				{Method: "startsWith", Arg: "hell"},
			},
			Expected: []interface{}{nil, true, false, true},
		},
		{
			Operations: []Operation{
				{Method: "search", Arg: "a"},
				{Method: "startsWith", Arg: "a"},
				{Method: "insert", Arg: "a"},
				{Method: "search", Arg: "a"},
			},
			Expected: []interface{}{false, false, nil, true},
		},
	}

	for i, tc := range tests {
		trie := Constructor()
		for j, op := range tc.Operations {
			var result interface{}
			switch op.Method {
			case "insert":
				trie.Insert(op.Arg)
				result = nil
			case "search":
				result = trie.Search(op.Arg)
			case "startsWith":
				result = trie.StartsWith(op.Arg)
			}

			if result != tc.Expected[j] {
				t.Errorf("Test %d, Operation %d (%s, %q): got %v, want %v",
					i, j, op.Method, op.Arg, result, tc.Expected[j])
			}
		}
	}
}
