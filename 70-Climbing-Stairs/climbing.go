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
	process(0)
	return dw
}

func process(sum int) {
	if sum > steps {
		return
	}
	if sum == steps {
		dw++
		return
	}

	sum1 := sum + 1
	sum2 := sum + 2
	process(sum1)
	process(sum2)
}
