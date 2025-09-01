package p160

import "testing"

// BenchmarkGetIntersectionNode benchmarks the performance of the getIntersectionNode function
// across different input sizes of linked lists.
//
// The benchmark test creates linked lists of varying sizes (Small, Medium, Large, XLarge)
// and measures the time taken to find the intersection node, if any.
//
// The test cases cover the following scenarios:
//   - Small: 10 and 15 nodes, with a skip of 5 nodes
//   - Medium: 100 and 150 nodes, with a skip of 50 nodes
//   - Large: 1,000 and 1,500 nodes, with a skip of 500 nodes
//   - XLarge: 10,000 and 15,000 nodes, with a skip of 5,000 nodes
//
// The benchmark results provide insights into the time complexity and scalability of the
// getIntersectionNode function, as well as its memory usage and overall efficiency.
//
// The benchmark test also demonstrates the use of the Go testing framework's built-in
// benchmarking capabilities, including the ability to run sub-benchmarks for different
// input sizes and to profile the code's performance.
//
// To run the benchmark test, use the following command:
//   go test -bench=.
//
// The benchmark results will be displayed, including the number of iterations per second
// and the time taken per iteration for each test case.
func BenchmarkGetIntersectionNode(b *testing.B) {
	// Create test cases with different list sizes
	testCases := []struct {
		name  string
		sizeA int
		sizeB int
		skip  int
	}{
		{"Small", 10, 15, 5},
		{"Medium", 100, 150, 50},
		{"Large", 1000, 1500, 500},
		{"XLarge", 10000, 15000, 5000},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			headA, headB := createListsWithSize(tc.sizeA, tc.sizeB, tc.skip)
			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = getIntersectionNode(headA, headB)
			}
		})
	}
}

func createListsWithSize(sizeA, sizeB, skip int) (*ListNode, *ListNode) {
	headA := createListWithSize(sizeA)
	headB := createListWithSize(sizeB)

	// Connect the lists at the skip position
	currA, currB := headA, headB
	for i := 0; i < skip; i++ {
		currA = currA.Next
		currB = currB.Next
	}
	connectLists(currA, currB)

	return headA, headB
}

func createListWithSize(size int) *ListNode {
	if size == 0 {
		return nil
	}

	head := &ListNode{Val: 0}
	curr := head
	for i := 1; i < size; i++ {
		curr.Next = &ListNode{Val: i}
		curr = curr.Next
	}
	return head
}

func connectLists(headA, headB *ListNode) {
	currA := headA
	for currA.Next != nil {
		currA = currA.Next
	}
	currA.Next = headB
}
