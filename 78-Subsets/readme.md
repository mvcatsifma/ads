## 78. Subsets

https://leetcode.com/problems/subsets/

## Notes

[Wikipedia - Backtracking](https://en.wikipedia.org/wiki/Backtracking)

```text
procedure backtrack(P, c) is
    if reject(P, c) then return
    if accept(P, c) then output(P, c)
    s ← first(P, c)
    while s ≠ NULL do
        backtrack(P, s)
        s ← next(P, s)
```

## Execution Table

| Step | Call   | Action                | subset before return | Output?     |
| ---- | ------ | --------------------- | -------------------- | ----------- |
| 1    | dfs(0) | include 1             | [1]                  | no          |
| 2    | dfs(1) | include 2             | [1,2]                | no          |
| 3    | dfs(2) | include 3             | [1,2,3]              | no          |
| 4    | dfs(3) | base case             | [1,2,3]              | **[1,2,3]** |
| 5    | dfs(2) | exclude 3 (backtrack) | [1,2]                | no          |
| 6    | dfs(3) | base case             | [1,2]                | **[1,2]**   |
| 7    | dfs(1) | exclude 2 (backtrack) | [1]                  | no          |
| 8    | dfs(2) | include 3             | [1,3]                | no          |
| 9    | dfs(3) | base case             | [1,3]                | **[1,3]**   |
| 10   | dfs(2) | exclude 3 (backtrack) | [1]                  | no          |
| 11   | dfs(3) | base case             | [1]                  | **[1]**     |
| 12   | dfs(0) | exclude 1 (backtrack) | []                   | no          |
| 13   | dfs(1) | include 2             | [2]                  | no          |
| 14   | dfs(2) | include 3             | [2,3]                | no          |
| 15   | dfs(3) | base case             | [2,3]                | **[2,3]**   |
| 16   | dfs(2) | exclude 3 (backtrack) | [2]                  | no          |
| 17   | dfs(3) | base case             | [2]                  | **[2]**     |
| 18   | dfs(1) | exclude 2 (backtrack) | []                   | no          |
| 19   | dfs(2) | include 3             | [3]                  | no          |
| 20   | dfs(3) | base case             | [3]                  | **[3]**     |
| 21   | dfs(2) | exclude 3 (backtrack) | []                   | no          |
| 22   | dfs(3) | base case             | []                   | **[]**      |

## Recursion Tree

                                  dfs(0)
                        /-------------------------\
                include 1                       exclude 1
                  [1]                              []
                 dfs(1)                           dfs(1)
            /---------------\               /---------------\
    include 2           exclude 2     include 2           exclude 2
       [1,2]               [1]           [2]                 []
      dfs(2)             dfs(2)        dfs(2)              dfs(2)
      /------\           /-----\       /-----\             /-----\
 include 3  ex3    include 3  ex3  include 3  ex3     include 3  ex3
  [1,2,3]  [1,2]     [1,3]    [1]    [2,3]     [2]       [3]      []
   dfs(3)  dfs(3)    dfs(3)  dfs(3)  dfs(3)   dfs(3)    dfs(3)   dfs(3)
     |       |         |       |       |        |         |        |
   [1,2,3] [1,2]     [1,3]   [1]     [2,3]    [2]       [3]       []
     out     out       out     out     out      out       out       out
