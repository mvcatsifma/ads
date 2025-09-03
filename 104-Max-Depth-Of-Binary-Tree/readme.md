## 104. Max Depth of Binary Tree

https://leetcode.com/problems/maximum-depth-of-binary-tree/description/

## Benchmark Analysis: Same Node Count Comparison

### Performance Impact of Tree Shape

**Clear pattern emerges** - skewed trees are consistently **2.5-3x slower** than balanced trees with the same node count:

| Nodes | Balanced | Skewed | Slowdown |
|-------|----------|--------|----------|
| 100   | 433ns    | 1,073ns| **2.5x** |
| 1,000 | 4,259ns  | 11,308ns| **2.7x** |
| 10,000| 42,652ns | 124,156ns| **2.9x** |

### Why Skewed Trees Are Slower

**Function call overhead dominates:**

```go
// Balanced tree (100 nodes, ~7 depth):
// - ~7 recursive calls per traversal path
// - Multiple nodes processed per call stack level

// Skewed tree (100 nodes, 100 depth):  
// - 100 recursive calls for single path
// - One node per call stack level
```

**Key factors:**
1. **Stack frame overhead** - 100 calls vs 7 calls
2. **CPU cache pressure** - deeper call stack
3. **Branch prediction** - linear vs tree-like branching patterns

### Performance Per Node

| Tree Type | 100 nodes | 1,000 nodes | 10,000 nodes |
|-----------|-----------|-------------|--------------|
| Balanced  | 4.33 ns/node | 4.26 ns/node | 4.27 ns/node |
| Skewed    | 10.73 ns/node | 11.31 ns/node | 12.42 ns/node |

**Observations:**
- **Balanced trees**: Consistent ~4.3ns per node (excellent scaling)
- **Skewed trees**: Degrading performance as depth increases (stack pressure)

### Algorithm Complexity Confirmed

Both scale **linearly with node count** (O(N)), but with different constants:
- **Balanced**: ~4.3ns per node
- **Skewed**: ~11-12ns per node

### Production Implications

1. **Tree balancing matters** for performance-critical applications
2. **Stack overflow risk** - skewed trees will hit recursion limits first
3. **Memory locality** - balanced trees have better cache performance

### Apple M3 Pro Insights

- Handles deep recursion well (10K depth without issues)
- Shows clear performance difference between call patterns
- Consistent linear scaling confirms efficient memory management

**Conclusion:** While both are O(N), tree shape significantly impacts the constant factor - balanced trees are ~3x more efficient for the same work.