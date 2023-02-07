package main

func countConsistentStrings(allowed string, words []string) int {
	allowedMap := make(map[string]bool)
	for _, a := range allowed {
		allowedMap[string(a)] = true
	}

	var count int
	for _, w := range words {
		allAllowed := true
		for _, c := range w {
			if ok := allowedMap[string(c)]; !ok {
				allAllowed = false
			}
		}
		if allAllowed {
			count++
		}
	}

	return count
}
