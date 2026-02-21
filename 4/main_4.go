package main

func Difference(s1, s2 []string) []string {
	var result []string

	for _, value1 := range s1 {
		foundValue := false

		for _, value2 := range s2 {
			if value1 == value2 {
				foundValue = true
				break
			}
		}
		if !foundValue {
			result = append(result, value1)
		}
	}
	return result
}
