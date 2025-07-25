package p66

import "fmt"

func main() {
	//digits := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	digits := []int{9}
	fmt.Println(digits)
	digits = plusOne(digits)
	fmt.Println(digits)
}

func plusOne(digits []int) []int {
	l := len(digits)
	ones := make([]int, l)
	ones[l-1] = 1

	for i := l - 1; i >= 0; i-- {
		d := digits[i]
		o := ones[i]
		s := d + o
		if s == 10 {
			digits[i] = 0
			if i-1 >= 0 {
				ones[i-1] = 1
			} else {
				digits = append([]int{1}, digits...)
			}
		} else {
			digits[i] = s
		}
	}

	return digits
}
