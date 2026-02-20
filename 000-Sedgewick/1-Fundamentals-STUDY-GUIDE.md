# Chapter 1: Fundamentals - Study Guide

## For Working Engineers

This chapter covers the foundation you need to understand all other algorithms. Skip the academic theory; focus on what you'll actually use.

## Section 1.1: Basic Programming Model
**Skip this** - You already know how to program. Maybe skim for Sedgewick's specific notation if needed.

## Section 1.2: Data Abstraction
**Skip most of it** - You understand objects and APIs. Key takeaway: think about data structure interfaces vs. implementations.

## Section 1.3: Stacks, Queues, Bags
**Essential - Read thoroughly**

### Why It Matters
- Stacks: Function call stacks, undo/redo, parsing, backtracking
- Queues: Job processing, message queues, BFS traversal
- Bags: When you need to collect items but don't care about order

### Key Implementations
- **Array-based**: Fast, fixed size, need to handle resizing
- **Linked-list-based**: Dynamic, pointer overhead, no resizing issues

### Practical Tips
- Use array-based for performance-critical code
- Use linked-list when you need true O(1) insertions anywhere
- Modern languages often provide both (Go: slice vs. list)

### Real-World Uses
- Stacks: Expression evaluation, browser history, syntax parsing
- Queues: Task schedulers, print spoolers, network buffers
- Deques: Sliding window problems, maintaining maximums

## Section 1.4: Analysis of Algorithms
**Read but don't obsess**

### What You Need
- Understand Big-O notation: O(1), O(log n), O(n), O(n log n), O(n²)
- Know when to care: n < 100? Probably doesn't matter. n > 1M? Better care.
- Best/average/worst case: Know they exist, focus on average

### Skip
- Formal mathematical proofs
- Detailed mathematical analysis
- Tilde notation minutiae

### Practical Rules
- O(n²) is fine for n < 1000
- O(n log n) scales to millions
- O(log n) is effectively free
- Constant factors matter in practice (2n vs. 100n)

## Section 1.5: Union-Find (Case Study)
**Essential - This is gold**

### Why It Matters
You already have implementations in this directory! This is a perfect example of:
- Problem: Connectivity, grouping, finding components
- Evolution: Naive → Better → Optimal
- Trade-offs: Code complexity vs. performance

### Key Insight
Watch the progression:
1. **Quick-Find**: O(1) find, O(n) union - too slow
2. **Quick-Union**: Better but can degenerate
3. **Weighted Quick-Union**: O(log n) - practical winner

### Real-World Applications
- Network connectivity
- Image segmentation
- Kruskal's MST algorithm (coming in Chapter 4)
- Percolation problems
- Social network "friend groups"

### Study Approach
1. Read about the problem
2. Look at your implementations in `1.5-UnionFind/`
3. Compare the benchmark results in `perf.md`
4. Understand *why* weighting helps

### Implementation Tips
- Path compression is an easy optimization (halve path length on each find)
- Union-by-rank is slightly better than union-by-size but more complex
- For most practical problems, weighted quick-union is plenty

## What to Skip in Chapter 1
- Mathematical induction proofs
- Detailed memory models
- Academic programming exercises
- Formal specifications

## What to Focus On
- Union-Find progression (you have implementations!)
- Big-O intuition (not formal derivations)
- Stack/Queue interfaces and when to use them
- Understanding trade-offs

## Study Time Estimate
- Section 1.3: 2-3 hours (code examples, try implementations)
- Section 1.4: 1 hour (skim, focus on Big-O table)
- Section 1.5: 3-4 hours (deep dive, run benchmarks, understand progression)

**Total: 6-8 hours for practical understanding**

## Next Steps
After Chapter 1, you'll have the foundation to understand:
- Why sorting algorithms differ (Chapter 2)
- How search trees balance performance (Chapter 3)
- Graph algorithm efficiency (Chapter 4)
