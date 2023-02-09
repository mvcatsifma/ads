package main

func main() {

}

func mySqrt(x int) int {
	var result int

IntegerLoop:
	for i := 0; i <= x; i++ {
		if i*i <= x {
			result = i
		} else {
			break IntegerLoop
		}
	}

	return result
}
