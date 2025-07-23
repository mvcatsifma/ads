## 234. Palindrome Linked List

https://leetcode.com/problems/palindrome-linked-list/description/

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

