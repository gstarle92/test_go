package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	workers := 10
	tasks := 100
	wg.Add(workers)
	go Pool(workers, tasks)
	wg.Wait()
}
func Pool(workers int, tasks int) {
	ch := make(chan int)

	for i := 0; i < workers; i++ {
		go WorkerFun(&wg, ch)
	}

	for i := 0; i < tasks; i++ {
		ch <- i
	}
	close(ch)

}

func WorkerFun(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	for {
		task, ok := <-ch
		if !ok {
			return
		}
		fmt.Println("Completed task ", task)
		time.Sleep(time.Second * 1)
	}
}
