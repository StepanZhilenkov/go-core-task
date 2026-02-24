package main

import (
	"math"
	"testing"
)

func TestMain(t *testing.T) {
	in := make(chan uint8)
	out := make(chan float64)

	go func() {
		defer close(out)
		for val := range in {
			out <- math.Pow(float64(val), 3)
		}
	}()

	in <- 2
	close(in)

	res := <-out
	if res != 8 {
		t.Errorf("Ожидали 8, получили %f", res)
	}

}
