// Package p981 implements a time-based key-value store that supports storing multiple values
// for the same key with different timestamps and retrieving values based on timestamp queries.
package p981

// TimeMap represents a time-based key-value data structure.
// It stores multiple timestamped values for each key and supports efficient timestamp-based queries.
type TimeMap struct {
	m map[string][]tmValue // Maps each key to a slice of timestamped values
}

// Constructor creates and returns a new empty TimeMap instance.
func Constructor() TimeMap {
	m := make(map[string][]tmValue)
	return TimeMap{
		m: m,
	}
}

// Set stores a key-value pair with the given timestamp.
// Values for the same key are stored in chronological order to enable binary search.
func (t *TimeMap) Set(key string, value string, timestamp int) {
	// Initialize slice for new keys
	if _, ok := t.m[key]; !ok {
		t.m[key] = make([]tmValue, 0)
	}
	// Append new timestamped value (assumes timestamps are added in non-decreasing order)
	t.m[key] = append(t.m[key], tmValue{
		value:     value,
		timestamp: timestamp,
	})
}

// Get retrieves the value associated with the key at the largest timestamp <= given timestamp.
// Returns empty string if no such value exists.
func (t *TimeMap) Get(key string, timestamp int) string {
	if v, ok := t.m[key]; !ok {
		return "" // Key doesn't exist
	} else {
		// Binary search for the largest timestamp <= given timestamp
		if idx := binarySearch(v, timestamp); idx != -1 {
			return v[idx].value
		}
	}
	return "" // No timestamp <= given timestamp found
}

// tmValue represents a timestamped value entry.
type tmValue struct {
	value     string // The stored value
	timestamp int    // The timestamp when this value was set
}

// binarySearch finds the index of the largest timestamp <= target timestamp.
// Returns -1 if no such timestamp exists.
// Time Complexity: O(log n) where n is the number of values for the key.
var binarySearch = func(values []tmValue, timestamp int) int {
	l, r := 0, len(values)-1

	res := -1 // Tracks the rightmost valid index found
	for l <= r {
		mid := l + (r-l)/2
		if values[mid].timestamp <= timestamp {
			res = mid   // Update result to current valid index
			l = mid + 1 // Look for a potentially larger valid timestamp on the right
		} else {
			r = mid - 1 // Current timestamp too large, search left
		}
	}

	return res
}
