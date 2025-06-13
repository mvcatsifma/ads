package main

import "fmt"

func main() {
	i := numberOfSteps(14)
	fmt.Println(i)
}

func numberOfSteps(num int) int {
	i := 0
	process(num, &i)
	return i
}

func process(num int, step *int) {
	if num == 0 { // base case
		return
	}

	if num%2 == 0 {
		num = num / 2
	} else {
		num--
	}

	*step++
	process(num, step)
}
