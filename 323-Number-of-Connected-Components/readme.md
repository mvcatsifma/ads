## 323. Number of Connected Components in an Undirected Graph

https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph/description/

## Approach

Counted the number of connected components.

Steps:
- build adjacency list
- track visited
- track number of components
- iterate over edges -> DFS
- for each completed DFS call, increase number
- return number