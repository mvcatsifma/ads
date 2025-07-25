## 19. Remove Nth Node From End of List

https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/

## Approach: Two-Pass Linked List Node Removal with Dummy Head

Title: "Two-Pass Linked List Node Removal with Dummy Head"

Algorithm Analysis:

Phase 1: Length Calculation
- First traversal counts total nodes
- Time: O(n) for full list traversal
- Space: O(1) using single pointer

Phase 2: Node Removal
- Second traversal to (length-n) position
- Uses dummy node to handle edge cases (like removing first node)
- Time: O(n) for traversal
- Space: O(1) for dummy node

Total Complexity:
- Time: O(n) - requires two full list traversals
- Space: O(1) - only uses pointer variables and one dummy node

Key Features:
1. Dummy node pattern eliminates edge cases
2. Two-pass approach trades time for simplicity
3. No extra data structures needed
4. In-place modification
5. Handles all valid positions (1 to length)

Potential Optimizations:
- Could be done in one pass using two pointers
- Current approach prioritizes readability over performance

Example Flow:
```
Input: 1->2->3->4->5, n=2
Phase 1: count length=5
Phase 2: move to position 3 (length-n=3)
Result: 1->2->3->5
```

The algorithm favors clarity and robustness over optimal performance, using the dummy head pattern to simplify edge case handling.