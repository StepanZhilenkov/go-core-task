package main

import "testing"

func TestGeneratorRandomNumber(t *testing.T) {
	n := 5
	ch := GeneratorRandomNumber(n)

	count := 0
	for val := range ch {
		count++
		if val < 0 || val >= 100 {
			t.Errorf("Число %d вне диапазона [0, 100) на итерации %d", val, count)
		}
	}

	if count != n {
		t.Errorf("Ожидалось %d чисел, получено %d", n, count)
	}

	if _, ok := <-ch; ok {
		t.Error("Канал должен быть пуст и закрыт")
	}

}
