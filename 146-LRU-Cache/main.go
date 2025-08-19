package p46

import "math"

type LRUCache struct {
	capacity int
	nums     map[int]int
	age      map[int]int
	curr     int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		nums:     make(map[int]int),
		age:      make(map[int]int),
		curr:     0,
	}
}

func (c *LRUCache) Get(key int) int {
	if v, ok := c.nums[key]; ok {
		c.curr++
		c.age[key] = c.curr
		return v
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	c.curr++
	c.nums[key] = value
	c.age[key] = c.curr

	if len(c.nums) > c.capacity {
		var l = math.MaxInt
		var e int
		for k, v := range c.age {
			if v < l {
				l = v
				e = k
			}
		}
		delete(c.nums, e)
		delete(c.age, e)
	}
}
