package main

import (
	"slices"
	"testing"
)

func TestDifference(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date"}
	slice2 := []string{"banana", "date"}

	newSlice := []string{"apple", "cherry"}
	res := Difference(slice1, slice2)

	if !slices.Equal(res, newSlice) {
		t.Errorf("Ошибка: Получили %v, хотели %v", res, newSlice)
	}

}
