package main

import (
	"slices"
	"testing"
)

func TestSliceExample(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	expected := []int{2, 4, 6}

	res := sliceExample(input)
	if !slices.Equal(res, expected) {
		t.Errorf("Ожидали: %v, получили: %v", expected, res)
	}
}

func TestAddElements(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []int{1, 2, 3, 4}

	res := addElements(input, 4)
	if !slices.Equal(res, expected) {
		t.Errorf("Ожидали: %v, получили: %v", expected, res)
	}
}

func TestCopySlice(t *testing.T) {
	input := []int{1, 2, 3}
	cp := copySlice(input)

	if !slices.Equal(input, cp) {
		t.Errorf("Значения в копии отличаются")
	}
	if len(cp) > 0 {
		cp[0] = 999
		if input[0] == 999 {
			t.Errorf("Это не копия, изменение дубликата затронуло оригинал")
		}
	}
}

func TestRemoveElement(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expected := []int{1, 2, 4}

	res := removeElement(input, 2)

	if !slices.Equal(res, expected) {
		t.Errorf("Ожидали: %v, получили: %v", expected, res)
	}
}
