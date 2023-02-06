package main

import "fmt"

func main() {
	b := xorOperation(5, 0)
	//b := xorOperation(4, 3)
	fmt.Println(b)
}

func xorOperation(n int, start int) int {
	var nums []int
	for i := 0; i < n; i++ {
		nums = append(nums, start+2*i)
	}

	var result int
	for _, num := range nums {
		result = result ^ num
	}

	return result
}
