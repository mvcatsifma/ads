package main

func main() {

}

func maximumWealth(accounts [][]int) int {
	var result int
	for _, a := range accounts {
		var wealth int
		for _, m := range a {
			wealth += m
		}
		if wealth > result {
			result = wealth
		}
	}
	return result
}
