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