package p412

import "strconv"

func fizzBuzz(n int) []string {
	result := make([]string, n)
	for i := 1; i <= n; i++ {
		d3 := i%3 == 0
		d5 := i%5 == 0
		if d3 && d5 {
			result[i-1] = "FizzBuzz"
			continue
		}
		if d3 {
			result[i-1] = "Fizz"
			continue
		}
		if d5 {
			result[i-1] = "Buzz"
			continue
		}
		result[i-1] = strconv.Itoa(i)
	}
	return result
}
