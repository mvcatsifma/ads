## 261. Graph Valid Tree

https://leetcode.com/problems/graph-valid-tree/description/

## Approach

Validate given graph is a valid DAG (no cycles, one connected component).

Steps:
- build adjacency list
- DFS
- track number of components
- track cycles
- return true if number of components is 1 and no cycle found