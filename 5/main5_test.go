package main

import (
	"reflect"
	"testing"
)

func TestGetIntersection(t *testing.T) {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	foundValue, result := GetIntersection(a, b)

	if foundValue != true {
		t.Errorf("Ошибка: ожидалось true, получили %v", foundValue)
	}

	expected := []int{64, 3}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ошибка: ожидали %v, получили %v", expected, result)
	}
}
