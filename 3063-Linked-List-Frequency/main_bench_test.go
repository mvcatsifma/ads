package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkFrequenciesOfElements(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000, 100000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			list := buildLinkedListWithSize(size)
			b.ResetTimer()
			b.ReportAllocs() // Added allocation reporting for all sizes

			for i := 0; i < b.N; i++ {
				frequenciesOfElements(list)
			}
		})
	}
}

func buildLinkedListWithSize(size int) *ListNode {
	if size <= 0 {
		return nil
	}

	dummy := &ListNode{}
	curr := dummy
	for i := 0; i < size; i++ {
		curr.Next = &ListNode{Val: rand.Intn(1000)}
		curr = curr.Next
	}
	return dummy.Next
}
