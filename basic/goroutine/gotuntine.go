package main

import (
	"fmt"
)

func doWorker(id int, c chan int, w worker) {
	for n := range c {
		fmt.Printf("%d : %c\n", id, n)
		w.done <- true
	}
}

type worker struct {
	in   chan int
	done chan bool
}

func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWorker(id, w.in, w)
	return w
}

func chanDemo() {
	//	创建 channel 用make
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
		//<-workers[i].done
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
		//<-workers[i].done
	}
	//time.Sleep(time.Second)

	for _, worker := range workers {
		<-worker.done
		<-worker.done
	}
}

func main() {
	chanDemo()

}
