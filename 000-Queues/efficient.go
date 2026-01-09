package p000Queues

type EfficientIntQueue struct {
	items []int
	head  int // pointer to head of queue
}

func NewEfficientIntQueue() *EfficientIntQueue {
	return &EfficientIntQueue{items: make([]int, 0)}
}

func (q *EfficientIntQueue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *EfficientIntQueue) Dequeue() (int, bool) {
	if len(q.items) == 0 || q.head >= len(q.items) {
		return 0, false // Queue is empty, or head is out of bounds
	}
	item := q.items[q.head]
	q.head++
	return item, true
}

func (q *EfficientIntQueue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *EfficientIntQueue) Size() int {
	return len(q.items)
}
