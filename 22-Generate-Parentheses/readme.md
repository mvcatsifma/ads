## 22. Generate Parentheses

https://leetcode.com/problems/generate-parentheses/description/

# Approach 1: Brute Force Parentheses Generator

## Description

A **brute force implementation** for generating valid parentheses combinations that creates all possible permutations of parentheses characters and filters for valid sequences. This initial approach prioritizes simplicity and correctness over performance optimization.

### Algorithm Overview

The implementation follows a straightforward three-step process:
1. **String Construction**: Build base string with n pairs of parentheses `"()()()..."`
2. **Permutation Generation**: Generate all possible character arrangements
3. **Validation & Filtering**: Check each permutation for valid parentheses matching

### Implementation Characteristics

- **Correctness**: Produces accurate results for all valid inputs
- **Simplicity**: Clear, readable logic that's easy to understand and verify
- **Educational Value**: Demonstrates problem requirements and validation logic
- **Brute Force Approach**: Exhaustive search without algorithmic optimizations

### Performance Limitations

- **Factorial Time Growth**: Each increment of n dramatically increases runtime
- **Memory Intensive**: Massive allocation overhead from string operations
- **Wasteful Computation**: Over 99% of generated permutations are discarded
- **Practical Limit**: Becomes impractical for n > 4

### Use Cases

**Suitable for:**
- Small inputs (n ≤ 3) where performance isn't critical
- Learning and educational purposes
- Algorithm benchmarking baseline
- Initial prototyping and problem exploration

**Limitations:**
- Not suitable for production use with larger inputs
- Performance degrades exponentially with input size
- High memory consumption for moderate values of n

This implementation serves as a functional baseline that demonstrates the problem solution through exhaustive search, trading performance for implementation simplicity.

# Approach 2: Generate All Valid Parentheses Combinations

## Description
This function generates all possible combinations of well-formed parentheses for a given number of pairs `n`. It uses a backtracking algorithm to build valid parentheses strings by ensuring that:

1. The total number of opening parentheses never exceeds `n`
2. The number of closing parentheses never exceeds the number of opening parentheses at any point
3. Each complete string contains exactly `2*n` characters (n opening + n closing parentheses)

**Parameters:**
- `n` (int): The number of parentheses pairs to generate

**Returns:**
- `[]string`: A slice containing all valid parentheses combinations

**Example:**
- For `n=2`, returns: `["(())", "()()")`
- For `n=3`, returns: `["((()))", "(()())", "(())()", "()(())", "()()()"]`

The algorithm explores all possible paths through recursive backtracking, pruning invalid branches early to ensure efficiency.

## Performance Analysis

**Time Complexity:** O(4^n / √n)
- The number of valid parentheses combinations follows the nth Catalan number: C(n) = (2n)! / ((n+1)! × n!)
- This asymptotically approaches 4^n / (√π × n^(3/2)), which simplifies to O(4^n / √n)
- Each valid combination requires O(n) time to construct the string

**Space Complexity:** O(4^n / √n)
- **Output space:** Stores all valid combinations, each of length 2n
- **Recursion stack:** O(n) depth due to maximum string length of 2n
- **Working space:** O(n) for building each string during backtracking

**Optimization Notes:**
- **Early pruning:** The algorithm efficiently prunes invalid branches by checking constraints before recursing
- **String concatenation:** Uses string concatenation (s+"(" and s+")") which creates new strings each time - could be optimized using a byte slice or string builder for better memory efficiency
- **Memory allocation:** The result slice grows dynamically, which may cause occasional reallocations

**Scalability:**
- n=1: 1 combination
- n=3: 5 combinations
- n=5: 42 combinations
- n=10: 16,796 combinations
- n=15: 9,694,845 combinations

The exponential growth makes this approach impractical for large values of n (typically n > 15-20).
