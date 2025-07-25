package lib

// ListNode represents a node in a singly-linked list.
// Each node contains an integer value and a pointer to the next node.
type ListNode struct {
	// Val holds the node's integer value
	Val int

	// Next points to the next node in the list
	// or is nil if this is the last node
	Next *ListNode
}

// CreateLinkedList creates a singly-linked list from a slice of integers.
// It returns nil if the input slice is empty, otherwise returns a pointer
// to the head of the newly created list.
//
// Example:
//
//	nums := []int{1, 2, 3}
//	list := CreateLinkedList(nums) // creates: 1->2->3
func CreateLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	return CreateLinkedListFromNums(nums, -1)
}

// CreateLinkedListFromNums creates a linked list from a slice of integers,
// optionally creating a cycle at the specified position.
//
// Parameters:
//   - nums: slice of integers to create nodes from
//   - pos: position to create cycle (-1 for no cycle, 0-based index otherwise)
//
// Returns:
//   - *ListNode: head of the created linked list
//
// Example:
//   nums=[3,2,0,-4], pos=1 creates: 3->2->0->-4->2->...
//   nums=[1,2], pos=-1 creates: 1->2->nil
func CreateLinkedListFromNums(nums []int, pos int) *ListNode {
	// Handle empty input
	if len(nums) == 0 {
		return nil
	}

	// Create head node
	head := &ListNode{Val: nums[0]}
	current := head

	var tailNext *ListNode
	if pos == 0 {
		tailNext = head
	}

	// Iterate through remaining numbers
	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
		if i == pos {
			tailNext = current
		}
	}

	if tailNext != nil {
		current.Next = tailNext
	}

	return head
}

// CreateLinkedListWithSize generates a linked list of specified size.
// If withCycle is true, it creates a cycle by connecting the last node
// to the middle node of the list. If withCycle is false, creates a regular
// linear linked list.
//
// Parameters:
//   - size: number of nodes in the list
//   - withCycle: whether to create a cycle in the list
//
// Returns:
//   - *ListNode: head of the created linked list
func CreateLinkedListWithSize(size int, withCycle bool) *ListNode {
	if size == 0 {
		return nil
	}

	head := &ListNode{Val: 0}
	current := head
	var cyclePoint *ListNode

	// Create the list
	for i := 1; i < size; i++ {
		current.Next = &ListNode{Val: i}
		current = current.Next
		// Store middle node for creating cycle
		if i == size/2 {
			cyclePoint = current
		}
	}

	// Create cycle if requested
	if withCycle {
		current.Next = cyclePoint
	}

	return head
}
