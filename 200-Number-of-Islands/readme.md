## 200. Number of Islands

https://leetcode.com/problems/number-of-islands/description/

## Approach

- iterate over all cells (i,j) starting at (0,0)
- if cell value is 1 start DFS and not seen
- base case value is 0 return true
- if 1 search adjacent cells if not seen
- otherwise continue
- save seen cells