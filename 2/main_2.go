package main

import (
	"fmt"
	"math/rand"
)

func sliceExample(input []int) []int {
	var result []int
	for _, v := range input {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

func addElements(input []int, value int) []int {
	return append(input, value)
}

func copySlice(input []int) []int {
	newSlice := make([]int, len(input))
	copy(newSlice, input)
	return newSlice
}

func removeElement(input []int, index int) []int {
	if index < 0 || index >= len(input) {
		return input
	}
	return append(input[:index], input[index+1:]...)
}

func main() {
	originalSlice := make([]int, 10)
	for i := range originalSlice {
		originalSlice[i] = rand.Intn(100)
	}

	fmt.Println(originalSlice)
	fmt.Println(sliceExample(originalSlice))
	fmt.Println(addElements(originalSlice, 200))
	cp := copySlice(originalSlice)
	fmt.Println(cp)
	fmt.Println(removeElement(originalSlice, 2))
}
