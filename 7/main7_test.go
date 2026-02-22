package main

import "testing"

func TestMerge(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 10
		ch1 <- 20
		close(ch1)
	}()

	go func() {
		ch2 <- 30
		close(ch2)
	}()

	merged := Merge(ch1, ch2)

	sum := 0
	count := 0

	for val := range merged {
		sum += val
		count++
	}

	if count != 3 {
		t.Errorf("Ожидали 3 числа, получили %d", count)
	}

	if sum != 60 {
		t.Errorf("Ожидали сумму 60, получили %d", sum)
	}
}
