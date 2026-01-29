## 684. Redundant Connection

https://leetcode.com/problems/redundant-connection/description/

## Approach

Detect cycle, return pair of v and w.

- build subslices of edges, removing one edge each time starting from the back
- for that edge list build adjacency list
- check if tree has cycle
- if not return removed edge
- if yes continue with next edge list
