package unionfind

// WeightedQuickUnion implements the union-find (disjoint-set) data structure using a tree-based approach.
// It tracks a partition of elements into disjoint sets and supports
// two primary operations: union (merge sets) and find (identify set membership).
//
// Algorithm: Quick-Union (tree-based, no optimizations)
// Each element points to its parent, forming trees where the root represents the component.
//
// Applications:
//   - Network connectivity (are two computers connected?)
//   - Percolation theory (does water flow through a system?)
//   - Kruskal's minimum spanning tree algorithm
//   - Least common ancestor problems
//   - Image processing (connected components)
//
// Time Complexity (basic Quick-Union without optimizations):
//   - Find: O(tree height) - worst case O(n), average O(log n)
//   - Union: O(tree height) - worst case O(n), average O(log n)
//   - Connected: O(tree height) - worst case O(n), average O(log n)
//
// Space Complexity: O(n) where n is the number of elements
//
// Note: This implementation can be optimized with:
//   - Path compression: flatten tree during find operations
//   - Union by rank/size: attach smaller tree to larger tree
//   With both optimizations: O(α(n)) amortized, where α is inverse Ackermann
type QuickUnion struct {
	id    []int // Parent pointer for each element (id[i] = parent of i)
	count int   // Number of disjoint components
}

// NewQuickUnion creates a new Quick-Union data structure with n elements.
// Initially, each element is in its own component (n separate components).
// Each element is its own root (id[i] = i).
//
// Example:
//   uf := NewQuickUnion(5)  // Creates components {0}, {1}, {2}, {3}, {4}
//                           // Tree structure: 0  1  2  3  4 (all separate roots)
func NewQuickUnion(n int) *QuickUnion {
	id := make([]int, n)
	for i := range id {
		id[i] = i // Each element is its own parent (root)
	}
	return &QuickUnion{
		id:    id,
		count: n,
	}
}

// connected returns true if elements p and q are in the same component.
// Compares the roots of the trees containing p and q.
//
// Example:
//   uf.union(0, 1)
//   uf.union(1, 2)
//   uf.connected(0, 2)  // returns true (both have root 2)
//   uf.connected(0, 3)  // returns false (different roots)
//
// Time: O(tree height) - requires two find operations
func (u *QuickUnion) connected(p int, q int) bool {
	return u.find(p) == u.find(q)
}

// union merges the components containing elements p and q.
// If p and q are already in the same component, this is a no-op.
// Links the root of p's tree to the root of q's tree.
// Decrements the component count by 1 if a merge occurs.
//
// Example:
//   uf.union(0, 1)  // Links root of 0 to root of 1: 0→1
//   uf.union(2, 3)  // Links root of 2 to root of 3: 2→3
//   uf.union(1, 2)  // Links root of {0,1} to root of {2,3}: 1→3
//                   // Final tree: 0→1→3, 2→3
//
// Implementation: Finds roots of both elements and links one root to the other.
// Creates a tree structure where elements point to their parent.
//
// Time: O(tree height) - two find operations plus O(1) link
func (u *QuickUnion) union(p int, q int) {
	pID := u.find(p)
	qID := u.find(q)

	if pID == qID {
		return // Already in same component
	}

	u.id[pID] = qID // Link root of p to root of q
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
// Time: O(tree height) - worst case O(n) for degenerate tree, average O(log n)
func (u *QuickUnion) find(p int) int {
	for p != u.id[p] {
		p = u.id[p] // Follow parent pointer upward
	}
	return p
}

// Count returns the number of disjoint components.
// Starts at n (all elements separate) and decreases with each union.
//
// Example:
//   uf := NewQuickUnion(5)  // count = 5
//   uf.union(0, 1)          // count = 4
//   uf.union(2, 3)          // count = 3
//   uf.Count()              // returns 3
//
// Time: O(1)
func (u *QuickUnion) Count() int {
	return u.count
}
