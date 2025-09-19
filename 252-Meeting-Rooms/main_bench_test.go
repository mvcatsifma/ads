package p252

import "testing"

// BenchmarkCanAttendMeetings benchmarks the canAttendMeetings function with 1000 intervals.
//
// Results on Apple M3 Pro (darwin/arm64):
// BenchmarkCanAttendMeetings-12    	  334791	      3614 ns/op	   24672 B/op	       4 allocs/op
//
// Analysis:
// - ~3.6μs per operation: Exceptional performance for 1000 intervals
// - 24KB memory per op: Memory breakdown:
//   * testData slice: 1000 × 24 bytes (slice header) = 24KB
//   * Individual int arrays are shared/referenced, not copied deeply
//   * copy() only copies slice headers, not the underlying int arrays
// - Only 4 allocations: Minimal memory fragmentation
//   * 1 alloc: testData slice header ([][]int)
//   * 1 alloc: underlying array for 1000 slice headers (24KB)
//   * 2 allocs: Go's sort.Slice internal temporary allocations
//
// Performance insights:
// - Memory efficiency: copy() on [][]int only copies slice descriptors
// - The underlying [2]int arrays are shared between original and copy
// - Apple M3 Pro optimization: ARM64 + efficient cache utilization
// - Throughput: ~277K operations/second - production ready
//
// Memory model:
// - int size: platform-dependent (at least 32 bits, typically 64 bits on arm64)
// - Slice header: 24 bytes (pointer + len + cap, each 8 bytes on 64-bit)
// - Actual int data is shared, not duplicated by copy()
func BenchmarkCanAttendMeetings(b *testing.B) {
	// Generate test data with overlapping meetings
	intervals := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		start := i * 2
		end := start + 3 // Creates some overlaps
		intervals[i] = []int{start, end}
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		// Create a copy since the function modifies the slice
		testData := make([][]int, len(intervals))
		copy(testData, intervals)
		canAttendMeetings(testData)
	}
}
