package priorityqueus

// IndexMinPQ is an indexed min priority queue implementation using a binary heap.
// It associates integer indices (0 to maxN-1) with keys and supports:
// - Insert a key at a given index
// - Delete the minimum key
// - Change the key at a given index
// - Delete the key at a given index
//
// The data structure uses three arrays:
// - pq[i]: the binary heap (stores client indices, 1-indexed)
// - qp[i]: inverse of pq, qp[pq[i]] = pq[qp[i]] = i (or -1 if i not in queue)
// - keys[i]: the actual keys indexed by the client's index
type IndexMinPQ struct {
	pq   []int // binary heap using 1-based indexing
	qp   []int // inverse: qp[pq[i]] = pq[qp[i]] = i
	n    int   // number of elements on pq
	keys []Key // items with priorities
}

// NewIndexMinPQ creates a new indexed min priority queue with capacity for indices 0 to maxN-1.
func NewIndexMinPQ(maxN int) *IndexMinPQ {
	queue := &IndexMinPQ{
		keys: make([]Key, maxN+1),
		pq:   make([]int, maxN+1),
		qp:   make([]int, maxN+1),
	}
	for i := 0; i <= maxN; i++ {
		queue.qp[i] = -1 // -1 indicates index i is not in the queue
	}
	return queue
}

// insert associates key v with clientIdx.
// Maintains heap order by swimming the new element up to its proper position.
func (m *IndexMinPQ) insert(clientIdx int, v Key) {
	m.n += 1
	m.pq[m.n] = clientIdx
	m.qp[clientIdx] = m.n
	m.keys[clientIdx] = v
	m.swim(m.n)
}

// minKey returns the minimum key in the priority queue.
func (m *IndexMinPQ) minKey() Key {
	return m.keys[m.pq[1]]
}

// delMin removes and returns the index associated with the minimum key.
func (m *IndexMinPQ) delMin() int {
	indexOfMin := m.pq[1]
	m.exch(1, m.n)
	m.n -= 1
	m.sink(1)
	// Clean up: clear the key and mark client index as not in queue
	// Note: pq[m.n+1] is not cleared because n tracks the boundary
	m.keys[m.pq[m.n+1]] = ""
	m.qp[m.pq[m.n+1]] = -1
	return indexOfMin
}

// minIndex returns the index associated with the minimum key.
func (m *IndexMinPQ) minIndex() int {
	return m.pq[1]
}

// changeKey changes the key associated with clientIdx to v.
func (m *IndexMinPQ) changeKey(clientIdx int, v Key) {
	m.keys[clientIdx] = v
	m.swim(m.qp[clientIdx])
	m.sink(m.qp[clientIdx])
}

// delete removes the key associated with clientIdx.
func (m *IndexMinPQ) delete(clientIdx int) {
	heapPos := m.qp[clientIdx]
	m.exch(heapPos, m.n)
	m.n -= 1
	m.swim(heapPos)
	m.sink(heapPos)
	// Clean up: clear the key and mark client index as not in queue
	// Note: pq[m.n+1] is not cleared because n tracks the boundary
	m.keys[clientIdx] = ""
	m.qp[clientIdx] = -1
}

// size returns the number of elements in the priority queue.
func (m *IndexMinPQ) size() int {
	return m.n
}

// isEmpty returns true if the priority queue has no elements.
func (m *IndexMinPQ) isEmpty() bool {
	return m.n == 0
}

// contains returns true if clientIdx is in the priority queue.
func (m *IndexMinPQ) contains(clientIdx int) bool {
	return m.qp[clientIdx] != -1
}

// less compares keys of two heap positions.
// Returns true if key at pq[heapPos1] < key at pq[heapPos2].
// Note: compares keys, not heap indices.
func (m *IndexMinPQ) less(heapPos1 int, heapPos2 int) bool {
	return m.keys[m.pq[heapPos1]] < m.keys[m.pq[heapPos2]]
}

// exch exchanges two elements in the heap and maintains the inverse mapping.
// Swaps pq[heapPos1] and pq[heapPos2], and updates qp to reflect the new positions.
func (m *IndexMinPQ) exch(heapPos1 int, heapPos2 int) {
	m.pq[heapPos1], m.pq[heapPos2] = m.pq[heapPos2], m.pq[heapPos1]
	m.qp[m.pq[heapPos1]] = heapPos1
	m.qp[m.pq[heapPos2]] = heapPos2
}

// swim restores heap order by moving element at heap position k up the tree.
// Used after inserting a new element or decreasing a key.
// Continues until the element reaches the root or is >= its parent.
func (m *IndexMinPQ) swim(k int) {
	for k > 1 && m.less(k, k/2) {
		m.exch(k, k/2)
		k = k / 2
	}
}

// sink restores heap order by moving element at heap position k down the tree.
// Used after removing the minimum element or increasing a key.
// At each step, exchanges with the smaller child and continues until heap order is restored.
func (m *IndexMinPQ) sink(k int) {
	for k*2 <= m.n {
		j := k * 2
		// Choose the smaller of the two children
		if j < m.n && m.less(j+1, j) {
			j++
		}
		// If parent is not greater than smaller child, heap order is satisfied
		if !m.less(j, k) {
			break
		}
		m.exch(k, j)
		k = j
	}
}
