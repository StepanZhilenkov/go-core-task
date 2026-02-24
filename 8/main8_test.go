package main

import "testing"

func TestCustomWG(t *testing.T) {
	n := 10
	wg := NewCustomWG(n)

	counter := 0

	for i := 0; i < n; i++ {
		go func() {
			counter++
			wg.Done()
		}()
	}
	wg.Wait()

	if counter != n {
		t.Errorf("Ожидали %d, получили %d", n, counter)
	}
}
