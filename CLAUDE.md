# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

This repository contains LeetCode problem solutions and algorithm implementations from Sedgewick's "Algorithms" textbook. All code is written in Go (Go 1.23).

## Project Structure

- **Problem directories** (e.g., `1-two-sum/`, `206-Reverse-Linked-List/`): Individual LeetCode solutions. Each directory contains:
  - `main.go`: Solution implementation
  - `main_test.go`: Test cases
  - `readme.md`: Problem description and link
  - Some include `main_bench_test.go` for benchmarks

- **`lib/`**: Shared helper code for common data structures:
  - `tree.go`: Binary tree utilities (`TreeNode`, `BuildTree`, `IsSameTree`)
  - `linked_list.go`: Linked list utilities (`ListNode`, `CreateLinkedList`, etc.)
  - `graph.go`: Graph utilities (`GraphNode`, `CreateGraphFromAdjList`)

- **`000-Sedgewick/`**: Implementations from Sedgewick's "Algorithms" textbook organized by chapter:
  - `1-Fundamentals/unionfind/`: Union-Find data structure implementations (QuickFind, QuickUnion, WeightedQuickUnion)
  - `4-Graphs/`: Graph algorithms

- **`000-*` directories**: Organized study materials:
  - `000-Backtracking/`: Backtracking algorithms
  - `000-Heap/`: Heap data structures
  - `000-Queues/`: Queue implementations
  - `000-Trees/`: Tree algorithms
  - `000-Recursion/`: Recursive solutions

- **`0000-Speed-Drills/`**: Practice implementations of basic data structures

## Creating a New LeetCode Problem

When given a LeetCode problem number, create a new problem directory following these steps:

1. **Create directory**: Use format `{number}-{problem-name}/` (e.g., `206-Reverse-Linked-List/`)

2. **Create `main.go`**:
   - Add package declaration: `package p{problem-number}` (e.g., `package p206`)
   - Import relevant utilities from `lib/` if needed (tree, linked list, graph)
   - Add stub function signature matching the LeetCode problem (DO NOT implement)
   - Include parameter and return types from problem description

3. **Create `readme.md`**:
   - Include problem title, description, and URL
   - Format:
     ```markdown
     ## {Problem Title}

     {Problem URL}

     {Problem Description}
     ```

4. **Create `main_test.go`**:
   - Package: `package p{problem-number}`
   - Use table-driven test structure
   - Generate test cases from LeetCode examples
   - Use `lib.BuildTree()`, `lib.CreateLinkedList()`, etc. for test data setup
   - For tree problems, use `math.MaxInt` to represent nil nodes

5. **Create `notes.md`**:
   - Track analysis and approach development
   - Format:
     ```markdown
     # Notes

     ## Time Complexity
     - TODO: Analyze time complexity
     - Brute force: TODO
     - Optimized: TODO

     ## Space Complexity
     - TODO: Analyze space complexity
     - Brute force: TODO
     - Optimized: TODO

     ## Edge Cases
     - [ ] TODO: Empty input
     - [ ] TODO: Single element
     - [ ] TODO: Duplicates
     - [ ] TODO: Negative numbers
     - [ ] TODO: Maximum constraints
     - [ ] TODO: Minimum constraints

     ## Brute Force Baseline
     TODO: Describe naive approach
     - Algorithm:
     - Time: O(?)
     - Space: O(?)
     - Why it works:
     - Limitations:

     ## Optimizations
     TODO: Better approaches
     ```

6. **Verify setup**:
   - Run `go test -v ./{problem-directory}/` to ensure it compiles
   - Test should fail since function is not implemented yet

## Common Commands

### Running Tests
```bash
# Run all tests in the repository
go test -v ./...

# Run tests for a specific problem
go test -v ./206-Reverse-Linked-List/

# Run a single test
go test -v ./206-Reverse-Linked-List/ -run TestName
```

### Running Benchmarks
```bash
# Run all benchmarks
go test -bench=. ./...

# Run benchmarks for a specific problem
go test -bench=. ./206-Reverse-Linked-List/

# Run benchmarks with memory stats
go test -bench=. -benchmem ./206-Reverse-Linked-List/
```

### Building
```bash
# Build all packages
go build -v ./...
```

### Module Management
```bash
# Download dependencies
go mod download

# Tidy dependencies
go mod tidy
```

## Code Conventions

### Package Naming
Each problem directory uses a package name format: `p{problem-number}` (e.g., `package p206` for problem 206).

### Test Structure
Tests use table-driven design with struct fields matching the problem parameters. Helper functions from `lib/` are used to construct test data structures (e.g., `lib.BuildTree()`, `lib.CreateLinkedList()`).

### Nil Node Representation
For tree problems, use `math.MaxInt` to represent nil nodes in array representations:
```go
[]int{3, 9, 20, math.MaxInt, math.MaxInt, 15, 7}
```

### Dependencies
- `github.com/stretchr/testify`: Assertions and test utilities
- `github.com/google/go-cmp`: Deep equality comparisons in tests

## CI/CD

GitHub Actions workflow (`.github/workflows/go.yml`) runs on push and PR to main:
- Builds all packages
- Runs all tests
