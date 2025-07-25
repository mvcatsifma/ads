package p1720

import "fmt"

func main() {
	fmt.Println(1 ^ 0)
	fmt.Println(0 ^ 2)
	fmt.Println(2 ^ 1)
}

func decode(encoded []int, first int) []int {
	result := []int{first}
	for i, v := range encoded {
		result = append(result, v^result[i])
	}
	return result
}
