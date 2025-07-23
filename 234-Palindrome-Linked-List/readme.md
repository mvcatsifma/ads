## 234. Palindrome Linked List

https://leetcode.com/problems/palindrome-linked-list/description/

## Approach 2 - Fast/Slow Pointer with Partial Reversal

Algorithm Summary:
1. Find Middle Point:
   - Use fast/slow pointer technique
   - Slow moves 1 step, fast moves 2 steps
   - When fast reaches end, slow is at middle
   - Time: O(n/2)

2. Reverse Second Half:
   - Start from middle node (slow.Next)
   - Reverse links in second half of list
   - Time: O(n/2)

3. Compare Halves:
   - First half starts from original head
   - Second half starts from reversed portion
   - Compare values until either pointer reaches null
   - Time: O(n/2)

Technical details:
- Space complexity: O(1) - only uses pointer variables
- Time complexity: O(n)
- No extra data structures needed
- In-place reversal of second half
- Handles edge cases (nil, single node)

Example:
Original:  1 -> 2 -> 2 -> 1
Middle:    1 -> 2 | 2 -> 1
Reversed:  1 -> 2 | 1 -> 2
Compare:   1==1, 2==2 ✓

## Approach 1 - Slice-Based Node Collection with Two-Pointer Comparison

1. Traverse the linked list:
    - Append each node pointer to a slice
    - Continue until reaching a node with nil Next pointer
    - Time: O(n), Space: O(n)

2. Initialize two pointers:
    - i starting at index 0 (beginning of slice)
    - j starting at index len(nodes)-1 (end of slice)

3. While i <= j:
    - Compare node values: nodes[i].Val vs nodes[j].Val
    - If values don't match, return false
    - Increment i, decrement j
    - Time: O(n/2)

4. If loop completes (all values matched), return true

Technical details:
- Space complexity: O(n) where n is number of nodes
- Time complexity: O(n)
- Uses slice of *ListNode pointers, not values
- Single pass through linked list for collection
- Single pass through half the collected nodes for comparison

