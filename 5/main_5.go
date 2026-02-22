package main

func GetIntersection(a, b []int) (bool, []int) {
	set := make(map[int]bool)

	for _, val := range a {
		set[val] = true
	}

	var result []int

	for _, val := range b {
		if set[val] {
			result = append(result, val)
			set[val] = false
		}
	}
	return len(result) > 0, result
}
