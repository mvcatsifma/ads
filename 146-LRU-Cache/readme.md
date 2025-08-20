## 146. LRU Cache

https://leetcode.com/problems/lru-cache/description/

You're absolutely right! A cache miss on `Get()` simply returns -1 without any evictions or insertions. Let me correct that:

# Approach: Optimized LRU Cache with Doubly-Linked List

## Description

A high-performance **Least Recently Used (LRU) Cache** implementation in Go using a doubly-linked list combined with a hash map for O(1) operations. This implementation efficiently maintains insertion order and provides constant-time access, updates, and evictions.

### Key Features

- **O(1) Time Complexity**: All operations (Get, Put, Add, Remove) execute in constant time
- **Doubly-Linked List**: Maintains access order with efficient insertion/deletion at both ends
- **Sentinel Nodes**: Uses dummy head and tail nodes to simplify edge case handling
- **Hash Map Integration**: Direct pointer access to list nodes for instant lookups
- **Memory Efficient**: Minimal overhead with direct node references

### Architecture

The cache combines two data structures for optimal performance:
- **Hash Map (`m`)**: Maps keys to list node pointers for O(1) lookups
- **Doubly-Linked List**: Maintains LRU order with most recent items at the tail
- **Sentinel Nodes**: Dummy head and tail nodes eliminate null pointer checks

```
head <-> [LRU item] <-> [item] <-> [MRU item] <-> tail
```

### Implementation Highlights

- **Smart Eviction**: Automatically removes least recently used items from the head
- **Order Maintenance**: Recently accessed items move to tail (most recent position)
- **Efficient Updates**: Existing keys are removed and re-added to update their position
- **Clean API**: Simple Get/Put interface with standard LRU semantics

### Usage

```go
// Create cache with capacity of 3
cache := Constructor(3)

// Store key-value pairs
cache.Put(1, 100)
cache.Put(2, 200)
cache.Put(3, 300)

// Access updates position (moves to most recent)
value := cache.Get(2) // returns 200, moves key 2 to tail

// Exceeding capacity evicts LRU item
cache.Put(4, 400) // evicts key 1 (least recently used)
```

### Performance Characteristics

- **Get (Cache Hit)**: O(1) - Hash lookup + list reordering
- **Get (Cache Miss)**: O(1) - Hash lookup only, returns -1
- **Put (New Key)**: O(1) - Hash insertion + list addition + potential eviction
- **Put (Existing Key)**: O(1) - Hash lookup + list reordering + value update
- **Space**: O(capacity) - Linear space usage

### Advantages Over Array-Based Implementations

- No linear search for eviction candidates
- Efficient middle-element removal/insertion
- Natural LRU ordering through list structure
- Minimal memory fragmentation
- Optimal for high-frequency access patterns

This implementation is production-ready and suitable for performance-critical applications requiring fast, predictable cache operations.