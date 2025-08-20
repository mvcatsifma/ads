## 150. Evaluate Reverse Polish Notation

https://leetcode.com/problems/evaluate-reverse-polish-notation/description/

# Reverse Polish Notation (RPN) Calculator with Fixed-Size Stack

## Description

A high-performance **Reverse Polish Notation (RPN) expression evaluator** implemented in Go using a custom fixed-size integer stack. This implementation efficiently processes postfix mathematical expressions with support for the four basic arithmetic operations: addition, subtraction, multiplication, and division.

### Key Features

- **Stack-Based Evaluation**: Uses a custom fixed-size stack optimized for integer operations
- **Complete RPN Support**: Handles all four basic arithmetic operators (+, -, *, /)
- **Memory Efficient**: Pre-allocates stack space based on input size to minimize allocations
- **Zero Error Handling Overhead**: Assumes well-formed input for maximum performance
- **Integer Arithmetic**: Processes 32-bit signed integers with proper operator precedence

### Algorithm Overview

The evaluator follows the standard RPN algorithm:
1. **Numbers**: Push operands onto the stack
2. **Operators**: Pop two operands, apply operation, push result back
3. **Result**: Final stack element contains the computed value

### Implementation Highlights

- **Custom Stack**: Fixed-size integer stack with O(1) push/pop operations
- **Operator Precedence**: Maintains correct operand order for non-commutative operations (-, /)
- **Efficient Parsing**: Direct string-to-integer conversion using `strconv.ParseInt`
- **Memory Optimization**: Single allocation for the entire stack based on token count

### Usage

```go
// Simple arithmetic: 3 + 4 = 7
tokens := []string{"3", "4", "+"}
result := evalRPN(tokens) // returns 7

// Complex expression: ((15 / (7 - 2)) * 3) - 4 = 5
tokens = []string{"15", "7", "2", "-", "/", "3", "*", "4", "-"}
result = evalRPN(tokens) // returns 5
```

### Performance Characteristics

- **Time Complexity**: O(n) where n is the number of tokens
- **Space Complexity**: O(n) for the stack storage
- **Memory Allocations**: Single allocation for stack creation
- **Cache Friendly**: Contiguous array-based stack implementation

### Stack Operations

- **Push**: O(1) - Add operand to stack top
- **Pop**: O(1) - Remove and return stack top element
- **No Bounds Checking**: Optimized for well-formed RPN expressions

This implementation is ideal for performance-critical applications requiring fast evaluation of mathematical expressions in postfix notation, such as calculators, expression parsers, or mathematical computation engines.