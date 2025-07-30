## 2. Add Two Numbers

https://leetcode.com/problems/add-two-numbers/description/

### Add Two Numbers Represented as Reversed Linked Lists 

#### Description
This algorithm performs addition of two non-negative integers where each number is represented as a linked list with digits stored in reverse order (least significant digit first). The function traverses both lists simultaneously, computing the sum digit by digit while managing carry propagation. It constructs a new linked list representing the result in the same reversed format.

**Algorithm Steps:**
1. Initialize a dummy head node and carry variable
2. Iterate through both input lists until all nodes are processed and no carry remains
3. For each position, sum the current digits from both lists plus any carry
4. Create a new node with the units digit (sum % 10)
5. Update carry for the next iteration (sum / 10)
6. Advance pointers in both input lists if nodes exist
7. Return the result list starting from dummy.Next

**Key Characteristics:**
- **Time Complexity**: O(max(m, n)) where m and n are the lengths of input lists
- **Space Complexity**: O(max(m, n)) for the output list
- **Single-pass algorithm** with constant extra space (excluding output)
- **Handles unequal length inputs** gracefully
- **Manages carry propagation** including final carry digits
- **Preserves input lists** without modification