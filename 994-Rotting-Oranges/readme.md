## 994 Rotting Oranges

https://leetcode.com/problems/rotting-oranges/description/

## Approach

- iterate over the cells
- on timer tick
- for earch cell with a rotting orange
- do a BFS search marking adjacent fresh oranges rotting
- check if all cells have rotting oranges or ar empty
- if false timer++ repeat
- if true done
- edge cases:
- no rotting oranges return -1