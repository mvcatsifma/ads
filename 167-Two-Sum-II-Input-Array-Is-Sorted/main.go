package main

func twoSum(numbers []int, target int) []int {
	m := make(map[int]int) // num -> index
	for i, num := range numbers {
		m[num] = i
	}

	for i, n1 := range numbers {
		n2 := target - n1 // we're looking for a second number so that n2+n1==target
		if j, ok := m[n2]; ok {
			return []int{i + 1, j + 1} // 1-indexed array
		}
	}

	return []int{}
}
