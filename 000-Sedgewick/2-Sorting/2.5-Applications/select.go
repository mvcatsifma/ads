package applications

import (
	"math/rand"
	"time"
)

// selectKthSmallest finds the kth smallest element in an unsorted array using quickselect.
//
// Algorithm: Uses partitioning from quicksort but only recurses into the partition
// containing the kth element. After partitioning around index j:
// - arr[0...j-1] contains elements < arr[j]
// - arr[j] is in its final sorted position
// - arr[j+1...n-1] contains elements > arr[j]
//
// If j == k, we found the kth smallest element.
// If j > k, the kth smallest is in the left partition.
// If j < k, the kth smallest is in the right partition.
//
// Performance characteristics:
// - Time: O(N) average case, O(N²) worst case
// - Space: O(1) - in-place
// - Note: Linear time on average (N + N/2 + N/4 + ... ≈ 2N)
//
// The array is shuffled before selection to guarantee probabilistic O(N) performance.
//
// Parameters:
// - arr: input array (will be modified/partitioned in place)
// - k: zero-based index of element to find (k=0 finds minimum, k=len(arr)-1 finds maximum)
//
// Returns: the kth smallest element
func selectKthSmallest(arr []int, k int) int {
	if k < 0 || k >= len(arr) {
		panic("k out of bounds")
	}

	// Shuffle array to guarantee probabilistic O(N) performance
	shuffle(arr)

	lo, hi := 0, len(arr)-1
	for hi > lo {
		j := partition(arr, lo, hi)
		if j == k {
			return arr[k]
		} else if j > k {
			hi = j - 1 // kth element is in left partition
		} else { // j < k
			lo = j + 1 // kth element is in right partition
		}
	}
	return arr[k]
}

// shuffle randomly shuffles an array in-place using Fisher-Yates algorithm.
// Each element has equal probability of ending up at any position.
func shuffle(arr []int) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := len(arr)
	for i := n - 1; i > 0; i-- {
		// Pick random index from 0 to i (inclusive)
		j := rng.Intn(i + 1)
		// Swap arr[i] with arr[j]
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// partition rearranges arr[lo:hi+1] so that arr[lo:j] < arr[j] <= arr[j+1:hi]
// Returns the final position of the pivot element.
// Uses random pivot selection to avoid worst-case performance.
func partition(arr []int, lo int, hi int) int {
	// Random pivot selection: choose random index and swap with first element
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := lo + rng.Intn(hi-lo+1)
	arr[lo], arr[randomIndex] = arr[randomIndex], arr[lo]

	// Initialize scan pointers: i scans left→right, j scans right→left
	i, j := lo, hi+1

	// Choose pivot element (now randomized)
	v := arr[lo]

	for {
		// Scan from left: find element >= pivot to swap
		i++
		for i <= hi && arr[i] < v {
			i++
		}

		// Scan from right: find element <= pivot to swap
		j--
		for j >= lo && v < arr[j] {
			j--
		}

		// Check if pointers have crossed
		if i >= j {
			break
		}

		// Swap elements that are out of place
		arr[i], arr[j] = arr[j], arr[i]
	}

	// Place pivot in its final sorted position
	arr[lo], arr[j] = arr[j], arr[lo]

	return j // Return pivot's final position
}
