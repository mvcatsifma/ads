package main

var dw int
var steps int

func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	dw = 0
	steps = n
	process([]int{})
	return dw
}

func process(arr []int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	if sum > steps {
		return
	}
	if sum == steps {
		dw++
		return
	}

	process(append(arr, 1))
	process(append(arr, 2))
}
