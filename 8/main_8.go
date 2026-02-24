package main

type CustomWG struct {
	done chan int
	size int
}

func NewCustomWG(n int) *CustomWG {
	return &CustomWG{
		done: make(chan int, n),
		size: n,
	}
}

func (cwg *CustomWG) Done() {
	cwg.done <- 1
}

func (cwg *CustomWG) Wait() {
	for i := 0; i < cwg.size; i++ {
		<-cwg.done
	}
}
