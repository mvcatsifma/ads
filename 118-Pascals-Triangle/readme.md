## 118. Pascal's Triangle

https://leetcode.com/problems/pascals-triangle/description/

## Description

Generates Pascal's triangle with the specified number of rows. Pascal's triangle is a triangular array of numbers where each element is the sum of the two numbers directly above it. The triangle starts with 1 at the apex, and each row begins and ends with 1.

### Example Output
For `numRows = 5`:
```
    1
   1 1
  1 2 1
 1 3 3 1
1 4 6 4 1
```

### Algorithm
- **Row 0**: Contains only `[1]`
- **Subsequent rows**: Edge elements are always `1`, interior elements are calculated as `triangle[row-1][col-1] + triangle[row-1][col]`

### Complexity
- **Time**: O(numRows²) - generates all elements in the triangle
- **Space**: O(numRows²) - stores the complete triangle

### Edge Cases
- Returns empty slice `[][]int{}` when `numRows` is 0
- Single row `[[1]]` when `numRows` is 1

### Optimizations
- Pre-sets edge elements to `1` to avoid redundant calculations
- Direct array access eliminates bounds checking overhead
- Efficient memory allocation with known row sizes