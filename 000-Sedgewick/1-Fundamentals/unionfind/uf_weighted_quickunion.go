package unionfind

// WeightedQuickUnion implements the union-find (disjoint-set) data structure with union-by-size optimization.
// It tracks a partition of elements into disjoint sets using a tree-based approach where
// smaller trees are always attached to larger trees to keep trees balanced.
//
// Algorithm: Quick-Union with Union-by-Size (Weighted Quick-Union)
// Each element points to its parent, forming trees. When merging, the smaller tree
// is attached to the root of the larger tree, preventing tall, degenerate trees.
//
// Applications:
//   - Network connectivity (are two computers connected?)
//   - Percolation theory (does water flow through a system?)
//   - Kruskal's minimum spanning tree algorithm
//   - Least common ancestor problems
//   - Image processing (connected components)
//
// Time Complexity (Weighted Quick-Union):
//   - Find: O(log n) - tree height is at most log n
//   - Union: O(log n) - two finds plus O(1) link
//   - Connected: O(log n) - two find operations
//
// Space Complexity: O(n) where n is the number of elements
//
// Improvement over basic Quick-Union: Guarantees logarithmic tree height by
// always attaching smaller tree to larger tree. Can be further optimized with
// path compression for O(α(n)) amortized time, where α is inverse Ackermann.
type WeightedQuickUnion struct {
	id    []int // Parent pointer for each element (id[i] = parent of i)
	sz    []int // Size of tree rooted at i (only meaningful for roots)
	count int   // Number of disjoint components
}

// NewWeightedQuickUnion creates a new Weighted Quick-Union data structure with n elements.
// Initially, each element is in its own component (n separate components).
// Each element is its own root with tree size 1.
//
// Example:
//   uf := NewWeightedQuickUnion(5)  // Creates components {0}, {1}, {2}, {3}, {4}
//                                   // id = [0, 1, 2, 3, 4], sz = [1, 1, 1, 1, 1]
func NewWeightedQuickUnion(n int) *WeightedQuickUnion {
	id := make([]int, n)
	for i := range id {
		id[i] = i // Each element is its own parent (root)
	}
	sz := make([]int, n)
	for i := range sz {
		sz[i] = 1 // Each tree has size 1 initially
	}
	return &WeightedQuickUnion{
		id:    id,
		count: n,
		sz:    sz,
	}
}

// connected returns true if elements p and q are in the same component.
// Compares the roots of the trees containing p and q.
//
// Example:
//   uf.union(0, 1)
//   uf.union(1, 2)
//   uf.connected(0, 2)  // returns true (both have same root)
//   uf.connected(0, 3)  // returns false (different roots)
//
// Time: O(log n) - requires two find operations
func (u *WeightedQuickUnion) connected(p int, q int) bool {
	return u.find(p) == u.find(q)
}

// union merges the components containing elements p and q.
// If p and q are already in the same component, this is a no-op.
// Links the root of the smaller tree to the root of the larger tree.
// Updates the size of the resulting tree.
// Decrements the component count by 1 if a merge occurs.
//
// Example:
//   uf.union(0, 1)  // Tree sizes: 1, 1 → attach either to other → size becomes 2
//   uf.union(2, 3)  // Tree sizes: 1, 1 → attach either to other → size becomes 2
//   uf.union(1, 2)  // Tree sizes: 2, 2 → attach either to other → size becomes 4
//
// Implementation: Always attaches smaller tree to larger tree to maintain balance.
// This ensures tree height never exceeds log n.
//
// Time: O(log n) - two find operations plus O(1) link and size update
func (u *WeightedQuickUnion) union(p int, q int) {
	pID := u.find(p)
	qID := u.find(q)

	if pID == qID {
		return // Already in same component
	}

	// Attach smaller tree to larger tree
	if u.sz[pID] < u.sz[qID] {
		u.id[pID] = qID        // Link root of smaller tree (p) to larger tree (q)
		u.sz[qID] += u.sz[pID] // Update size of resulting tree
	} else {
		u.id[qID] = pID        // Link root of smaller tree (q) to larger tree (p)
		u.sz[pID] += u.sz[qID] // Update size of resulting tree
	}

	u.count--
}

// find returns the root of the tree containing element p.
// Traverses parent pointers upward until reaching a root node.
// A root is identified by id[p] == p (element points to itself).
//
// Example:
//   Tree structure: 0→1→2 (0's parent is 1, 1's parent is 2, 2 is root)
//   uf.find(0)  // returns 2 (follows 0→1→2)
//   uf.find(1)  // returns 2 (follows 1→2)
//   uf.find(2)  // returns 2 (already at root)
//
// Time: O(log n) - tree height is at most log n due to union-by-size
func (u *WeightedQuickUnion) find(p int) int {
	for p != u.id[p] {
		p = u.id[p] // Follow parent pointer upward
	}
	return p
}

// Count returns the number of disjoint components.
// Starts at n (all elements separate) and decreases with each union.
//
// Example:
//   uf := NewWeightedQuickUnion(5)  // count = 5
//   uf.union(0, 1)                  // count = 4
//   uf.union(2, 3)                  // count = 3
//   uf.Count()                      // returns 3
//
// Time: O(1)
func (u *WeightedQuickUnion) Count() int {
	return u.count
}
