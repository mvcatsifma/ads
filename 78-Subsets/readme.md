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