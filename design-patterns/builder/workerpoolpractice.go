package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	workers := 5
	takslist := 10

	ch := make(chan int)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go workerworks(&wg, ch, i)
	}

	for i := 0; i < takslist; i++ {
		ch <- i
	}
	close(ch)
	wg.Wait()
}
func workerworks(wg *sync.WaitGroup, ch chan int, worker int) {
	defer wg.Done()

	for {
		task, ok := <-ch
		if !ok {
			return
		}
		time.Sleep(time.Second)
		fmt.Println(worker, " worker Completed task----->", task)

	}
}
