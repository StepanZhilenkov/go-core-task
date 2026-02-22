package main

import (
	"math/rand"
)

func GeneratorRandomNumber(n int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for i := 0; i < n; i++ {
			ch <- rand.Intn(100)
		}
	}()
	return ch
}
