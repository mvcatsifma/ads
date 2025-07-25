package p167

func twoSum(numbers []int, target int) []int {
	start := 0
	end := len(numbers) - 1

	for {
		if start == end {
			break
		}
		n1 := numbers[start]
		n2 := numbers[end]

		sum := n1 + n2
		if sum == target {
			return []int{start + 1, end + 1}
		}
		if sum < target {
			start++
		}
		if sum > target {
			end--
		}
	}

	return []int{}
}
