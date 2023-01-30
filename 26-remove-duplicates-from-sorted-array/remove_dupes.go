package main

import "fmt"

func main() {
	arr := []int{1, 1, 2}
	//arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k := removeDuplicates(arr)
	fmt.Println(k)
	fmt.Println(arr)
}

func removeDuplicates(nums []int) int {
	numToBool := make(map[int]bool)
	var idx int
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if ok := numToBool[num]; !ok {
			numToBool[num] = true
			nums[idx] = num
			idx++
		}
	}
	return idx
}
