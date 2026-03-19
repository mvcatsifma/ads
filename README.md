# leetcode

Personal solutions to LeetCode problems and algorithm implementations from Sedgewick's *Algorithms* textbook, written in Go.

## Structure

- **`{number}-{problem}/`** — LeetCode solutions, each with `main.go`, `main_test.go`, and `readme.md`
- **`000-Sedgewick/`** — Sedgewick textbook implementations, organized by chapter
- **`000-Backtracking/`, `000-Heap/`, `000-Queues/`, `000-Recursion/`, `000-Trees/`** — Topic-based study implementations
- **`0000-Speed-Drills/`** — Practice implementations of core data structures
- **`lib/`** — Shared utilities (binary tree, linked list, graph helpers)

## Running Tests

```bash
# All tests
go test ./...

# Single problem
go test ./206-Reverse-Linked-List/

# With verbose output
go test -v ./206-Reverse-Linked-List/
```

## License

Copyright (c) 2026 Thomas Delnoij. All rights reserved. See [LICENSE](LICENSE).
