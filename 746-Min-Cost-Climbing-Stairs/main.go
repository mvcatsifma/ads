package p746

import "math"

// minCostClimbingStairs calculates the minimum cost to reach the top of a staircase
// using dynamic programming. You can climb either 1 or 2 steps at a time, and each
// step has an associated cost. You can start from either step 0 or step 1.
//
// The algorithm works backwards from the top, calculating the minimum cost to reach
// the end from each position by choosing the cheaper of the next two possible steps.
//
// Time complexity: O(n), Space complexity: O(1) additional space (modifies input)
func minCostClimbingStairs(cost []int) int {
	cost = append(cost, 0)

	for i := len(cost) - 3; i >= 0; i-- {
		cost[i] += int(math.Min(float64(cost[i+1]), float64(cost[i+2])))
	}

	return int(math.Min(float64(cost[0]), float64(cost[1])))
}
