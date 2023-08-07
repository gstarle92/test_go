package main

import (
	"fmt"
	"sync"
	"time"
)

type Student struct {
	Id    int
	FName string
	LName string
}

func main() {
	var wg sync.WaitGroup
	workers := 5
	tasks := 20
	wg.Add(workers)
	go Pool(&wg, workers, tasks)
	wg.Wait()
}

func Pool(wg *sync.WaitGroup, workers int, tasks int) {

	ch := make(chan Student)
	for i := 0; i < workers; i++ {
		go Workers(wg, ch, i)
	}
	for i := 0; i < tasks; i++ {
		ch <- Student{
			Id:    i,
			FName: "firstName",
			LName: "lastName",
		}
	}
	close(ch)
}

func Workers(wg *sync.WaitGroup, ch chan Student, worker int) {
	defer wg.Done()
	for {
		result, ok := <-ch
		if !ok {
			break
		}
		fmt.Printf("Task %d Cmpleted by worker %d \n ", result.Id, worker)
		time.Sleep(time.Second * 1)
	}

}
