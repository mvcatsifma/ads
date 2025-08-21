## 739. Daily Temperatures

https://leetcode.com/problems/daily-temperatures/description/

## Title
**Daily Temperatures - Next Warmer Day Calculator**

## Description
**Problem**: Given an array of daily temperatures, calculate how many days you have to wait after each day to experience a warmer temperature. If no warmer day exists in the future, return 0 for that day.

**Algorithm**: Monotonic Decreasing Stack
- Uses a stack to maintain indices of temperatures in decreasing order
- When a warmer temperature is found, pop all cooler temperatures from the stack and calculate their waiting days
- Each temperature index is pushed and popped at most once, ensuring O(n) time complexity

**Key Characteristics**:
- **Time Complexity**: O(n) - linear scan with amortized constant stack operations
- **Space Complexity**: O(n) - stack storage and result array
- **Pattern**: Classic monotonic stack problem for "next greater element" variants
- **Applications**: Weather forecasting, resource scheduling, time-series analysis

This algorithm efficiently solves the "next greater element" class of problems by maintaining a monotonic data structure that eliminates the need for nested loops, reducing complexity from O(n²) brute force to O(n) optimal solution.

