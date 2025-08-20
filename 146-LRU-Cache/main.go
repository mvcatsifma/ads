package p46

type LRUCache struct {
	capacity int
	head     *listNode
	tail     *listNode
	m        map[int]*listNode
}

func Constructor(capacity int) LRUCache {
	head := &listNode{}
	tail := &listNode{}
	head.next = tail
	tail.prev = head

	return LRUCache{
		capacity: capacity,
		head:     head,
		tail:     tail,
		m:        make(map[int]*listNode),
	}
}

func (c *LRUCache) Get(key int) int {
	if n, ok := c.m[key]; ok {
		c.Remove(key)
		c.Add(key, n.value)
		return n.value
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if _, ok := c.m[key]; ok {
		c.Remove(key)
		c.Add(key, value)
	} else {
		c.Add(key, value)
	}

	if len(c.m) > c.capacity {
		delete(c.m, c.head.next.key)
		c.head.next = c.head.next.next
		c.head.next.prev = c.head
	}
}

func (c *LRUCache) Add(key int, value int) {
	n := &listNode{
		key:   key,
		value: value,
		prev:  c.tail.prev,
		next:  c.tail,
	}

	c.tail.prev.next = n
	c.tail.prev = n

	c.m[key] = n
}

func (c *LRUCache) Remove(key int) {
	n := c.m[key]
	n.prev.next = n.next
	n.next.prev = n.prev
}

type listNode struct {
	key   int
	value int
	next  *listNode
	prev  *listNode
}
