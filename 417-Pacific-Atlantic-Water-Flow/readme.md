## 417. Pacific Atlantic Water Flow

https://leetcode.com/problems/pacific-atlantic-water-flow/

## Approach

- for each of the cells in:
    * (0,0) - (0, cols-1)
    * (0,0) - (row-1, 0)
    * (0, cols-1) - (rows-1, cols-1)
    * (rows-1, 0) - (rows-1, cols-1)
* do a DFS 
  * explore adjacent cells
  * handle out of bound
  * stop on neighbor lower than current