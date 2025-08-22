package p853

import (
	"sort"
)

type pair struct {
	pos   int
	speed int
}

// carFleet calculates the number of car fleets that will arrive at the target.
// Cars at the same position with the same speed form a fleet, and faster cars
// behind slower cars will catch up and form a fleet (traveling at the slower speed).
//
// Algorithm:
// 1. Pair each car's position with its speed
// 2. Sort cars by position (descending - closest to target first)
// 3. Calculate time to reach target for each car
// 4. If a car takes less/equal time than the car ahead, they form a fleet
//
// Parameters:
//   - target: the destination position
//   - position: slice of car positions
//   - speed: slice of car speeds (same length as position)
//
// Returns: number of distinct car fleets that reach the target
func carFleet(target int, position []int, speed []int) int {
	pairs := make([]pair, len(position)) // Value types, not pointers
	for i := 0; i < len(position); i++ {
		pairs[i] = pair{pos: position[i], speed: speed[i]} // Direct assignment
	}

	sortPairsByPos(pairs)

	result := make([]float64, 0, len(pairs))
	for _, p := range pairs {
		result = append(result, float64(target-p.pos)/float64(p.speed))
		if len(result) >= 2 && result[len(result)-1] <= result[len(result)-2] {
			result = result[:len(result)-1]
		}
	}

	return len(result)
}

// Sort pairs by position in descending order
func sortPairsByPos(pairs []pair) { // Changed from []*pair to []pair
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].pos > pairs[j].pos
	})
}
