package p973

import (
	"container/heap"
	"math"
)

// Point represents a 2D coordinate point with its distance from the origin.
// Used for efficiently finding the k closest points to the origin (0,0).
type Point struct {
	x int     // X coordinate
	y int     // Y coordinate
	d float64 // Euclidean distance from origin: sqrt(x² + y²)
}

// PointHeap implements a min-heap of Points ordered by distance from origin.
// Satisfies the heap.Interface for use with the container/heap package.
// Points with smaller distances have higher priority (min-heap).
type PointHeap []Point

// Len returns the number of points in the heap.
func (h PointHeap) Len() int { return len(h) }

// Less compares two points by their distance from origin (min-heap: smaller distance = higher priority).
func (h PointHeap) Less(i, j int) bool { return h[i].d < h[j].d }

// Swap exchanges two points in the heap.
func (h PointHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push adds a point to the heap. Uses pointer receiver because it modifies slice length.
func (h *PointHeap) Push(x any) {
	*h = append(*h, x.(Point))
}

// Pop removes and returns the point with minimum distance. Uses pointer receiver.
func (h *PointHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// kClosest finds the k closest points to the origin (0,0) using Euclidean distance.
// Returns the k points with the smallest distances from the origin.
// Example: kClosest([[1,1],[2,2],[3,3]], 2) returns [[1,1],[2,2]]
//
// Approach: Uses a min-heap to efficiently extract the k smallest distances.
// Calculates Euclidean distance for each point, adds to heap, then pops k smallest.
// Time Complexity: O(n log n) where n is the number of points (heap operations).
// Space Complexity: O(n) for storing all points in the heap.
func kClosest(points [][]int, k int) [][]int {
	h := &PointHeap{}
	heap.Init(h)

	// Add all points to heap with calculated distances
	for _, p := range points {
		x, y := p[0], p[1]
		// Calculate Euclidean distance: √(x² + y²)
		d := math.Sqrt(math.Pow(float64(x), 2) + math.Pow(float64(y), 2))
		heap.Push(h, Point{x: x, y: y, d: d})
	}

	// Extract k closest points from min-heap
	var result [][]int
	for range k {
		point := heap.Pop(h).(Point)
		result = append(result, []int{point.x, point.y})
	}

	return result
}
