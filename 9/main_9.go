package main

import (
	"fmt"
	"math"
)

func main() {
	numsCh := make(chan uint8)
	resultCh := make(chan float64)

	go func() {
		defer close(numsCh)
		for i := uint8(1); i <= 5; i++ {
			numsCh <- i
		}
	}()

	go func() {
		defer close(resultCh)
		for val := range numsCh {
			resultCh <- math.Pow(float64(val), 3)
		}
	}()

	for res := range resultCh {
		fmt.Println(res)
	}
}
