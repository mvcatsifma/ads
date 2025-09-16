package p703

import "slices"

type KthLargest struct {
	k    int
	nums []int
}

func Constructor(k int, nums []int) KthLargest {
	return KthLargest{
		k:    k,
		nums: nums,
	}
}

func (k *KthLargest) Add(val int) int {
	k.nums = append(k.nums, val)
	slices.Sort(k.nums)
	return k.nums[len(k.nums)-k.k] // fixme: potential index out of bounds error
}
