# Chapter 3: Searching - Study Guide

## For Working Engineers

This chapter is about finding things fast. You'll learn the data structures behind databases, caches, and indexes.

## The Big Picture
- **Section 3.1**: Symbol tables - the interface (think: map, dict, hash table)
- **Section 3.2**: Binary search trees - the classic tree
- **Section 3.3**: Balanced trees - keeping performance reliable
- **Section 3.4**: Hash tables - the practical winner
- **Section 3.5**: Applications - when to use what

## Section 3.1: Symbol Tables
**Essential - The foundation**

### What Is This?
Symbol table = key-value store = map = dictionary = associative array

**Operations:**
- Put(key, value)
- Get(key) → value
- Delete(key)
- Contains(key)

### Why It Matters
This is one of the most used data structures in all of software:
- Databases: indexes
- Compilers: symbol tables
- Caches: cache lookups
- Configuration: key-value stores

### Implementations Preview
1. **Sequential search** (linked list): O(n) - too slow
2. **Binary search** (sorted array): O(log n) search, O(n) insert - read-heavy only
3. **Binary search trees**: O(log n) average - but can degenerate
4. **Balanced BSTs**: O(log n) guaranteed
5. **Hash tables**: O(1) average - the practical winner

### Skip
- Sequential search details
- Binary search in arrays (you know this)

### Focus
- Understanding the interface
- Trade-offs between implementations

## Section 3.2: Binary Search Trees (BST)
**Important - The gateway to balanced trees**

### What You Need to Know
A BST maintains this property:
- Left subtree: all keys < node's key
- Right subtree: all keys > node's key

**Operations:**
- Search: O(log n) average, O(n) worst case
- Insert: O(log n) average, O(n) worst case
- Delete: Tricky but doable

### Why Study This?
- Foundation for all balanced tree variants
- Simple to implement
- Good for understanding tree traversals
- Interview staple

### The Problem
BSTs can become unbalanced (degenerate into linked list):
```
Insert: 1, 2, 3, 4, 5 → becomes a linked list!
    1
     \
      2
       \
        3
```

### Tree Traversals (Essential)
- **In-order**: Left → Node → Right (gives sorted order!)
- **Pre-order**: Node → Left → Right
- **Post-order**: Left → Right → Node
- **Level-order**: BFS (breadth-first)

### Practical Applications
- Range queries: Find all keys in [low, high]
- Rank: How many keys less than k?
- Floor/Ceiling: Largest key ≤ k / Smallest key ≥ k

### When to Use
- Educational purposes
- When you control insertion order (can keep balanced)
- When you need ordered traversal

### When to Avoid
- Random insertions (use balanced tree or hash table)
- Performance-critical code without balancing

## Section 3.3: Balanced Search Trees
**Essential for understanding, but use libraries**

### Red-Black Trees
Most common balanced BST:
- Java's TreeMap
- C++ std::map
- Linux kernel

**Guarantees:**
- O(log n) search, insert, delete
- Self-balancing through rotations
- At most 2x the height of perfect BST

### Skip the Details
- Color rules and proofs
- Rotation mechanics
- Implementation details

### What You Need to Know
- Balanced BSTs guarantee O(log n)
- Use your language's built-in (TreeMap, std::map)
- Maintains sorted order (unlike hash table)

### When to Use Balanced BSTs
1. **Need sorted order**: Iterate keys in order
2. **Range queries**: Find all keys in [low, high]
3. **Order statistics**: Find kth largest element
4. **Floor/ceiling operations**: Nearest key ≤ or ≥ k

### When to Use Hash Table Instead
- Don't need sorted order
- Just need fast lookup/insert
- Can tolerate occasional resize

### B-Trees (Brief mention)
- Used in databases and filesystems
- Optimized for disk access
- Wide trees (lots of keys per node)
- Not typically implemented by hand

## Section 3.4: Hash Tables
**Essential - Master this**

### Why This Is King
- O(1) average case for search, insert, delete
- Most standard library "maps" use this
- Practical winner for most use cases

### How It Works
1. **Hash function**: key → integer → index
2. **Collision handling**: What when two keys hash to same index?

### Collision Resolution

#### 1. Separate Chaining (Easier)
- Each bucket is a linked list
- Collisions just add to the list
- Simple, handles clustering well
- Used in: Java HashMap, Go map

#### 2. Open Addressing (More complex)
- All entries in array
- Collision → probe for next open spot
- Linear probing: Try index+1, index+2, ...
- Better cache performance, harder to implement

### Hash Function Requirements
- Deterministic: Same key → same hash
- Uniform distribution: Minimize collisions
- Fast to compute

### Practical Tips
1. **Load factor**: Keep items/capacity < 0.75
2. **Prime table size**: Reduces clustering
3. **Resize**: Double size when load factor exceeded
4. **Don't hash mutable objects**: Key changes → lost in hash table!

### Real-World Applications
- **Caches**: URL → cached response
- **Database indexes**: Fast lookups
- **Deduplication**: Seen this before?
- **Counting**: Word frequencies
- **Memoization**: Cache function results

### When Hash Table Wins
- Fast average-case lookup required
- Don't need sorted order
- Don't need range queries
- Keys have good hash function

### When Hash Table Loses
- Need sorted iteration
- Need range queries
- Need min/max efficiently
- Keys hard to hash

## Section 3.5: Applications
**Read for practical context**

### Key Takeaways

#### Sets
- Set = symbol table with keys only (no values)
- Use for membership testing
- Hash set vs. tree set (same trade-offs)

#### Dictionary Clients
- Spell checkers
- Allowlists/blocklists
- File indexing
- Concordances

#### Indexing
- **Inverted index**: word → list of documents
- Used in search engines
- Hash table of word → postings list

#### Sparse Vectors/Matrices
- Store only non-zero entries
- Symbol table: index → value
- Efficient for sparse data

## Practical Decision Guide

### Need Fast Lookup Only?
→ **Hash Table**
- Go: `map[K]V`
- Python: `dict`
- Java: `HashMap`

### Need Sorted Order?
→ **Balanced BST**
- Go: No built-in (use library or implement)
- Python: No built-in (use `sortedcontainers`)
- Java: `TreeMap`
- C++: `std::map`

### Need Fastest Possible?
→ **Hash Table** with tuning:
- Pre-size to avoid resizing
- Good hash function
- Monitor load factor

### Need Range Queries?
→ **Balanced BST**
- Find all keys in [low, high]
- Fast floor/ceiling
- Order statistics

## What to Skip
- Detailed BST deletion algorithms
- Red-black tree rotation proofs
- Mathematical hash function analysis
- Every hash table variant

## What to Focus On
- Hash table collision handling (separate chaining)
- When to use hash table vs. BST
- Understanding O(1) vs. O(log n) trade-offs
- BST traversals (especially in-order)
- Practical applications

## Study Time Estimate
- Section 3.1: 1 hour (symbol table interface)
- Section 3.2: 3-4 hours (BST implementation, traversals)
- Section 3.3: 1-2 hours (skim balanced trees, use libraries)
- Section 3.4: 3-4 hours (hash table implementation)
- Section 3.5: 1 hour (applications)

**Total: 9-12 hours**

## Exercises Worth Doing
1. Implement simple BST with insert, search, in-order traversal
2. Implement hash table with separate chaining
3. LeetCode problems using hash tables
4. Compare hash table vs. BST performance for your use case

## Next Steps
Chapter 4 (Graphs) will use these data structures for complex algorithms like shortest paths and minimum spanning trees.
