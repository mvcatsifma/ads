## 138. Copy List with Random Pointer

https://leetcode.com/problems/copy-list-with-random-pointer/description/

## Approach: Two-Pass Hash Map Deep Copy of Linked List with Random Pointers

Title: "Two-Pass Hash Map Deep Copy of Linked List with Random Pointers"

Description:
This function creates a deep copy of a linked list where each node contains both a next pointer and a random pointer that can point to any node in the list or be null. The solution uses a two-pass approach with a hash map to maintain node relationships.

Key Features:
1. First pass: Creates new nodes and maps original→copy relationships
2. Second pass: Connects Next and Random pointers using the map
3. Handles nil input and random pointers
4. Preserves original list structure
5. O(n) time and space complexity

Example:
```
Input: 7 -> 13 -> 11 -> 10 -> 1
       ↑    ↑     ↑    ↑    ↑
      nil   7    10   11    7
       
Output: Deep copy with same structure
```

Performance (from benchmarks):
- Linear time scaling (~0.46μs for 10 nodes, ~0.83ms for 10K nodes)
- Consistent memory usage (~90-100 bytes per node)
- Approximately one allocation per node

The solution prioritizes clarity and correctness while maintaining good performance characteristics through efficient hash map usage.

## Approach: Single-Pass Hash Map Deep Copy of Linked List with Random Pointers

Same as above but using a single pass.
