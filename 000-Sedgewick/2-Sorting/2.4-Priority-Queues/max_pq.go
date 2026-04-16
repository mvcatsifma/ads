package priorityqueus

// Key represents the type of elements stored in the priority queue
type Key string

// MaxPQ is a binary max heap implementation.
// The heap is represented as a 1-indexed array where:
// - Root is at index 1
// - Parent of node k is at k/2
// - Children of node k are at 2k and 2k+1
// - Invariant: pq[k] >= pq[2k] and pq[k] >= pq[2k+1] for all k
type MaxPQ struct {
	pq []Key // heap-ordered binary tree stored in array pq[1..n]
	n  int   // number of elements in the priority queue
}

// insert adds a new key to the priority queue.
// Maintains heap order by swimming the new element up to its proper position.
func (m *MaxPQ) insert(v Key) {
	m.n += 1
	m.pq[m.n] = v
	m.swim(m.n)
}

// size returns the number of elements in the priority queue.
func (m *MaxPQ) size() int {
	return m.n
}

// isEmpty returns true if the priority queue has no elements.
func (m *MaxPQ) isEmpty() bool {
	return m.n == 0
}

// less compares two elements in the heap.
// Returns true if pq[i] < pq[j].
func (m *MaxPQ) less(i int, j int) bool {
	return m.pq[i] < m.pq[j]
}

// exch exchanges the elements at indices i and j.
func (m *MaxPQ) exch(i int, j int) {
	m.pq[i], m.pq[j] = m.pq[j], m.pq[i]
}

// swim restores heap order by moving element at index k up the tree.
// Used after inserting a new element that may be larger than its parent.
// Continues until the element reaches the root or is smaller than its parent.
func (m *MaxPQ) swim(k int) {
	for k > 1 && m.less(k/2, k) {
		m.exch(k/2, k)
		k = k / 2
	}
}

// sink restores heap order by moving element at index k down the tree.
// Used after removing the maximum element or when a node becomes smaller than its children.
// At each step, exchanges with the larger child and continues until heap order is restored.
func (m *MaxPQ) sink(k int) {
	for k*2 <= m.n {
		j := k * 2
		// Choose the larger of the two children
		if j < m.n && m.less(j, j+1) {
			j++
		}
		// If parent is not less than larger child, heap order is satisfied
		if !m.less(k, j) {
			break
		}
		m.exch(j, k)
		k = j
	}
}

// NewMaxPQ creates a new max priority queue with capacity for maxN elements.
// The underlying array has size maxN+1 to accommodate 1-based indexing.
func NewMaxPQ(maxN int) *MaxPQ {
	return &MaxPQ{
		pq: make([]Key, maxN+1),
	}
}
