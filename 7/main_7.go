package main

import "sync"

func Merge(in ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	out := make(chan int)

	output := func(c <-chan int) {
		defer wg.Done()

		for n := range c {
			out <- n
		}
	}
	wg.Add(len(in))

	for _, c := range in {
		go output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
