package p118

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	result[0] = []int{1}

	for n := 1; n < numRows; n++ {
		result[n] = make([]int, n+1)
		for k := 0; k < n+1; k++ {
			v1 := getNthIntOrZero(result[n-1], k-1)
			v2 := getNthIntOrZero(result[n-1], k)
			result[n][k] = v1 + v2
		}
	}

	return result
}

func getNthIntOrZero(slice []int, n int) int {
	if n < 0 || n >= len(slice) {
		return 0
	}
	return slice[n]
}
