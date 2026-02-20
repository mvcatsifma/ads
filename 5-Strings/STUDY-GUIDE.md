# Chapter 5: Strings - Study Guide

## For Working Engineers

String processing is everywhere: text editors, compilers, search engines, bioinformatics, log parsing. This chapter covers specialized algorithms for working with text.

## Reality Check
**Modern languages have good built-in string libraries.** Study this chapter to:
- Understand what's happening under the hood
- Optimize performance-critical text processing
- Prepare for specialized applications (search engines, DNA sequencing)
- Learn interesting algorithms (tries are cool!)

## Section 5.1: String Sorts
**Skim unless doing specialized string work**

### Key Insight
Strings are different from numbers:
- Variable length
- Can compare character-by-character
- Can use character position for sorting

### Algorithms

#### LSD (Least Significant Digit) Radix Sort
**When to use:**
- All strings same length (e.g., dates, phone numbers)
- Fixed alphabet size (e.g., DNA: ACGT)

**Complexity**: O(WN) where W = string length, N = number of strings
- Faster than comparison-based O(N log N) for short strings

**Practical use**: Database sorting of fixed-width fields

#### MSD (Most Significant Digit) Radix Sort
**When to use:**
- Variable-length strings
- English words, arbitrary text

**Complexity**: O(N + W) best case, O(WN) average

**Watch out**: Recursion overhead, small subarrays

#### Three-Way String Quicksort
**Hybrid approach**: Like quicksort but character-aware

**Advantage**: Handles variable-length strings well
**Complexity**: O(N log N) average

### Practical Takeaway
For general string sorting, use:
1. Built-in sort (probably optimized)
2. Three-way string quicksort if implementing yourself
3. LSD radix sort only for fixed-length, small alphabet

## Section 5.2: Tries
**Essential - Powerful for many applications**

### What Is a Trie?
Tree where:
- Each node represents a character
- Path from root to node spells a string
- Commonly: 26-way tree (one child per letter)

### R-way Trie
```
          root
         /  |  \
        c   s   t
       /    |    \
      a     e     h
     /      |      \
    t       a       e
```
Words: "cat", "sea", "the"

**Operations:**
- Search: O(L) where L = key length
- Insert: O(L)
- Delete: O(L)

**Space**: Can be large - R pointers per node

### Why Tries Matter

**Advantages:**
- Search time independent of dataset size
- Prefix operations (find all words starting with "pre")
- Longest prefix match
- No hash collisions

**Applications:**
1. **Autocomplete**: Find all words with prefix
2. **Spell checkers**: "Did you mean...?"
3. **IP routing**: Longest prefix match
4. **Phone directories**: T9 predictive text
5. **Search suggestions**: Google autocomplete

### Ternary Search Tries (TST)
**Space optimization**: Each node has 3 children (less, equal, greater)

**Trade-off**: Saves space, slightly slower than R-way trie

**When to use**: Large alphabets (Unicode), memory constrained

### Practical Tips
- Use TST for large alphabets
- Compress common paths (Patricia trie, radix tree)
- Store additional data at nodes (word frequency, definitions)

## Section 5.3: Substring Search
**Essential for text processing**

### The Problem
Find pattern P in text T

**Applications:**
- Text editors (Ctrl+F)
- Search engines
- Virus scanners
- DNA sequence matching
- Log file analysis

### Algorithms

#### Brute Force
**Complexity**: O(NM) where N = text length, M = pattern length
**When acceptable**: Short patterns, small texts

#### Knuth-Morris-Pratt (KMP)
**Idea**: Precompute pattern info to avoid backing up in text

**Complexity**: O(N + M) - linear!
**Space**: O(M) for DFA/failure function

**Advantage**: Never backs up in text (good for streaming)
**Disadvantage**: Complex to implement, modest constant factor

**When to use:**
- Can't back up in text (streams, tapes)
- Multiple searches with same pattern
- Theoretical interest

#### Boyer-Moore
**Idea**: Scan pattern from right to left, skip sections when mismatch

**Complexity**: O(N/M) best case (sublinear!), O(NM) worst case
**Typical**: O(N) in practice

**Advantage**: Fast in practice, especially for long patterns
**Disadvantage**: Complex preprocessing

**When to use:**
- Text editors, search utilities
- Large texts, medium patterns
- Real-world text (not worst-case data)

#### Rabin-Karp
**Idea**: Hash pattern, hash text windows, compare hashes

**Complexity**: O(N + M) expected, O(NM) worst case
**Uses**: Rolling hash function

