package unionfind

// QuickFind implements the union-find (disjoint-set) data structure.
// It efficiently tracks a partition of elements into disjoint sets and supports
// two primary operations: union (merge sets) and find (identify set membership).
//
// Applications:
//   - Network connectivity (are two computers connected?)
//   - Percolation theory (does water flow through a system?)
//   - Kruskal's minimum spanning tree algorithm
//   - Least common ancestor problems
//   - Image processing (connected components)
//
// Time Complexity (with path compression and union by rank):
//   - Find: O(α(n)) amortized, where α is inverse Ackermann (effectively constant)
//   - Union: O(α(n)) amortized
//   - Connected: O(α(n)) amortized
//
// Space Complexity: O(n) where n is the number of elements
type QuickFind struct {
	id    []int // Component identifier for each element (id[i] = parent of i)
	count int   // Number of disjoint components
}

// NewQuickFind creates a new union-find data structure with n elements.
// Initially, each element is in its own component (n separate components).
//
// Example:
//   uf := NewQuickFind(5)  // Creates components {0}, {1}, {2}, {3}, {4}
func NewQuickFind(n int) *QuickFind {
	id := make([]int, n)
	for i := range id {
		id[i] = i // Each element starts in its own component
	}
	return &QuickFind{
		id:    id,
		count: n,
	}
}

// connected returns true if elements p and q are in the same component.
// In Quick-Find, this is a simple comparison of component identifiers.
//
// Example:
//   uf.union(0, 1)
//   uf.union(1, 2)
//   uf.connected(0, 2)  // returns true (same component id)
//   uf.connected(0, 3)  // returns false (different component ids)
//
// Time: O(1) - constant time comparison
func (u *QuickFind) connected(p int, q int) bool {
	return u.id[p] == u.id[q]
}

// union merges the components containing elements p and q.
// If p and q are already in the same component, this is a no-op.
// Updates all elements in p's component to have q's component identifier.
// Decrements the component count by 1 if a merge occurs.
//
// Example:
//   uf.union(0, 1)  // Merges components {0} and {1} into {0,1}
//   uf.union(2, 3)  // Merges components {2} and {3} into {2,3}
//   uf.union(1, 2)  // Merges {0,1} and {2,3} into {0,1,2,3}
//
// Implementation: Scans entire array and updates all elements with pID to qID.
// This ensures all elements in the merged component share the same identifier,
// enabling O(1) find operations at the cost of O(n) union operations.
//
// Time: O(n) - must scan entire array to update component identifiers
func (u *QuickFind) union(p int, q int) {
	pID := u.find(p)
	qID := u.find(q)

	if pID == qID {
		return
	}

	for i, v := range u.id {
		if v == pID {
			u.id[i] = qID
		}
	}
	u.count--
}

// find returns the component identifier for element p.
// In Quick-Find, all elements in the same component share the same identifier.
// This is a direct array lookup with no tree traversal needed.
//
// Example:
//   uf.union(0, 1)  // After union, id[0] == id[1]
//   uf.union(1, 2)  // After union, id[0] == id[1] == id[2]
//   uf.find(0)      // returns component id (e.g., 2)
//   uf.find(2)      // returns same id (same component)
//   uf.find(3)      // returns different id (different component)
//
// Time: O(1) - constant time array access
func (u *QuickFind) find(p int) int {
	return u.id[p]
}