**Advantages:**
- Simple to implement
- Easily extends to multiple pattern search
- Good for DNA/string matching

**When to use:**
- Multiple pattern search
- Need simplicity
- Hash-based approach fits problem

### Practical Advice
1. **Default**: Use your language's built-in (e.g., `strings.Contains` in Go)
2. **Multiple patterns**: Rabin-Karp or Aho-Corasick
3. **Very long texts**: Boyer-Moore
4. **Streaming data**: KMP

## Section 5.4: Regular Expressions
**Read but use libraries**

### What Are Regular Expressions?
Pattern language for text matching:
- `a*`: Zero or more 'a'
- `(a|b)`: 'a' or 'b'
- `.`: Any character
- `[a-z]`: Any lowercase letter

### NFA (Nondeterministic Finite Automaton)
Sedgewick shows how to build NFA from regex:
- States represent positions in pattern
- Transitions on characters or epsilon
- Use graph reachability (BFS/DFS)

### Why Study This?
- Understand what happens when you use regex
- Appreciate complexity (catastrophic backtracking)
- Know when regex is overkill

### Practical Reality
**Use your language's regex library:**
- Go: `regexp` package
- Python: `re` module
- Java: `Pattern`, `Matcher`

**When regex is good:**
- Complex patterns (email validation, parsing)
- Quick prototypes
- Text processing scripts

**When regex is bad:**
- Parsing nested structures (use proper parser)
- Performance-critical paths (can be slow)
- Simple substring search (overkill)

### Regex Performance Tips
1. Anchor patterns (`^`, `$`) when possible
2. Avoid catastrophic backtracking
3. Compile pattern once, reuse
4. Use simple string search for simple patterns

## Section 5.5: Data Compression
**Skim unless specialized need**

### Overview
**Goal**: Represent data with fewer bits

**Two types:**
1. **Lossless**: Original data recoverable (ZIP, gzip)
2. **Lossy**: Approximation acceptable (JPEG, MP3)

### Algorithms (Brief)

#### Run-Length Encoding
Simple: "AAAABBBCC" → "4A3B2C"
**Use**: Images with solid colors, fax machines

#### Huffman Coding
**Idea**: Variable-length codes, frequent characters get shorter codes

**Uses priority queue** (from Chapter 2!)

**Applications**: ZIP, JPEG, MP3
**Complexity**: O(N log N) to build tree

#### LZW Compression
**Idea**: Build dictionary of common patterns

**Used in**: GIF, Unix compress, V.42bis modems

### Practical Takeaway
**Use existing compression libraries:**
- Go: `compress/gzip`, `compress/zlib`
- Python: `gzip`, `zlib`
- Java: `java.util.zip`

**Understand for:**
- Appreciating file format complexity
- Optimizing data transmission
- Interview questions

## What to Skip
- Detailed string sort implementations
- Formal regex engine construction
- Every compression algorithm variant
- Mathematical proofs

## What to Focus On
- Tries for prefix operations
- Understanding substring search trade-offs
- When regex is appropriate
- Boyer-Moore concept (skip on mismatch)
- Huffman coding with priority queue

## Study Time Estimate
- Section 5.1: 1-2 hours (skim, understand radix sort concept)
- Section 5.2: 3-4 hours (implement trie, understand prefix operations)
- Section 5.3: 3-4 hours (understand substring search algorithms)
- Section 5.4: 2 hours (skim regex NFA, focus on practical use)
- Section 5.5: 1-2 hours (skim compression, understand Huffman)

**Total: 10-14 hours**

## Practical Exercises
1. Implement trie for autocomplete
2. Compare substring search algorithms on real text
3. Solve LeetCode trie problems
4. Build simple regex matcher (just `*`, `.`, character match)
5. Implement Huffman coding

## Real Projects to Try
1. **Autocomplete system**: Trie-based suggestions
2. **Text search tool**: Boyer-Moore or built-in
3. **Log parser**: Regex for pattern extraction
4. **Dictionary app**: Trie for word lookup and prefix search

## When Chapter 5 Matters Most
- Building text editors or IDEs
- Search engine development
- Bioinformatics (DNA sequence matching)
- Log processing and analysis
- Network protocols (compression)
- Compiler development

## Final Thoughts
Strings are specialized but important. Focus on:
1. Tries (universally useful)
2. Understanding substring search (even if using built-in)
3. Regex usage (not implementation)
4. When to use specialized vs. general algorithms

Most day-to-day engineering uses built-in string functions, but understanding these algorithms helps you:
- Debug performance issues
- Choose right tool for the job
- Optimize when needed
- Appreciate what your language provides
